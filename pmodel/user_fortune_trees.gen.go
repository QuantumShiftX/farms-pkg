package pmodel

import "github.com/QuantumShiftX/golib/stores/gormx"

const TableNameUserFortuneTree = "user_fortune_trees"

// UserFortuneTree 发财树表（每位用户一棵树，用于定时生产金币并在收获后重新开始计时）
type UserFortuneTree struct {
	ID                 int64 `gorm:"column:id;primaryKey;comment:发财树唯一标识ID"`
	UserID             int64 `gorm:"column:user_id;not null;uniqueIndex;comment:用户ID，唯一约束保证每个用户只有一棵树"`
	VIPLevelID         int64 `gorm:"column:vip_level_id;not null;default:1;index;comment:VIP等级ID，影响金币产出和生产周期"`
	CycleStartTime     int64 `gorm:"column:cycle_start_time;comment:当前周期开始时间（时间戳，秒）"`
	IsReadyToHarvest   int64 `gorm:"column:is_ready_to_harvest;not null;default:1;comment:是否可以收获标志：1-不可收获，2-可收获"`
	LastHarvestTime    int64 `gorm:"column:last_harvest_time;comment:上次收获时间（时间戳，秒）"`
	NextHarvestTime    int64 `gorm:"column:next_harvest_time;index;comment:下次可收获时间（时间戳，秒）"`
	CycleHarvestedIcon int64 `gorm:"column:cycle_harvested_icon;not null;default:0;comment:当前期收获金币数量"`
	TotalHarvested     int64 `gorm:"column:total_harvested;not null;default:0;comment:历史总收获金币数量统计"`
	//
	gormx.Model
}

// TableName 设置表名
func (*UserFortuneTree) TableName() string {
	return TableNameUserFortuneTree
}
