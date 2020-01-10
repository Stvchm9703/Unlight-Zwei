package RawMod

import (
// // "github.com/jinzhu/gorm"
// _ "github.com/jinzhu/gorm/dialects/sqlite"
)

type CharaCardRawMod struct {
	ID int `gorm:"column:id;type:int(11);primary_key:true;" json:"id,omitempty"`
	//
	NameJP    string `gorm:"column:name;type:text" json:"name_jp,omitempty"`
	AbNameJP  string `gorm:"column:ab_name;type:text" json:"ab_name_jp,omitempty"`
	CaptionJP string `gorm:"column:caption;type:text" json:"caption_jp,omitempty"`
	//
	Level    int `gorm:"column:level;type:int(11)" json:"level,omitempty"`
	Hp       int `gorm:"column:hp;type:int(11)" json:"hp,omitempty"`
	Ap       int `gorm:"column:ap;type:int(11)" json:"ap,omitempty"`
	Dp       int `gorm:"column:dp;type:int(11)" json:"dp,omitempty"`
	Rarity   int `gorm:"column:rarity;type:int(11)" json:"rarity,omitempty"`
	DeckCost int `gorm:"column:deck_cost;type:int(11)" json:"deck_cost,omitempty"`
	Slot     int `gorm:"column:slot;type:int(11)" json:"slot,omitempty"`
	//
	StandImage    string `gorm:"column:stand_image;type:text" json:"stand_image,omitempty"`
	CharaImage    string `gorm:"column:chara_image;type:text" json:"chara_image,omitempty"`
	ArtifactImage string `gorm:"column:artifact_image;type:text" json:"artifact_image,omitempty"`
	BgImage       string `gorm:"column:bg_image;type:text" json:"bg_image,omitempty"`
	//
	CharactorID int    `gorm:"column:charactor_id;type:int(11)" json:"charactor_id,omitempty"`
	NextID      int    `gorm:"column:next_id;type:int(11)" json:"next_id,omitempty"`
	Kind        int    `gorm:"column:kind;type:int(11)" json:"kind,omitempty"`
	CreatedAt   string `gorm:"column:created_at;type:text" json:"created_at,omitempty"`

	//
	NameEn    string `gorm:"column:name_en;type:text" json:"name_en,omitempty"`
	AbNameEn  string `gorm:"column:ab_name_en;type:text" json:"ab_name_en,omitempty"`
	CaptionEn string `gorm:"column:caption_en;type:text" json:"caption_en,omitempty"`
	//
	NameFr    string `gorm:"column:name_fr;type:text" json:"name_fr,omitempty"`
	AbNameFr  string `gorm:"column:ab_name_fr;type:text" json:"ab_name_fr,omitempty"`
	CaptionFr string `gorm:"column:caption_fr;type:text" json:"caption_fr,omitempty"`
	//
	NameKr    string `gorm:"column:name_kr;type:text" json:"name_kr,omitempty"`
	AbNameKr  string `gorm:"column:ab_name_kr;type:text" json:"ab_name_kr,omitempty"`
	CaptionKr string `gorm:"column:caption_kr;type:text" json:"caption_kr,omitempty"`
	//
	NameScn    string `gorm:"column:name_scn;type:text" json:"name_scn,omitempty"`
	AbNameScn  string `gorm:"column:ab_name_scn;type:text" json:"ab_name_scn,omitempty"`
	CaptionScn string `gorm:"column:caption_scn;type:text" json:"caption_scn,omitempty"`
	//
	NameTcn    string `gorm:"column:name_tcn;type:text" json:"name_tcn,omitempty"`
	AbNameTcn  string `gorm:"column:ab_name_tcn;type:text" json:"ab_name_tcn,omitempty"`
	CaptionTcn string `gorm:"column:caption_tcn;type:text" json:"caption_tcn,omitempty"`
	//
	NameIna    string `gorm:"column:name_ina;type:text" json:"name_ina,omitempty"`
	AbNameIna  string `gorm:"column:ab_name_ina;type:text" json:"ab_name_ina,omitempty"`
	CaptionIna string `gorm:"column:caption_ina;type:text" json:"caption_ina,omitempty"`
	//
	NameThai    string `gorm:"column:name_thai;type:text" json:"name_thai,omitempty"`
	AbNameThai  string `gorm:"column:ab_name_thai;type:text" json:"ab_name_thai,omitempty"`
	CaptionThai string `gorm:"column:caption_thai;type:text" json:"caption_thai,omitempty"`
	Feats 		[]FeatRawMod `json:"feats"`
}

func (*CharaCardRawMod) TableName() string {
	return "chara_cards"
}
