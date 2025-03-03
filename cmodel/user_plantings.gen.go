package cmodel

import (
	"github.com/QuantumShiftX/golib/stores/gormx"
	"github.com/google/uuid"
)

const TableNameUserPlanting = "user_plantings"

// UserPlanting represents the user_plantings table in ClickHouse
type UserPlanting struct {
	ID              uuid.UUID `gorm:"type:UUID;default:generateUUIDv4();primaryKey" json:"id"`
	PlantID         int64     `gorm:"column:plant_id;comment:'种植记录ID'" json:"plant_id"`
	UserID          int64     `gorm:"column:user_id;comment:'用户ID'" json:"user_id"`
	FarmID          int64     `gorm:"column:farm_id;comment:'农场ID'" json:"farm_id"`
	LandID          int64     `gorm:"column:land_id;comment:'土地ID'" json:"land_id"`
	CropName        string    `gorm:"column:crop_name;type:LowCardinality(String);comment:'作物名称'" json:"crop_name"`
	PlantTime       int64     `gorm:"column:plant_time;comment:'种植时间（Unix时间戳）'" json:"plant_time"`
	WaterCount      int64     `gorm:"column:water_count;default:0;comment:'已浇水总次数'" json:"water_count"`
	FertilizerCount int64     `gorm:"column:fertilizer_count;default:0;comment:'已施肥总次数'" json:"fertilizer_count"`
	IsHarvested     int64     `gorm:"column:is_harvested;default:1;comment:'是否已收获：1-未收获，2-已收获'" json:"is_harvested"`
	HarvestedAt     int64     `gorm:"column:harvested_at;comment:'收获时间（Unix时间戳）'" json:"harvested_at"`
	CoinsEarned     int64     `gorm:"column:coins_earned;comment:'作物产生的金币数量'" json:"coins_earned"`
	Status          int64     `gorm:"column:status;default:1;comment:'生长状态：1-种子阶段，2-发芽阶段，3-成长阶段，4-开花阶段，5-结果阶段，6-成熟阶段'" json:"status"`
	//
	gormx.Model
}

// TableName specifies the table name for the model
func (*UserPlanting) TableName() string {
	return TableNameUserPlanting
}
