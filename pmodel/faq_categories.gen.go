package pmodel

import (
	"github.com/QuantumShiftX/golib/stores/gormx"
	"gorm.io/datatypes"
)

const TableNameFaqCategory = "faq_categories"

// FAQ分类表模型
type FaqCategory struct {
	ID           int64          `gorm:"primaryKey;autoIncrement;column:id;type:bigint"`
	CategoryName datatypes.JSON `gorm:"column:category_name;type:jsonb;not null;default:'{\"en\":\"\"}'"`
	Sort         int64          `gorm:"column:sort;type:bigint;not null"`
	Status       int8           `gorm:"column:status;type:smallint;not null;default:1"`
	Remarks      string         `gorm:"column:remarks;type:text"`
	//
	gormx.OperationBaseModel
	gormx.Model
}

func (*FaqCategory) TableName() string {
	return TableNameFaqCategory
}
