// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package pmodel

import (
	"github.com/QuantumShiftX/golib/stores/gormx"
	"gorm.io/datatypes"
)

const TableNameCurrencyInfo = "currency_infos"

// CurrencyInfo mapped from table <currency_infos>
type CurrencyInfo struct {
	ID                int64                       `gorm:"column:id;primaryKey;autoIncrement:true;comment:自增主键，唯一标识每条记录" json:"id"`                                                 // 自增主键，唯一标识每条记录
	Rank              int64                       `gorm:"column:rank;not null;comment:排名，表示货币在列表中的排序位置" json:"rank"`                                                               // 排名，表示货币在列表中的排序位置
	CurrencyName      string                      `gorm:"column:currency_name;not null;comment:货币名称，例如“印度卢比”" json:"currency_name"`                                                // 货币名称，例如“印度卢比”
	CountryName       string                      `gorm:"column:country_name;comment:所属国家，例如“印度”，允许为空（如“波场”）" json:"country_name"`                                                 // 所属国家，例如“印度”，允许为空（如“波场”）
	CurrencyCode      string                      `gorm:"column:currency_code;not null;uniqueIndex:idx_unique_currency_code,priority:1;comment:货币代码，例如“INR”" json:"currency_code"` // 货币代码，例如“INR”
	CurrencyFlagURL   string                      `gorm:"column:currency_flag_url;comment:货币国旗图标的 URL，保存图片的路径或链接" json:"currency_flag_url"`                                        // 货币国旗图标的 URL，保存图片的路径或链接
	ThousandSeparator string                      `gorm:"column:thousand_separator;comment:千分位符号，例如“,”，表示格式化显示" json:"thousand_separator"`                                         // 千分位符号，例如“,”，表示格式化显示
	CurrencySymbol    string                      `gorm:"column:currency_symbol;comment:货币符号，例如“₹”" json:"currency_symbol"`                                                        // 货币符号，例如“₹”
	ExchangeRate      string                      `gorm:"column:exchange_rate;not null;default:1:1;comment:货币兑换比例，默认值为 1:1" json:"exchange_rate"`                                  // 货币兑换比例，默认值为 1:1
	CurrencyType      int64                       `gorm:"column:currency_type;not null;default:1;comment:货币类型：1-法定货币，2-数字货币，3-平台币" json:"currency_type"`                           // 货币类型：1-法定货币，2-数字货币，3-平台币
	IsEnabled         int8                        `gorm:"column:is_enabled;not null;default:1;comment:是否启用，默认启用" json:"is_enabled"`                                                // 是否启用，默认启用
	Protocol          datatypes.JSONSlice[string] `gorm:"column:protocol;comment: 虚拟货币支持的协议，例如USDT支持['TRC-20','ERC-20'，'BEP-20']" json:"protocol"`                                 //  虚拟货币支持的协议，例如USDT支持["TRC-20","ERC-20"，"BEP-20"]
	//
	gormx.OperationBaseModel
	gormx.Model
}

// TableName CurrencyInfo's table name
func (*CurrencyInfo) TableName() string {
	return TableNameCurrencyInfo
}
