package pmodel

import "github.com/QuantumShiftX/golib/stores/gormx"

const TableNameNotificationTranslation = "notification_translations"

// NotificationTranslation represents the notification_translations table in the database.
type NotificationTranslation struct {
	ID           int64  `gorm:"primaryKey;column:id" json:"id"`                             // 唯一标识ID
	TemplateID   int64  `gorm:"column:template_id;not null" json:"template_id"`             // 关联的通知模板ID
	LanguageCode string `gorm:"column:language_code;size:10;not null" json:"language_code"` // 语言编码，如 'zh', 'en'
	Title        string `gorm:"column:title;not null" json:"title"`                         // 通知标题
	Content      string `gorm:"column:content;not null" json:"content"`                     // 通知内容
	Remarks      string `gorm:"column:remarks" json:"remarks"`                              // 备注，描述模板的用途或说明
	//
	gormx.OperationBaseModel
	gormx.Model
}

// TableName sets the table name for NotificationTranslation
func (*NotificationTranslation) TableName() string {
	return TableNameNotificationTranslation
}
