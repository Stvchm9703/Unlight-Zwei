package Mod

import (
	// "github.com/jinzhu/gorm"
	// _ "github.com/jinzhu/gorm/dialects/sqlite"
	rm "exporter/RawMod"
	"fmt"
	"exporter/Fbs"
	flatbuffers"github.com/google/flatbuffers/go"
	// "io/ioutil"
)

type CCardMod struct {
	ID      *int               `gorm:"column:id;type:int(11)" json:"id,omitempty"`
	CardSet []*CCardModCardSet `json:"card_set,omitempty"`

	Name    *ModLangPara `gorm:"column:name;type:text" json:"name,omitempty"`
	AbName  *ModLangPara `gorm:"column:ab_name;type:text" json:"ab_name,omitempty"`
	Caption *ModLangPara `gorm:"column:caption;type:text" json:"caption,omitempty"`
	//
	Skill []*CCardModSkillSet `json:"skill,omitempty"`
}

type ModLangPara struct {
	Jp   *string `json:"jp,omitempty"`
	En   *string `json:"en,omitempty"`
	Fr   *string `json:"fr,omitempty"`
	Kr   *string `json:"kr,omitempty"`
	Scn  *string `json:"scn,omitempty"`
	Tcn  *string `json:"tcn,omitempty"`
	Ina  *string `json:"ina,omitempty"`
	Thai *string `json:"thai,omitempty"`
}

type CCardModCardSet struct {
	//
	ID *int `gorm:"column:id;type:int(11)" json:"id,omitempty"`
	//
	Level    *int `gorm:"column:level;type:int(11)" json:"level,omitempty"`
	Hp       *int `gorm:"column:hp;type:int(11)" json:"hp,omitempty"`
	Ap       *int `gorm:"column:ap;type:int(11)" json:"ap,omitempty"`
	Dp       *int `gorm:"column:dp;type:int(11)" json:"dp,omitempty"`
	Rarity   *int `gorm:"column:rarity;type:int(11)" json:"rarity,omitempty"`
	DeckCost *int `gorm:"column:deck_cost;type:int(11)" json:"deck_cost,omitempty"`
	Slot     *int `gorm:"column:slot;type:int(11)" json:"slot,omitempty"`
	//
	StandImage       *string   `gorm:"column:stand_image;type:text" json:"stand_image,omitempty"`
	StandImageSet    *ImageSet `json:"stand_image_set,omitempty"`
	CharaImage       *string   `gorm:"column:chara_image;type:text" json:"chara_image,omitempty"`
	CharaImageSet    *ImageSet `json:"stand_image_set,omitempty"`
	ArtifactImage    *string   `gorm:"column:artifact_image;type:text" json:"artifact_image,omitempty"`
	ArtifactImageSet *ImageSet `json:"artifact_image_set,omitempty"`
	BgImage          *string   `gorm:"column:bg_image;type:text" json:"bg_image,omitempty"`
	BgImageSet       *ImageSet `json:"bg_image_set,omitempty"`
	//
	// CharactorID string `gorm:"column:charactor_id;type:int(11)"`
	NextID    *int    `gorm:"column:next_id;type:int(11)" json:"next_id,omitempty"`
	Kind      *int    `gorm:"column:kind;type:int(11)" json:"kind,omitempty"`
	CreatedAt *string `gorm:"column:created_at;type:text" json:"created_at,omitempty"`
	//
	Skill        []*CCardModSkillSet `json:"skill,omitempty"`
	SkillPointer []*int              `json:"skill_pointer,omitempty"`
}

