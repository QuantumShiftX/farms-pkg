package cmodel

import (
	"github.com/QuantumShiftX/golib/stores/gormx"
)

const TableNameUserPlanting = "user_plantings"

// UserPlanting represents the user_plantings table in ClickHouse
type UserPlanting struct {
	ID              string `gorm:"column:id;type:UUID;default:generateUUIDv4();primaryKey"`
	PlantID         int64  `gorm:"column:plant_id;comment:'种植记录ID'"`
	UserID          int64  `gorm:"column:user_id;comment:'用户ID'"`
	FarmID          int64  `gorm:"column:farm_id;comment:'农场ID'"`
	LandID          int64  `gorm:"column:land_id;comment:'土地ID'"`
	StorageID       int64  `gorm:"column:storage_id;comment:'用户的仓库物品ID'"`
	PlantTime       int64  `gorm:"column:plant_time;comment:'种植时间（Unix时间戳）'"`
	WaterCount      int64  `gorm:"column:water_count;default:0;comment:'已浇水总次数'"`
	FertilizerCount int64  `gorm:"column:fertilizer_count;default:0;comment:'已施肥总次数'"`
	IsHarvested     int8   `gorm:"column:is_harvested;default:1;comment:'是否已收获：1-未收获，2-已收获'"`
	HarvestedAt     int64  `gorm:"column:harvested_at;comment:'收获时间（Unix时间戳）'"`
	Status          int64  `gorm:"column:status;default:1;comment:'生长状态：1-种子阶段，2-发芽阶段，3-成长阶段，4-开花阶段，5-结果阶段，6-成熟阶段'"`

	// 快照字段
	ProductName                       string `gorm:"column:product_name;type:LowCardinality(String);comment:'种植时商品名称快照'"`
	CropName                          string `gorm:"column:crop_name;type:LowCardinality(String);comment:'种植时作物的作物名称快照'"`
	Price                             int64  `gorm:"column:price;comment:'种植时作物的价格快照'"`
	GrowthTime                        int64  `gorm:"column:growth_time;comment:'种植时作物的生长时间快照'"`
	CoinReward                        int64  `gorm:"column:coin_reward;comment:'种植时作物的金币奖励快照'"`
	WateringFrequency                 int64  `gorm:"column:watering_frequency;comment:'种植时作物的浇水频率快照'"`
	FertilizingFrequency              int64  `gorm:"column:fertilizing_frequency;comment:'种植时作物的施肥频率快照'"`
	SelfTimeReducedPerWater           int64  `gorm:"column:self_time_reduced_per_water;comment:'种植时每次自己浇水减少的时间快照'"`
	SelfWaterReductionPercentage      int64  `gorm:"column:self_water_reduction_percentage;comment:'种植时自己浇水减少时间的百分比快照'"`
	SelfTimeReducedPerFertilizer      int64  `gorm:"column:self_time_reduced_per_fertilizer;comment:'种植时每次自己施肥减少的时间快照'"`
	SelfFertilizerReductionPercentage int64  `gorm:"column:self_fertilizer_reduction_percentage;comment:'种植时自己施肥减少时间的百分比快照'"`
	FriendWaterTimeReduction          int64  `gorm:"column:friend_water_time_reduction;comment:'种植时好友浇水减少的时间快照'"`
	FriendWaterPercentage             int64  `gorm:"column:friend_water_percentage;comment:'种植时好友浇水百分比快照'"`
	FriendFertilizerTimeReduction     int64  `gorm:"column:friend_fertilizer_time_reduction;comment:'种植时好友施肥减少的时间快照'"`
	FriendFertilizerPercentage        int64  `gorm:"column:friend_fertilizer_percentage;comment:'种植时好友施肥百分比快照'"`

	//
	gormx.Model
}

// TableName specifies the table name for the model
func (*UserPlanting) TableName() string {
	return TableNameUserPlanting
}
