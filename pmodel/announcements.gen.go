package pmodel

import (
	"github.com/QuantumShiftX/golib/stores/gormx"
)

const TableNameAnnouncement = "announcements"

type Announcement struct {
	ID           int64  `gorm:"primaryKey;autoIncrement;column:id"`
	LanguageCode string `gorm:"column:language_code;size:10;not null"`
	Name         string `gorm:"column:name;size:255;not null"`
	Title        string `gorm:"column:title;size:255"`
	Content      string `gorm:"column:content;type:text"`
	Status       int8   `gorm:"column:status;default:1;not null"`
	Sort         int64  `gorm:"column:sort;default:0"`
	StartTime    int64  `gorm:"column:start_time;not null"` // 待展示/展示中/已过期
	EndTime      int64  `gorm:"column:end_time;not null"`   // 待展示/展示中/已过期
	Remark       string `gorm:"column:remark;size:500"`
	//
	gormx.OperationBaseModel
	gormx.Model
}

// TableName specifies the database table name
func (*Announcement) TableName() string {
	return TableNameAnnouncement
}
