package pmodel

import "github.com/QuantumShiftX/golib/stores/gormx"

const TableNameDigitalCurrencyMerchant = "digital_currency_merchants"

// DigitalCurrencyMerchant 币商配置表
type DigitalCurrencyMerchant struct {
	ID                  int64  `gorm:"primaryKey;column:id;type:int8;autoIncrement:true" json:"id"`                            // 自增主键
	Name                string `gorm:"column:name;type:varchar(500);not null" json:"name"`                                     // 币商名称
	Logo                string `gorm:"column:logo;type:varchar(500);not null" json:"logo"`                                     // 币商LOGO地址
	Protocol            string `gorm:"column:protocol;type:varchar(500);not null" json:"protocol"`                             // 币商协议
	ExchangeRate        int64  `gorm:"column:exchange_rate;type:int8;not null" json:"exchange_rate"`                           // 基准汇率(代币兑换USDT)类似1个代币兑换多少ustd的汇率
	UpwardFluctuation   int64  `gorm:"column:upward_fluctuation;type:int8;not null;default:0" json:"upward_fluctuation"`       // 上浮幅度（正向波动范围）
	DownwardFluctuation int64  `gorm:"column:downward_fluctuation;type:int8;not null;default:0" json:"downward_fluctuation"`   // 下浮幅度（负向波动范围）
	MinCollectionAmount int64  `gorm:"column:min_collection_amount;type:int8;not null;default:0" json:"min_collection_amount"` // 最小收购金额(代币)
	MaxCollectionAmount int64  `gorm:"column:max_collection_amount;type:int8;not null;default:0" json:"max_collection_amount"` // 最大收购金额(代币)
	PurchaseAmount      int64  `gorm:"column:purchase_amount;type:int8;not null;default:0" json:"purchase_amount"`             // 已收购金额(代币)
	OrderAmount         int64  `gorm:"column:order_amount;type:int8;not null;default:0" json:"order_amount"`                   // 已订购金额(USDT)
	Balance             int64  `gorm:"column:balance;type:int8;not null;default:0" json:"balance"`                             // 库存余额(USDT)
	Sort                int64  `gorm:"column:sort;type:int8;not null;default:0" json:"sort"`                                   // 排序
	IsUnlimitedBalance  int8   `gorm:"column:is_unlimited_balance;type:int2;not null;default:2" json:"is_unlimited_balance"`   // 是否无限库存余额 1 是，2：否
	IsOnline            int8   `gorm:"column:is_online;type:int2;not null;default:1" json:"is_online"`                         // 是否上架 1开启上架，2：下架
	RealExchangeRate    int64  `gorm:"-" json:"real_exchange_rate"`                                                            // 当前的实时汇率(代币兑换USDT)类似1个代币兑换多少ustd的汇率
	//
	gormx.OperationBaseModel
	gormx.Model
}

// TableName 表名
func (m *DigitalCurrencyMerchant) TableName() string {
	return TableNameDigitalCurrencyMerchant
}
