package pmodel

import "github.com/QuantumShiftX/golib/stores/gormx"

// 银行配置表
const TableNameBankConfig = "bank_configs"

// BankConfig mapped from table <bank_configs>
type BankConfig struct {
	ID             int64  `gorm:"column:id;primaryKey;autoIncrement:true;comment:主键" json:"id"`                                                                          // 主键
	CountryCode    string `gorm:"column:country_code;type:varchar(32);not null;index:idx_bank_configs_country_code;comment:国家/地区代码" json:"country_code"`             // 国家/地区代码
	CurrencyCode   string `gorm:"column:currency_code;type:varchar(32);not null;index:idx_bank_configs_currency_code_status,priority:1;comment:币种" json:"currency_code"` // 币种
	CollectionType string `gorm:"column:collection_type;type:varchar(64);not null;comment:收款类型，如bank/pix等" json:"collection_type"`                                   // 收款类型
	BankName       string `gorm:"column:bank_name;type:varchar(255);not null;comment:银行名称" json:"bank_name"`                                                           // 银行名称
	BankShowName   string `gorm:"column:bank_show_name;type:varchar(255);not null;comment:银行显示名称" json:"bank_show_name"`                                             // 银行显示名称
	BankIcon       string `gorm:"column:bank_icon;type:varchar(255);not null;comment:银行图标" json:"bank_icon"`                                                           // 银行图标
	Code           string `gorm:"column:code;type:varchar(255);comment:银行编码" json:"code"`                                                                              // 银行编码
	Fee            int64  `gorm:"column:fee;not null;default:0;comment:手续费" json:"fee"`                                                                                 // 手续费
	Description    string `gorm:"column:description;type:text;comment:银行描述" json:"description"`                                                                        // 银行描述
	Status         int8   `gorm:"column:status;type:smallint;not null;index:idx_bank_configs_currency_code_status,priority:2;comment:启用状态" json:"status"`              // 启用状态
	TopUpAt        int64  `gorm:"column:top_up_at;index:idx_bank_configs_top_up_at,priority:1,sort:desc;comment:置顶时间" json:"top_up_at"`                                // 置顶时间

	gormx.OperationBaseModel
	gormx.Model
}

// TableName BankConfig's table name
func (*BankConfig) TableName() string {
	return TableNameBankConfig
}
