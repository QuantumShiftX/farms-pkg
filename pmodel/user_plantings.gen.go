package pmodel

import (
	"github.com/QuantumShiftX/golib/stores/gormx"
)

const TableNameUserPlanting = "user_plantings"

// UserPlanting represents the user_plantings table in ClickHouse
type UserPlanting struct {
	ID              int64 `gorm:"column:id;primaryKey" json:"id"`
	PlantID         int64 `gorm:"column:plant_id;comment:'种植记录ID'" json:"plant_id"`
	UserID          int64 `gorm:"column:user_id;comment:'用户ID'" json:"user_id"`
	FarmID          int64 `gorm:"column:farm_id;comment:'农场ID'" json:"farm_id"`
	LandID          int64 `gorm:"column:land_id;comment:'土地ID'" json:"land_id"`
	StorageID       int64 `gorm:"column:storage_id;comment:'用户的仓库物品ID'" json:"storage_id"`
	PlantTime       int64 `gorm:"column:plant_time;comment:'种植时间（Unix时间戳）'" json:"plant_time"`
	WaterCount      int64 `gorm:"column:water_count;default:0;comment:'已浇水总次数'" json:"water_count"`
	FertilizerCount int64 `gorm:"column:fertilizer_count;default:0;comment:'已施肥总次数'" json:"fertilizer_count"`
	IsHarvested     int64 `gorm:"column:is_harvested;default:1;comment:'是否已收获：1、未收获(已种植) 2、可收获 3、已收获'" json:"is_harvested"`
	HarvestedAt     int64 `gorm:"column:harvested_at;comment:'收获时间（Unix时间戳）'" json:"harvested_at"`
	Status          int64 `gorm:"column:status;default:1;comment:'生长状态：1-种子阶段，2-发芽阶段，3-成长阶段，4-开花阶段，5-结果阶段，6-成熟阶段'" json:"status"`
	// 快照字段
	CropID                            int64  `gorm:"column:crop_id;comment:'作物ID（对应商店信息中的sku_id或者种子信息中的id）'" json:"crop_id"`
	CropName                          string `gorm:"column:crop_name;type:LowCardinality(String);comment:'种植时作物的作物名称快照'" json:"crop_name"`
	ProductName                       string `gorm:"column:product_name;type:LowCardinality(String);comment:'种植时商品名称快照'" json:"product_name"`
	Price                             int64  `gorm:"column:price;comment:'种植时作物的价格快照'" json:"price"`
	GrowthTime                        int64  `gorm:"column:growth_time;comment:'种植时作物的生长时间快照'" json:"growth_time"`
	CoinReward                        int64  `gorm:"column:coin_reward;comment:'种植时作物的金币奖励快照'" json:"coin_reward"`
	WateringFrequency                 int64  `gorm:"column:watering_frequency;comment:'种植时作物的浇水频率快照'" json:"watering_frequency"`
	FertilizingFrequency              int64  `gorm:"column:fertilizing_frequency;comment:'种植时作物的施肥频率快照'" json:"fertilizing_frequency"`
	SelfTimeReducedPerWater           int64  `gorm:"column:self_time_reduced_per_water;comment:'种植时每次自己浇水减少的时间快照'" json:"self_time_reduced_per_water"`
	SelfWaterReductionPercentage      int64  `gorm:"column:self_water_reduction_percentage;comment:'种植时自己浇水减少时间的百分比快照'" json:"self_water_reduction_percentage"`
	SelfTimeReducedPerFertilizer      int64  `gorm:"column:self_time_reduced_per_fertilizer;comment:'种植时每次自己施肥减少的时间快照'" json:"self_time_reduced_per_fertilizer"`
	SelfFertilizerReductionPercentage int64  `gorm:"column:self_fertilizer_reduction_percentage;comment:'种植时自己施肥减少时间的百分比快照'" json:"self_fertilizer_reduction_percentage"`
	FriendWaterTimeReduction          int64  `gorm:"column:friend_water_time_reduction;comment:'种植时好友浇水减少的时间快照'" json:"friend_water_time_reduction"`
	FriendWaterPercentage             int64  `gorm:"column:friend_water_percentage;comment:'种植时好友浇水百分比快照'" json:"friend_water_percentage"`
	FriendFertilizerTimeReduction     int64  `gorm:"column:friend_fertilizer_time_reduction;comment:'种植时好友施肥减少的时间快照'" json:"friend_fertilizer_time_reduction"`
	FriendFertilizerPercentage        int64  `gorm:"column:friend_fertilizer_percentage;comment:'种植时好友施肥百分比快照'" json:"friend_fertilizer_percentage"`
	//
	gormx.Model
}

// TableName specifies the table name for the model
func (*UserPlanting) TableName() string {
	return TableNameUserPlanting
}
