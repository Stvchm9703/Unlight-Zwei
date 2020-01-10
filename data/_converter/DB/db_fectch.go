package DB

import (
	rm "exporter/RawMod"
	"fmt"

	"github.com/jinzhu/gorm"
)

var db_path = "127.0.0.1:3306"
var db_name = "unlight_old_db"
var db_user = "root:"

func FetchCharactor() ([]rm.CharaCardRawMod, error) {
	db, err := gorm.Open("mysql", db_user+"@tcp("+db_path+")/"+db_name+"?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	var tmp_result []rm.CharaCardRawMod
	db.Model(&rm.CharaCardRawMod{}).Group("charactor_id").Having("level = 1 AND kind = 0").Order("charactor_id asc").Find(&tmp_result)
	return tmp_result, nil
}

func FetchFullCC(cc *int) ([]rm.CharaCardRawMod, error) {
	db, err := gorm.Open("mysql", db_user+"@tcp("+db_path+")/"+db_name+"?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var tmp_result []rm.CharaCardRawMod
	fmt.Println(*cc)
	db.Where(&rm.CharaCardRawMod{
		CharactorID: *cc,
		Kind:        0,
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
	return tmp_result, nil
}
