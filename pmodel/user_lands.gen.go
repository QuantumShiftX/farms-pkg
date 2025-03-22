package pmodel

import (
	"github.com/QuantumShiftX/golib/stores/gormx"
)

const TableNameUserLand = "user_lands"

// UserLand 用户土地信息表
type UserLand struct {
	ID       int64 `gorm:"column:id;primary_key" json:"id"`
	LandID   int64 `gorm:"column:land_id;comment:土地ID" json:"land_id"`
	UserID   int64 `gorm:"column:user_id;comment:用户ID" json:"user_id"`
	FarmID   int64 `gorm:"column:farm_id;comment:农场ID" json:"farm_id"`
	PlantID  int64 `gorm:"column:plant_id;comment:种植记录ID" json:"plant_id"`
	Position int64 `gorm:"column:position;comment:土地位置编号（1-20）" json:"position"`
	Status   int64 `gorm:"column:status;default:1;comment:土地状态：1-未开垦，2-已开垦未种植，3-已种植" json:"status"`
	//
	gormx.Model
}

// TableName 指定表名
func (*UserLand) TableName() string {
	return TableNameUserLand
}
