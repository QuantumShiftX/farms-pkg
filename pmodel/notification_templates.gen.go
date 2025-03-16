package pmodel

import (
	"github.com/QuantumShiftX/farms-pkg/constants"
	"github.com/QuantumShiftX/golib/stores/gormx"
)

const TableNameNotificationTemplate = "notification_templates"

// NotificationTemplate represents the notification_templates table in the database.
type NotificationTemplate struct {
	ID          int64                         `gorm:"primaryKey;column:id" json:"id"`                       // 通知模板唯一标识ID
	TypeCode    constants.NotificationType    `gorm:"column:type_code;not null" json:"type_code"`           // 类型编码，如 'daily_greeting'
	SubTypeCode constants.NotificationSubType `gorm:"column:subtype_code;not null" json:"sub_type_code"`    // 子类型编码，如 'morning'
	IsActive    int16                         `gorm:"column:is_active;not null;default:1" json:"is_active"` // 是否启用：1-启用，2-禁用
	Remarks     string                        `gorm:"column:remarks" json:"remarks"`                        // 备注，描述模板的用途或说明
	//
	gormx.OperationBaseModel
	gormx.Model
}

// TableName sets the table name for NotificationTemplate
func (*NotificationTemplate) TableName() string {
	return TableNameNotificationTemplate
}
