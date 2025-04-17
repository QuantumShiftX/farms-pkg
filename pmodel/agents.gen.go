package pmodel

import "github.com/QuantumShiftX/golib/stores/gormx"

const TableNameAgent = "agents"

// Agent 代理商信息表
type Agent struct {
	gormx.Model
	//
	Status                 int64 `gorm:"column:status;default:1" json:"status"`                                     // 状态：1表示正常，2表示禁用
	ID                     int64 `gorm:"column:id;primaryKey;autoIncrement" json:"id"`                              // 主键ID
	AgentID                int64 `gorm:"column:agent_id;not null;uniqueIndex" json:"agent_id"`                      // 代理ID，与用户ID一致
	ParentID               int64 `gorm:"column:parent_id" json:"parent_id"`                                         // 上级代理ID
	DirectChildrenCount    int   `gorm:"column:direct_children_count;default:0" json:"direct_children_count"`       // 直属下级代理人数
	DirectMembersCount     int   `gorm:"column:direct_members_count;default:0" json:"direct_members_count"`         // 直属会员数
	DirectDepositAmount    int   `gorm:"column:direct_deposit_amount;default:0" json:"direct_deposit_amount"`       // 直接存款总额（单位：微）
	DirectWithdrawalAmount int   `gorm:"column:direct_withdrawal_amount;default:0" json:"direct_withdrawal_amount"` // 直接提款总额（单位：微）
	DirectCommissionAmount int   `gorm:"column:direct_commission_amount;default:0" json:"direct_commission_amount"` // 直接佣金总额（单位：微）
	TotalChildrenCount     int   `gorm:"column:total_children_count;default:0" json:"total_children_count"`         // 团队总人数（包括所有层级下级）
	TotalMembersCount      int   `gorm:"column:total_members_count;default:0" json:"total_members_count"`           // 团队会员总数（包括所有层级下级的会员）
	TotalDepositAmount     int   `gorm:"column:total_deposit_amount;default:0" json:"total_deposit_amount"`         // 团队存款总额（单位：微）
	TotalWithdrawalAmount  int   `gorm:"column:total_withdrawal_amount;default:0" json:"total_withdrawal_amount"`   // 团队提款总额（单位：微）
	TotalCommissionAmount  int   `gorm:"column:total_commission_amount;default:0" json:"total_commission_amount"`   // 团队佣金总额（单位：微）
}

// TableName 指定表名
func (*Agent) TableName() string {
	return TableNameAgent
}
