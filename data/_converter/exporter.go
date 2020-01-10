package main

import (
	"encoding/json"
	"exporter/DB"
	"exporter/Fbs"
	m "exporter/Mod"
	rm "exporter/RawMod"

	"fmt"
	"os"
	"path/filepath"

	// "strconv"
	"time"

	"io/ioutil"

	flatbuffers "github.com/google/flatbuffers/go"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	mgo "gopkg.in/mgo.v2"
)

var db_path = "127.0.0.1:3306"
var db_name = "unlight_old_db"
var db_user = "root:"

var (
	mgo_path     = "127.0.0.1:27017"
	mgo_username = ""
	mgo_password = ""
	mgo_database = "unlight_asset"

	mgo_session *mgo.Session
	mgo_DBC     *mgo.Database
)

var export_folder_path = "A:\\Gitrepo\\Unlight-Zwei\\data\\output_ph1"

func main() {
	fmt.Println("start time", time.Now())
	mainRun()
	fmt.Println()
	fmt.Println("end time", time.Now())
}

func mainRun() {
	// mgo_session, err := ConnectMongoDB()
	// if err != nil {
	// 	fmt.Println(err)
	// } else{
	// 	mgo_DBC = mgo_session.DB(mgo_database)
	// 	defer mgo_session.Close()
	// }

	a, b := DB.FetchCharactor()
	if b != nil {
		fmt.Println(b)
	}
	// fmt.Println(a)

	var shortskill []*m.SkillShortSet
	for v := range a {
		tmp, _ := DB.FetchFullCC(&(a[v].CharactorID))
		fmt.Println("v:", v)
		var tt m.CCardMod
		tt.ImportCharaCardRaw(&tmp)
		ExportAsCCAsset(&tt)
		// ImportDataToMongo(mgo_DBC, &tt)
		shortskill = append(shortskill, (getSSSet(&tt))...)
	}
	ExportAsBackendAsset(shortskill)
}

func testRun() (bool, error) {
	db, err := gorm.Open("mysql", db_user+"@tcp("+db_path+")/"+db_name+"?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
	if err != nil {
		fmt.Println(err)
		return false, err
	}

	var tmp_result []rm.CharaCardRawMod
	db.Where(&rm.CharaCardRawMod{
		CharactorID: 1,
	}).Find(&tmp_result)

	for k := range tmp_result {
		var tt []rm.FeatInventRawMod
		db.Where(&rm.FeatInventRawMod{
			CharaCardID: tmp_result[k].ID,
		}).Find(&tt)
		var ty []rm.FeatRawMod
		for _, v1 := range tt {
			var tu []rm.FeatRawMod
			db.Where(&rm.FeatRawMod{
				ID: v1.FeatID,
			}).First(&tu)
			ty = append(ty, tu...)
		}
		tmp_result[k].Feats = ty
	}

	jsona, _ := json.Marshal(tmp_result)
	fmt.Println(string(jsona))

	return true, nil
}

func ExportAsCCAsset(dataset *m.CCardMod) {
	// create folder = cc-id
	file_name := fmt.Sprintf("%02d", *dataset.ID)
	CCfolder := filepath.Join(export_folder_path, "cc"+file_name)
	os.MkdirAll(CCfolder, os.ModePerm)

	file, _ := os.OpenFile(filepath.Join(CCfolder, "skill.json"), os.O_CREATE, os.ModePerm)
	encoder := json.NewEncoder(file)
	encoder.Encode(dataset.Skill)
	file.Close()

	var Otmp m.CCardMod = *dataset
	Otmp.Skill = nil
	for k := range Otmp.CardSet {
		(Otmp.CardSet[k]).Skill = nil
	}

	file, _ = os.OpenFile(filepath.Join(CCfolder, "card_set.json"), os.O_CREATE, os.ModePerm)
	encoder = json.NewEncoder(file)
	encoder.Encode(Otmp)
	file.Close()

	file, _ = os.OpenFile(filepath.Join(CCfolder, "raw_set.json"), os.O_CREATE, os.ModePerm)
	encoder = json.NewEncoder(file)
	encoder.Encode(dataset)
	file.Close()
}

// -----------------------------------------------------
// MongoDB Related
func ConnectMongoDB() (*mgo.Session, error) {
	Host := []string{
		mgo_path,
	}
	DBSess, err := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:    Host,
		Username: mgo_username,
		Password: mgo_password,
	})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	// DBConn = DBSess.DB(mgo_database)
	return DBSess, nil
}

func ImportDataToMongo(DBConn *mgo.Database, dataset *m.CCardMod) (bool, error) {
	failList := []map[string]interface{}{}

	// Charactor Card
	err := DBConn.C("CharactorCard").Insert(dataset)

	if err != nil {
		errmsg := map[string]interface{}{
			"modname":    "CharactorCard",
			"batchindex": dataset.ID,
			"refDoc":     dataset,
			"err":        err,
		}
		failList = append(failList, errmsg)
	}

	// Skill
	for _, v := range dataset.Skill {
		err := DBConn.C("Skill").Insert(v)
		if err != nil {
			errmsg := map[string]interface{}{
				"modname":    "Skill",
				"batchindex": dataset.ID,
				"refDoc":     dataset,
				"err":        err,
			}
			failList = append(failList, errmsg)
		}

	}
	if len(failList) > 0 {
		return false, ErrorMessage{
			When:     time.Now(),
			What:     "run import batch err:",
			ErrorObj: failList,
		}
	}
	return true, nil
}

