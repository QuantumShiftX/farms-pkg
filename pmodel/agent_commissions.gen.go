package pmodel

import "github.com/QuantumShiftX/golib/stores/gormx"

const TableNameAgentCommission = "agent_commissions"

// AgentCommission 代理佣金记录表
type AgentCommission struct {
	gormx.Model
	//
	ID           int64  `gorm:"column:id;primaryKey;autoIncrement" json:"id"`         // 主键ID
	OrderNo      string `gorm:"column:order_no;not null;index" json:"order_no"`       // 佣金订单号
	AgentID      int64  `gorm:"column:agent_id;not null;index" json:"agent_id"`       // 代理ID
	WalletType   int64  `gorm:"column:wallet_type;not null;index" json:"wallet_type"` // 钱包类型
	SourceType   int64  `gorm:"column:source_type;not null;index" json:"source_type"` // 来源类型：1=下级存款, 2=下级收获
	UserID       int64  `gorm:"column:user_id;index" json:"user_id"`                  // 关联会员ID
	Amount       int64  `gorm:"column:amount;not null" json:"amount"`                 // 佣金金额（单位：微）
	Rate         int64  `gorm:"column:rate;not null" json:"rate"`                     // 佣金比率（乘以1000000，例如：50%存为500000）
	BalanceAfter int64  `gorm:"column:balance_after;not null" json:"balance_after"`   // 变动后余额（单位：微）
	Remark       string `gorm:"column:remark" json:"remark"`                          // 备注信息
	Agent        *Agent `gorm:"foreignKey:AgentID;references:AgentID" json:"agent"`   // 关联代理

}

// TableName 指定表名
func (*AgentCommission) TableName() string {
	return TableNameAgentCommission
}
