// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameTOpmVoice = "t_opm_voice"

// TOpmVoice mapped from table <t_opm_voice>
type TOpmVoice struct {
	Id            int64  `gorm:"column:id;primaryKey" json:"id"`                   // 语控id
	VoiceNo       string `gorm:"column:voice_no" json:"voiceNo"`                   // 语控编码;xiaoai,tianmao,alex,google
	VoiceName     string `gorm:"column:voice_name" json:"voiceName"`               // 语控名称
	VoiceCategory string `gorm:"column:voice_category" json:"voiceCategory"`       // 语控支持的品类;json字符串(暂留.后续使用)
	VoiceDoc      string `gorm:"column:voice_doc" json:"voiceDoc"`                 // 参考文档;验证指导
	VoiceLogo     string `gorm:"column:voice_logo" json:"voiceLogo"`               // 语控图标
	VoiceEnable   int32  `gorm:"column:voice_enable;default:2" json:"voiceEnable"` // 是否启用
	VoiceDesc     string `gorm:"column:voice_desc" json:"voiceDesc"`               // 语音平台描述
	VoiceEnName   string `gorm:"column:voice_en_name" json:"voiceEnName"`          // 语控英文名称
}

// TableName TOpmVoice's table name
func (*TOpmVoice) TableName() string {
	return TableNameTOpmVoice
}