type CCardModSkillSet struct {
	ID     *int `gorm:"column:id;type:int(6)" json:"id,omitempty"`
	FeatNo *int `gorm:"column:feat_no;type:int(10)" json:"feat_no,omitempty"`
	Pow    *int `gorm:"column:pow;type:int(10)" json:"pow,omitempty"`

	DiceAttribute *string `gorm:"column:dice_attribute;type:varchar(255)" json:"dice_attribute,omitempty"`
	EffectImage   *string `gorm:"column:effect_image;type:varchar(255)" json:"effect_image,omitempty"`
	Condition     *string `gorm:"column:condition;type:varchar(255)" json:"condition,omitempty"`
	CreatedAt     *string `gorm:"column:created_at;type:varchar(255)" json:"created_at,omitempty"`

	Name    *ModLangPara `gorm:"column:name;type:text" json:"name,omitempty"`
	Caption *ModLangPara `gorm:"column:caption;type:text" json:"caption,omitempty"`
}
type SkillShortSet struct {
	ID     *int `gorm:"column:id;type:int(6)" json:"id,omitempty"`
	FeatNo *int `gorm:"column:feat_no;type:int(10)" json:"feat_no,omitempty"`
	Pow    *int `gorm:"column:pow;type:int(10)" json:"pow,omitempty"`
	DiceAttribute *string `gorm:"column:dice_attribute;type:varchar(255)" json:"dice_attribute,omitempty"`
	CreatedAt     *string `gorm:"column:created_at;type:varchar(255)" json:"created_at,omitempty"`
	Name    *string `gorm:"column:name;type:text" json:"name,omitempty"`
	TrigFunc []byte  `json:"trig_func,omitempty"`
}
type ImageSet struct {
	Name *string `json:"name,omitempty"`;
	Width  *int    `json:"width,omitempty"`
	Height *int    `json:"height,omitempty"`
}
func (this *CCardMod) ImportCharaCardRaw(cardList *[]rm.CharaCardRawMod) (bool, error) {
	// for k,v := range
	
	this.ID = &((*cardList)[0].CharactorID)
	fmt.Println(this.ID)
	this.Name = &ModLangPara{
		Jp : &((*cardList)[0].NameJP),
		En : &(*cardList)[0].NameEn,
		Fr : &(*cardList)[0].NameFr,
		Kr : &(*cardList)[0].NameKr,
		Scn: &(*cardList)[0].NameScn,
		Tcn: &(*cardList)[0].NameTcn,
		Ina: &(*cardList)[0].NameIna,
		Thai :&(*cardList)[0].NameThai,
	}

	this.AbName = &ModLangPara{
		Jp : &(*cardList)[0].AbNameJP,
		En : &(*cardList)[0].AbNameEn,
		Fr : &(*cardList)[0].AbNameFr,
		Kr : &(*cardList)[0].AbNameKr,
		Scn: &(*cardList)[0].AbNameScn,
		Tcn: &(*cardList)[0].AbNameTcn,
		Ina: &(*cardList)[0].AbNameIna,
		Thai :&(*cardList)[0].AbNameThai,
	}
	this.Caption = &ModLangPara{
		Jp : &(*cardList)[0].CaptionJP,
		En : &(*cardList)[0].CaptionEn,
		Fr : &(*cardList)[0].CaptionFr,
		Kr : &(*cardList)[0].CaptionKr,
		Scn: &(*cardList)[0].CaptionScn,
		Tcn: &(*cardList)[0].CaptionTcn,
		Ina: &(*cardList)[0].CaptionIna,
		Thai :&(*cardList)[0].CaptionThai,
	}
	for k := range *cardList{
		this.CardSet = append(
			this.CardSet, 
			BuildCardSet( (*cardList)[k]) )
	}
	var sklst []*CCardModSkillSet
	for k := range this.CardSet {
		for k1 := range (this.CardSet[k]).Skill {
			if len(sklst) == 0 {
				sklst = append(sklst, (this.CardSet[k]).Skill[k1])
			} else {
				skc := 0
				for _,v:= range sklst {
					if ( *v.ID  == *(this.CardSet[k]).Skill[k1].ID){
						skc++
					} 
				}
				if skc == 0 {
					sklst = append(sklst, (this.CardSet[k]).Skill[k1])
				}
			}
		}
	}
	this.Skill = sklst
	
	return false, nil
}

func BuildCardSet( raw (rm.CharaCardRawMod)) *CCardModCardSet {
	tmp :=  CCardModCardSet{
		ID : &(raw.ID),
		Level : &raw.Level,
		Hp : &raw.Hp,
		Ap : &raw.Ap,
		Dp : &raw.Dp,
		Rarity : &raw.Rarity ,
		DeckCost: &raw.DeckCost,
		Slot : &raw.Slot,
		StandImage : &raw.StandImage,
		CharaImage : &raw.CharaImage,
		ArtifactImage :&raw.ArtifactImage,
		BgImage:&raw.BgImage,
		NextID: &raw.NextID,
		Kind : &raw.Kind,
		CreatedAt: &raw.CreatedAt,
	}
	// tmp.ImportFeatRaw()
	for k := range raw.Feats{
		tmp.SkillPointer = append(tmp.SkillPointer, &(raw.Feats[k].ID))
		tmp.Skill = append(tmp.Skill , &CCardModSkillSet{
			ID : &raw.Feats[k].ID,
			FeatNo : &raw.Feats[k].FeatNo,
			Pow : &raw.Feats[k].Pow,
			DiceAttribute : &raw.Feats[k].DiceAttribute,
			EffectImage   : &raw.Feats[k].EffectImage,
			Condition  : &raw.Feats[k].Condition,
			CreatedAt : &raw.Feats[k].CreatedAt,    
			Name : &ModLangPara{
				Jp : &raw.Feats[k].NameJP,
				En : &raw.Feats[k].NameEn,
				Fr : &raw.Feats[k].NameFr,
				Kr : &raw.Feats[k].NameKr,
				Scn: &raw.Feats[k].NameScn,
				Tcn: &raw.Feats[k].NameTcn,
				Ina: &raw.Feats[k].NameIna,
				Thai :&raw.Feats[k].NameThai,
			},
			Caption : &ModLangPara{
				Jp : &raw.Feats[k].CaptionJP,
				En : &raw.Feats[k].CaptionEn,
				Fr : &raw.Feats[k].CaptionFr,
				Kr : &raw.Feats[k].CaptionKr,
				Scn: &raw.Feats[k].CaptionScn,
				Tcn: &raw.Feats[k].CaptionTcn,
				Ina: &raw.Feats[k].CaptionIna,
				Thai :&raw.Feats[k].CaptionThai,
			},
		})
	}
	
	return &tmp
}



