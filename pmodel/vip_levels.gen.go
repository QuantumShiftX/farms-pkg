// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package pmodel

import "github.com/QuantumShiftX/golib/stores/gormx"

const TableNameVipLevel = "vip_levels"

// VipLevel mapped from table <vip_levels>
type VipLevel struct {
	ID                              int64  `gorm:"column:id;primaryKey;autoIncrement:true;comment:自增主键，唯一标识每个VIP等级" json:"id"`                                                         // 自增主键，唯一标识每个VIP等级
	LevelName                       string `gorm:"column:level_name;not null;comment:VIP等级的名称" json:"level_name"`                                                                      // VIP等级的名称
	LevelVal                        int64  `gorm:"column:level_val;not null;default:0;comment:VIP等级的值" json:"level_val"`                                                               // VIP等级的值
	IsDefaultLevel                  int64  `gorm:"column:is_default_level;not null;default:2;comment:是否为默认VIP等级（1表示是，2表示否）" json:"is_default_level"`                                   // 是否为默认VIP等级（1表示是，2表示否）
	GrowthValue                     int64  `gorm:"column:growth_value;not null;comment:成长值达到此点即可升级为该VIP等级" json:"growth_value"`                                                        // 成长值达到此点即可升级为该VIP等级
	UpgradeRewardFertilizerNum      int64  `gorm:"column:upgrade_reward_fertilizer_num;not null;default:1;comment:升级时获得的肥料数量" json:"upgrade_reward_fertilizer_num"`                    // 升级时获得的肥料数量
	UpgradeRewardTreeMaturityPeriod int32  `gorm:"column:upgrade_reward_tree_maturity_period;not null;default:24;comment:发财树奖励的成熟周期，单位为小时" json:"upgrade_reward_tree_maturity_period"` // 发财树奖励的成熟周期，单位为小时
	UpgradeRewardTreeNum            int64  `gorm:"column:upgrade_reward_tree_num;not null;comment:升级时获得的发财树代币数量" json:"upgrade_reward_tree_num"`                                       // 升级时获得的发财树代币数量
	Remarks                         string `gorm:"column:remarks;comment:对VIP等级的备注说明" json:"remarks"`                                                                                  // 对VIP等级的备注说明
	IsEnabled                       int8   `gorm:"column:is_enabled;not null;default:1;comment:是否启用该VIP等级（1表示启用，2表示禁用）" json:"is_enabled"`                                             // 是否启用该VIP等级（1表示启用，2表示禁用）
	//
	gormx.OperationBaseModel
	gormx.Model
}

// TableName VipLevel's table name
func (*VipLevel) TableName() string {
	return TableNameVipLevel
}
