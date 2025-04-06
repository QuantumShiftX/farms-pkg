package pmodel

import "github.com/QuantumShiftX/golib/stores/gormx"

const TableNameExchangeRecord = "exchange_records"

// ExchangeRecord 用户换汇记录表
type ExchangeRecord struct {
	ID                int64  `gorm:"primaryKey;autoIncrement;column:id"`  // 记录ID，主键，自动增长
	UserID            int64  `gorm:"column:user_id;not null"`             // 用户ID
	Username          string `gorm:"column:username;not null"`            // 用户账号
	OrderNo           string `gorm:"column:order_no;not null;unique"`     // 订单编号，唯一标识一笔交易
	DigitalMerchantID int64  `gorm:"column:digital_merchant_id;not null"` // 币商 id
	ExchangePair      string `gorm:"column:exchange_pair;not null"`       // 兑换货币对，格式为"代币->USDT"
	ExchangeRate      int64  `gorm:"column:exchange_rate;not null"`       // 汇率
	PaymentAmount     int64  `gorm:"column:payment_amount;not null"`      // 支付金额
	ReceivedAmount    int64  `gorm:"column:received_amount;not null"`     // 获得金额
	OrderStatus       int16  `gorm:"column:order_status;not null"`        // 订单状态：1-处理中 2-处理成功 3-处理失败
	//
	gormx.OperationBaseModel
	gormx.Model
}

// TableName 设置数据库表名
func (*ExchangeRecord) TableName() string {
	return TableNameExchangeRecord
}
