package pmodel

import "github.com/QuantumShiftX/golib/stores/gormx"

const TableNameVipUpgradeHistory = "vip_upgrade_histories"

// VipUpgradeHistory 用户VIP等级升级历史记录表
type VipUpgradeHistory struct {
	ID                int64  `gorm:"column:id;primaryKey;autoIncrement:true;comment:自增主键ID" json:"id"`              // 自增主键ID
	UserID            int64  `gorm:"column:user_id;not null;index;comment:用户ID" json:"user_id"`                     // 用户ID
	OldVipLevelID     int64  `gorm:"column:old_vip_level_id;not null;comment:升级前VIP等级ID" json:"old_vip_level_id"`   // 升级前VIP等级ID
	NewVipLevelID     int64  `gorm:"column:new_vip_level_id;not null;comment:升级后VIP等级ID" json:"new_vip_level_id"`   // 升级后VIP等级ID
	GrowthValueBefore int64  `gorm:"column:growth_value_before;not null;comment:升级前成长值" json:"growth_value_before"` // 升级前成长值
	GrowthValueAfter  int64  `gorm:"column:growth_value_after;not null;comment:升级后成长值" json:"growth_value_after"`   // 升级后成长值
	GrowthValueDiff   int64  `gorm:"column:growth_value_diff;not null;comment:本次升级增加的成长值" json:"growth_value_diff"` // 本次升级增加的成长值
	Remark            string `gorm:"column:remark;default:'';comment:升级备注说明" json:"remark"`                         // 升级备注说明

	// 通用时间字段
	gormx.Model
}

// TableName 指定表名
func (*VipUpgradeHistory) TableName() string {
	return TableNameVipUpgradeHistory
}
