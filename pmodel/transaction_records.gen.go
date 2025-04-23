package pmodel

import "github.com/QuantumShiftX/golib/stores/gormx"

const TableNameTransactionRecord = "transaction_records"

// TransactionRecord 账变记录表
type TransactionRecord struct {
	ID              int64  `gorm:"primaryKey;autoIncrement;column:id" json:"id"`             // 主键，自动增长
	UserID          int64  `gorm:"column:user_id;not null" json:"user_id"`                   // 用户ID
	Username        string `gorm:"column:username;not null" json:"username"`                 // 用户账号
	OrderNo         string `gorm:"column:order_no;not null" json:"order_no"`                 // 流水编号
	CurrencyCode    string `gorm:"column:currency_code;not null" json:"currency_code"`       // 币种编码
	WalletType      int16  `gorm:"column:wallet_type;not null" json:"wallet_type"`           // 账变钱包类型（1:用户钱包 2:代理钱包 3:奖励钱包 4:USTD钱包）
	Category        int16  `gorm:"column:category;not null" json:"category"`                 // 账变大类交易类型
	SubCategory     int16  `gorm:"column:sub_category" json:"sub_category"`                  // 账变小类
	BalancePrevious int64  `gorm:"column:balance_previous;not null" json:"balance_previous"` // 变动前余额, 单位:微
	Amount          int64  `gorm:"column:amount;not null" json:"amount"`                     // 变动金额, 单位:微
	BalanceAfter    int64  `gorm:"column:balance_after;not null" json:"balance_after"`       // 变动后余额, 单位:微
	Remark          string `gorm:"column:remark" json:"remark"`                              // 备注信息
	TransactionTime int64  `gorm:"column:transaction_time;not null" json:"transaction_time"` // 交易时间
	RelatedID       string `gorm:"column:related_id;not null" json:"related_id"`             // 关联ID，如游戏ID，商品ID，活动ID
	MsgID           string `gorm:"column:msg_id" json:"msg_id"`                              // 消息队列消息ID
	//
	gormx.Model
}

// TableName 设置数据库表名
func (*TransactionRecord) TableName() string {
	return TableNameTransactionRecord
}