func ExtractSkill(raw *(rm.CharaCardRawMod)) []*CCardModSkillSet {
	var tmp []*CCardModSkillSet
	for _,v := range (*raw).Feats{
		tmp = append(tmp, &CCardModSkillSet{
			ID : &v.ID,
			FeatNo : &v.FeatNo,
			Pow : &v.Pow,
			DiceAttribute : &v.DiceAttribute,
			EffectImage   : &v.EffectImage,
			Condition  : &v.Condition,
			CreatedAt : &v.CreatedAt,    
			Name : &ModLangPara{
				Jp : &v.NameJP,
				En : &v.NameEn,
				Fr : &v.NameFr,
				Kr : &v.NameKr,
				Scn: &v.NameScn,
				Tcn: &v.NameTcn,
				Ina: &v.NameIna,
				Thai :&v.NameThai,
			},
			Caption : &ModLangPara{
				Jp : &v.CaptionJP,
				En : &v.CaptionEn,
				Fr : &v.CaptionFr,
				Kr : &v.CaptionKr,
				Scn: &v.CaptionScn,
				Tcn: &v.CaptionTcn,
				Ina: &v.CaptionIna,
				Thai :&v.CaptionThai,
			},
		})
	}
	return tmp
}

func (this *CCardModSkillSet) ExportAsShort() *SkillShortSet {
	return &SkillShortSet{
		ID : this.ID,
		FeatNo : this.FeatNo,
		Pow : this.Pow,
		DiceAttribute : this.DiceAttribute,
		CreatedAt : this.CreatedAt,    
		Name : this.Name.En,
	}
} 


func (this *SkillShortSet) ImportToFbs(fbb *flatbuffers.Builder) flatbuffers.UOffsetT {
	var (
		id int32
		fn int8
		pw int16
		da,n flatbuffers.UOffsetT
		 	
	) 
	if this.ID != nil{
		id = int32(*this.ID)
	}
	
	if this.FeatNo != nil{
		fn = int8(*this.FeatNo) 
	} 
	if  this.Pow != nil {
		pw = int16(*this.Pow)
	}
	
	if this.DiceAttribute != nil {
		da = fbb.CreateString(*this.DiceAttribute  )
	} 
	if this.Name !=nil{
		n = fbb.CreateString(*this.Name )
	}
	// fbb.CreateByteString(this.TrigFunc)
	
	Fbs.SkillStartTrigFuncVector(fbb, len(this.TrigFunc))
	for i:=len(this.TrigFunc); i > 0; i--{
		fbb.PrependByte(this.TrigFunc[i-1])
	}
	tf := fbb.EndVector(len(this.TrigFunc))

	Fbs.SkillStart(fbb) 
	{
		Fbs.SkillAddId(fbb, id)
		Fbs.SkillAddFeatNo(fbb,fn )
		Fbs.SkillAddPow(fbb, pw)
		Fbs.SkillAddDiceAttr(fbb, da)
		Fbs.SkillAddName(fbb, n)
		Fbs.SkillAddTrigFunc(fbb, tf)
	}
	return  Fbs.SkillEnd(fbb) 
}

func (this *SkillShortSet) ImportFmFbs(skillbyte *Fbs.Skill) {
	
	id := (skillbyte.Id())
	sid := int(id) 
	this.ID = &sid
	fn :=  int(skillbyte.FeatNo())
	this.FeatNo = &fn
	pow := int(skillbyte.Pow()) 
	this.Pow = &pow
	da := string(skillbyte.DiceAttr())
	this.DiceAttribute = &da
	nm := string(skillbyte.Name())
	this.Name = &nm 
	fmt.Println(skillbyte.TrigFunc(0))
	this.TrigFunc = skillbyte.TrigFuncBytes()
}
