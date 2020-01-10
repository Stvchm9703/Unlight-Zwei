package main

import (
	"encoding/json"
	mod "exporter/Mod"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	fmt.Println("start :", time.Now())
	main_proc()
	fmt.Println("end:", time.Now())
}

var (
	png_root_path  = "A:\\Gitrepo\\Unlight-Zwei\\data\\png\\char"
	json_root_path = "A:\\Gitrepo\\Unlight-Zwei\\data\\output_ph1"
	out_root_path  = "A:\\Gitrepo\\Unlight-Zwei\\data\\output_ph2"
)

func main_proc() {
	os.MkdirAll(out_root_path, os.ModePerm)
	log.Println("start up importer")

	var json_pathList, cardCode []string
	err := filepath.Walk(json_root_path,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() {
				json_pathList = append(json_pathList, path)
				cardCode = append(cardCode, strings.ReplaceAll(path, json_root_path+"\\", ""))
				fmt.Println(path)
			}
			return nil
		})
	if err != nil {
		log.Println(err)
	}

	for _, v := range cardCode {
		fmt.Println("in:", v)
		tar := filepath.Join(png_root_path, v)
		if _, err := os.Stat(tar); !os.IsNotExist(err) {
			cardMoveProc(&v)
		}
	}
}

func cardMoveProc(fileCode *string) {
	os.MkdirAll(filepath.Join(out_root_path, *fileCode), os.ModePerm)

	pngpath := filepath.Join(png_root_path, *fileCode)

	// copy(
	// 	filepath.Join(json_root_path, *fileCode, "card_set.json"),
	// 	filepath.Join(out_root_path, *fileCode, "card_set.json"),
	// )
	// copy(
	// 	filepath.Join(json_root_path, *fileCode, "skill_set.json"),
	// 	filepath.Join(out_root_path, *fileCode, "skill_set.json"),
	// )
	var tr []*mod.CCardModSkillSet

	jsonFile, err := os.Open(filepath.Join(json_root_path, *fileCode, "skill.json"))
	if err != nil {
		fmt.Println(err)
	}
	skill_source, _ := ioutil.ReadAll(jsonFile)
	jsonFile.Close()
	// fmt.Print(string(skill_source))
	json.Unmarshal(skill_source, &tr)
	// skill_source.Close()
	for k := range tr {
		if strings.Contains(*tr[k].EffectImage, "_skl00.swf") || strings.Contains(*tr[k].EffectImage, "_sklex00.swf") {
			*tr[k].EffectImage = "skill_0_1.png"
		} else if strings.Contains(*tr[k].EffectImage, "_skl01.swf") || strings.Contains(*tr[k].EffectImage, "_sklex01.swf") {
			*tr[k].EffectImage = "skill_1_1.png"
		} else if strings.Contains(*tr[k].EffectImage, "_skl02.swf") || strings.Contains(*tr[k].EffectImage, "_sklex02.swf") {
			*tr[k].EffectImage = "skill_2_1.png"
		} else if strings.Contains(*tr[k].EffectImage, "_skl03.swf") || strings.Contains(*tr[k].EffectImage, "_sklex03.swf") {
			*tr[k].EffectImage = "skill_3_1.png"
		}
	}
	newSkill_source, _ := os.OpenFile(filepath.Join(out_root_path, *fileCode, "skill.json"), os.O_CREATE, os.ModePerm)
	encoder := json.NewEncoder(newSkill_source)
	encoder.Encode(tr)

	var tyu mod.CCardMod
	jsonFile, err = os.Open(filepath.Join(json_root_path, *fileCode, "card_set.json"))
	if err != nil {
		fmt.Println(err)
	}
	card_source, _ := ioutil.ReadAll(jsonFile)

	// fmt.Println(string(card_source))
	json.Unmarshal(card_source, &tyu)
	// card_source.Close()
	for k := range tyu.CardSet {
		*tyu.CardSet[k].StandImage = "stand_1.png"
		*tyu.CardSet[k].ArtifactImage = "stand_f.png"
		*tyu.CardSet[k].BgImage = "stand_shad.png"
		if strings.Contains(*tyu.CardSet[k].CharaImage, "_01.swf") {
			*tyu.CardSet[k].CharaImage = "cover_1.png"
		} else if strings.Contains(*tyu.CardSet[k].CharaImage, "_02.swf") {
			*tyu.CardSet[k].CharaImage = "cover_2.png"
		} else if strings.Contains(*tyu.CardSet[k].CharaImage, "_03.swf") {
			*tyu.CardSet[k].CharaImage = "cover_3.png"
		} else if strings.Contains(*tyu.CardSet[k].CharaImage, "_04.swf") {
			*tyu.CardSet[k].CharaImage = "cover_4.png"
		} else if strings.Contains(*tyu.CardSet[k].CharaImage, "_05.swf") {
			*tyu.CardSet[k].CharaImage = "cover_5.png"
		} else if strings.Contains(*tyu.CardSet[k].CharaImage, "_a.swf") {
			*tyu.CardSet[k].CharaImage = "cover_a.png"
		} else if strings.Contains(*tyu.CardSet[k].CharaImage, "_b.swf") {
			*tyu.CardSet[k].CharaImage = "cover_b.png"
		} else if strings.Contains(*tyu.CardSet[k].CharaImage, "_c.swf") {
			*tyu.CardSet[k].CharaImage = "cover_c.png"
		} else if strings.Contains(*tyu.CardSet[k].CharaImage, "_d.swf") {
			*tyu.CardSet[k].CharaImage = "cover_d.png"
		} else if strings.Contains(*tyu.CardSet[k].CharaImage, "_e.swf") {
			*tyu.CardSet[k].CharaImage = "cover_e.png"
		}
	}
	newcard_source, _ := os.OpenFile(filepath.Join(out_root_path, *fileCode, "card_set.json"), os.O_CREATE, os.ModePerm)
	encoder = json.NewEncoder(newcard_source)
	encoder.Encode(tyu)

	err = filepath.Walk(pngpath,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			fmt.Println(path)
			if !info.IsDir() {
				fileName := strings.ReplaceAll(path, pngpath+"\\", "")
				fileName = strings.ReplaceAll(fileName, "\\", "_")
				copy(
					path,
					filepath.Join(out_root_path, *fileCode, fileName),
				)
			}
			return nil
		})
	if err != nil {
		log.Println(err)
	}

}

func copy(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}

func RewriteFile(src string) {

}
