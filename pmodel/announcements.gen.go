package pmodel

import (
	"github.com/QuantumShiftX/golib/stores/gormx"
)

const TableNameAnnouncement = "announcements"

type Announcement struct {
	ID           int64  `gorm:"primaryKey;autoIncrement;column:id" json:"id"`
	LanguageCode string `gorm:"column:language_code;size:10;not null" json:"language_code"`
	Name         string `gorm:"column:name;size:255;not null" json:"name"`
	Title        string `gorm:"column:title;size:255" json:"title"`
	Content      string `gorm:"column:content;type:text" json:"content"`
	Status       int8   `gorm:"column:status;default:1;not null" json:"status"`
	Sort         int64  `gorm:"column:sort;default:0" json:"sort"`
	StartTime    int64  `gorm:"column:start_time;not null" json:"start_time"` // 待展示/展示中/已过期
	EndTime      int64  `gorm:"column:end_time;not null" json:"end_time"`     // 待展示/展示中/已过期
	Remark       string `gorm:"column:remark;size:500" json:"remark"`
	//
	gormx.OperationBaseModel
	gormx.Model
}

// TableName specifies the database table name
func (*Announcement) TableName() string {
	return TableNameAnnouncement
}
