package pmodel

import (
	"github.com/QuantumShiftX/golib/stores/gormx"
	"gorm.io/datatypes"
)

const TableNameFaqItem = "faq_items"

// FAQ条目表模型
type FaqItem struct {
	ID         int64          `gorm:"primaryKey;autoIncrement;column:id;type:bigint"`
	CategoryID int64          `gorm:"column:category_id;type:bigint;not null"`
	Question   datatypes.JSON `gorm:"column:question;type:jsonb;not null;default:'{\"en\":\"\"}'"`
	Answer     datatypes.JSON `gorm:"column:answer;type:jsonb;not null;default:'{\"en\":\"\"}'"`
	Status     int8           `gorm:"column:status;type:smallint;not null;default:1"`
	Remarks    string         `gorm:"column:remarks;type:text"`
	//
	gormx.OperationBaseModel
	gormx.Model
}

func (*FaqItem) TableName() string {
	return TableNameFaqItem
}
