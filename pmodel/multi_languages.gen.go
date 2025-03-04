// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package pmodel

import "github.com/QuantumShiftX/golib/stores/gormx"

const TableNameMultiLanguage = "multi_languages"

// MultiLanguage mapped from table <multi_languages>
type MultiLanguage struct {
	ID            int64  `gorm:"column:id;primaryKey;autoIncrement:true;comment:主键Id" json:"id"`                           // 主键Id
	Key           string `gorm:"column:key;not null;index:multi_languages_key_index,priority:1;comment:多语言key" json:"key"` // 多语言key
	Text          string `gorm:"column:text;not null;comment:多语言文本" json:"text"`                                           // 多语言文本
	LanguageCode  string `gorm:"column:language_code;not null;comment:语言代码" json:"language_code"`                          // 语言代码
	//
	gormx.OperationBaseModel
	gormx.Model
}

// TableName MultiLanguage's table name
func (*MultiLanguage) TableName() string {
	return TableNameMultiLanguage
}
