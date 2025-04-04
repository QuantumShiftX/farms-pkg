package pmodel

import "github.com/QuantumShiftX/golib/stores/gormx"

const TableNameWalletTransaction = "wallet_transactions"

// WalletTransaction 钱包交易记录表 - 记录钱包的充值和提现记录
type WalletTransaction struct {
	ID              int64  `gorm:"column:id;primaryKey;autoIncrement" json:"id"`                   // 主键ID，自增
	WalletID        int64  `gorm:"column:wallet_id;not null;index" json:"wallet_id"`               // 关联的钱包ID
	TransactionType int8   `gorm:"column:transaction_type;not null;index" json:"transaction_type"` // 交易类型：1-入款、2-出款、3-归集
	AmountUsdt      int64  `gorm:"column:amount_usdt;default:0" json:"amount_usdt"`                // USDT金额
	AmountToken     int64  `gorm:"column:amount_token;default:0" json:"amount_token"`              // 平台币金额
	TransactionHash string `gorm:"column:transaction_hash;size:255;index" json:"transaction_hash"` // 交易哈希
	FromAddress     string `gorm:"column:from_address;size:255" json:"from_address"`               // 来源地址
	ToAddress       string `gorm:"column:to_address;size:255" json:"to_address"`                   // 目标地址
	Status          int8   `gorm:"column:status;default:1;index" json:"status"`                    // 状态：1-成功，2-失败，3-处理中
	Remarks         string `gorm:"column:remarks;type:text" json:"remarks"`                        // 备注
	//
	gormx.Model
}

// TableName 设置WalletTransaction的表名
func (*WalletTransaction) TableName() string {
	return TableNameWalletTransaction
}
