package pmodel

import "github.com/QuantumShiftX/golib/stores/gormx"

const TableNamePaymentConfig = "payment_configs"

// PaymentConfig 支付配置模型，对应数据库表 payment_configs
type PaymentConfig struct {
	ID                 int64  `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	PaymentConfigType  int64  `gorm:"column:payment_config_type;not null" json:"payment_config_type"`               // 配置类型: 1:充值(deposit), 2:提现(withdraw)
	PaymentChannelType int64  `gorm:"column:payment_channel_type;not null" json:"payment_channel_type"`             // 支付渠道类型: 1:USDT虚拟货币, 2:银行卡, 3:快捷支付
	Name               string `gorm:"column:name;type:varchar(500);not null" json:"name"`                           // 支付方式名称，如"USDT充值"、"银行卡提现"等
	CurrencyCode       string `gorm:"column:currency_code;type:varchar(500);not null" json:"currency_code"`         // 货币代码，如"USDT"、"CNY"等
	LogoURL            string `gorm:"column:logo_url;type:varchar(500)" json:"logo_url"`                            // LOGO图片URL地址，用于前端显示
	HasExchangeRate    int64  `gorm:"column:has_exchange_rate;not null;default:2" json:"has_exchange_rate"`         // 是否有汇率: 1:有, 2:无
	ExchangeRate       int64  `gorm:"column:exchange_rate;not null;default:0" json:"exchange_rate"`                 // 汇率值，1 USDT 兑换目标货币的汇率 1usdt兑换0.998美元 那么exchange_rate=0.998*1000000
	FixedFee           int64  `gorm:"column:fixed_fee;not null;default:0" json:"fixed_fee"`                         // 固定手续费金额，每笔交易收取的固定金额
	PercentageFee      int64  `gorm:"column:percentage_fee;not null;default:0" json:"percentage_fee"`               // 百分比手续费率，以1000000为单位，例如 20%=0.20=0.20×1,000,000=200,000
	Status             int64  `gorm:"column:status;not null;default:1" json:"status"`                               // 状态: 1:启用, 2:禁用
	Sort               int64  `gorm:"column:sort;sort:desc;comment:排序" json:"sort"`                                 // 排序
	MinLimitAmount     int64  `gorm:"column:min_limit_amount;not null;comment:'单笔最小限额，下限'" json:"min_limit_amount"` // 单笔最小限额
	MaxLimitAmount     int64  `gorm:"column:max_limit_amount;not null;comment:'单笔最大限额，上限'" json:"max_limit_amount"` // 单笔最大限额
	//
	gormx.OperationBaseModel
	gormx.Model
}

// TableName 指定模型对应的数据库表名
func (*PaymentConfig) TableName() string {
	return TableNamePaymentConfig
}