// ------------------------------------------------------
type ErrorMessage struct {
	When     time.Time
	What     string
	ErrorObj interface{}
}

func (e ErrorMessage) Error() string {
	return fmt.Sprintf("%v: %v \n %v", e.When, e.What, e.ErrorObj)
}

func getSSSet(dataset *m.CCardMod) (result []*m.SkillShortSet) {
	for k := range dataset.Skill {
		result = append(result, dataset.Skill[k].ExportAsShort())
	}
	return
}

func ExportAsBackendAsset(data []*m.SkillShortSet) {
	BEfolder := filepath.Join(export_folder_path, "backend")
	os.MkdirAll(BEfolder, os.ModePerm)

	file, _ := os.OpenFile(filepath.Join(BEfolder, "skill.bson"), os.O_CREATE, os.ModePerm)
	encoder := json.NewEncoder(file)
	encoder.Encode(data)
	file.Close()
}

func WriteSkillListFBC(shortskill []*m.SkillShortSet) []byte {
	builder := flatbuffers.NewBuilder(1024)
	var skillOffset []flatbuffers.UOffsetT
	for k := range shortskill {
		skillOffset = append(skillOffset,
			shortskill[k].ImportToFbs(builder))
		fmt.Print(k, ":")
	}
	fmt.Println()

	Fbs.SkillListStartSkillListVector(builder, len(skillOffset))
	for i := len(skillOffset); i > 0; i-- {
		builder.PrependUOffsetT(skillOffset[i-1])
	}
	skillfin := builder.EndVector(len(skillOffset))
	lu := builder.CreateString(time.Now().Format("2006-01-02 15:04:05"))

	Fbs.SkillListStart(builder)
	{
		Fbs.SkillListAddSkillList(builder, skillfin)
		Fbs.SkillListAddLastUpdated(builder, lu)
	}
	slfile := Fbs.SkillListEnd(builder)
	builder.Finish(slfile)
	return builder.FinishedBytes()
}

func test_FB_Compile() {
	a, b := DB.FetchCharactor()
	if b != nil {
		fmt.Println(b)
	}
	var shortskill []*m.SkillShortSet
	for v := range a {
		tmp, _ := DB.FetchFullCC(&(a[v].CharactorID))
		fmt.Println("v:", v)
		var tt m.CCardMod
		tt.ImportCharaCardRaw(&tmp)
		shortskill = append(shortskill, (getSSSet(&tt))...)
	}
	// ExportAsBackendAsset(shortskill)

	buf := WriteSkillListFBC(shortskill)
	// write file
	BEfolder := filepath.Join(export_folder_path, "backend")
	file, _ := os.OpenFile(filepath.Join(BEfolder, "skill.d.fbc"), os.O_CREATE, os.ModePerm)
	file.Write(buf)
	file.Close()
}

func test_FB_Decompile() {

	BEfolder := filepath.Join(export_folder_path, "backend")
	filebyte, _ := ioutil.ReadFile(filepath.Join(BEfolder, "skill.d.fbc"))
	SkillListBase := Fbs.GetRootAsSkillList(filebyte, 0)

	fmt.Println(string(SkillListBase.LastUpdated()))
	var skill_list []*m.SkillShortSet
	for i := SkillListBase.SkillListLength(); i > 0; i-- {
		tmp := new(Fbs.Skill)
		var tmpSkill m.SkillShortSet
		fmt.Println(i, ": de _ ")
		if SkillListBase.SkillList(tmp, i-1) {
			Id := tmp.Id()
			fmt.Print("ID:", Id)
			tmpSkill.ImportFmFbs(tmp)
			skill_list = append(skill_list, &tmpSkill)
		}
		fmt.Println("Skill no:", *tmpSkill.ID, "\t:", *tmpSkill.Name)
	}
}

func test_FB_compile_run() {
	BEfolder := filepath.Join(export_folder_path, "backend")
	filebyte, _ := ioutil.ReadFile(filepath.Join(BEfolder, "go.mod.c"))

	fmt.Println("test file: go.mod.c")
	fmt.Println(string(filebyte))

	test_skill := &m.SkillShortSet{
		TrigFunc: filebyte,
	}

	test_list := []*m.SkillShortSet{
		test_skill,
	}

	buf := WriteSkillListFBC(test_list)

	SkillListBase := Fbs.GetRootAsSkillList(buf, 0)

	fmt.Println(string(SkillListBase.LastUpdated()))
	// var skill_list []*m.SkillShortSet
	// for i := SkillListBase.SkillListLength(); i > 0 ; i-- {
	tmp := new(Fbs.Skill)
	fmt.Println(" test field")
	var TestSkill m.SkillShortSet
	if SkillListBase.SkillList(tmp, 0) {
		TestSkill.ImportFmFbs(tmp)
	}
	fmt.Println("Skill no:", *TestSkill.ID, "\t:", *TestSkill.Name)
	fmt.Println("raw file")
	fmt.Println(string(TestSkill.TrigFunc))
}
