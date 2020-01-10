package RawMod


type FeatRawMod struct {
	ID     int `gorm:"column:id;type:int(6);primary_key:true;" json:"id"`
	FeatNo int `gorm:"column:feat_no;type:int(10)" json:"feat_no"`
	Pow    int `gorm:"column:pow;type:int(10)" json:"pow"`

	//
	NameJP    string `gorm:"column:name;type:varchar(255)" json:"name_jp"`
	CaptionJP string `gorm:"column:caption;type:varchar(255)" json:"caption_jp"`
	//

	DiceAttribute string `gorm:"column:dice_attribute;type:varchar(255)" json:"dice_attribute"`
	EffectImage   string `gorm:"column:effect_image;type:varchar(255)" json:"effect_image"`
	Condition     string `gorm:"column:condition;type:varchar(255)" json:"condition"`
	CreatedAt     string `gorm:"column:created_at;type:varchar(255)" json:"created_at"`

	//
	NameEn    string `gorm:"column:name_en;type:varchar(255)" json:"name_en"`
	CaptionEn string `gorm:"column:caption_en;type:varchar(255)" json:"caption_en"`
	//
	NameFr    string `gorm:"column:name_fr;type:varchar(255)" json:"name_fr"`
	CaptionFr string `gorm:"column:caption_fr;type:varchar(255)" json:"caption_fr"`
	//
	NameKr    string `gorm:"column:name_kr;type:varchar(255)" json:"name_kr"`
	CaptionKr string `gorm:"column:caption_kr;type:varchar(255)" json:"caption_kr"`
	//
	NameScn    string `gorm:"column:name_scn;type:varchar(255)" json:"name_scn"`
	CaptionScn string `gorm:"column:caption_scn;type:varchar(255)" json:"caption_scn"`
	//
	NameTcn    string `gorm:"column:name_tcn;type:varchar(255)" json:"name_tcn"`
	CaptionTcn string `gorm:"column:caption_tcn;type:varchar(255)" json:"caption_tcn"`
	//
	NameIna    string `gorm:"column:name_ina;type:varchar(255)" json:"name_ina"`
	CaptionIna string `gorm:"column:caption_ina;type:varchar(255)" json:"caption_ina"`
	//
	NameThai    string `gorm:"column:name_thai;type:varchar(255)" json:"name_thai"`
	CaptionThai string `gorm:"column:caption_thai;type:varchar(255)" json:"caption_thai"`
}

func (*FeatRawMod) TableName() string {
	return "feats"
}

type FeatInventRawMod struct {
	ID int `gorm:"column:id;type:int(6)" json:"id"`
	CharaCardID int `gorm:"column:chara_card_id"`
	FeatID int `gorm:"column:feat_id"`
}

func (*FeatInventRawMod) TableName() string {
	return "feat_inventories"
}