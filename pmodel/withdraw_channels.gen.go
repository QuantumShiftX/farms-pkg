package pmodel

import (
	"github.com/QuantumShiftX/golib/stores/gormx"
)

const TableNameWithdrawChannel = "withdraw_channels"

// WithdrawChannel 表示提现渠道子表
type WithdrawChannel struct {
	gormx.OperationBaseModel
	gormx.Model
	Status int64 `gorm:"not null;default:1;comment:'通道状态：1-启用，2-禁用';index:idx_withdraw_channels_status"` // 通道启用状态
	//
	ID                    int64  `gorm:"primaryKey;autoIncrement;comment:'主键ID'"`                                                                    // 主键ID
	PaymentConfigID       int64  `gorm:"not null;comment:'支付配置ID，关联到payment_configs表的主键';index:idx_withdraw_channels_payment_config_id"`    // 支付配置ID
	CurrencyCode          string `gorm:"type:varchar(255);not null;comment:'币种代码，如\"USDT\"、\"CNY\"等';index:idx_withdraw_channels_currency"`      // 币种代码
	WithdrawChannelName   string `gorm:"type:varchar(255);not null;comment:'提现通道名称，用于前端展示'"`                                               // 提现通道名称
	WithdrawChannelCode   string `gorm:"type:varchar(255);not null;comment:'提现通道编码，系统内部唯一标识';index:idx_withdraw_channels_code_currency"` // 提现通道编码
	WithdrawChannelType   int16  `gorm:"not null;comment:'提现通道类型：1-USDT虚拟货币，2-银行卡，3-快捷支付';index:idx_withdraw_channels_type"`          // 提现通道类型
	WithdrawChannelLogo   string `gorm:"type:varchar(500);not null;comment:'提现通道logo图片URL'"`                                                     // 提现通道logo
	WithdrawMerchantID    int64  `gorm:"not null;comment:'提现商户配置ID，关联到支付商户表';index:idx_withdraw_channels_merchant_id"`                   // 提现商户配置ID
	BankID                int64  `gorm:"comment:'银行ID，关联到银行信息表';index:idx_withdraw_channels_bank_id"`                                        // 银行ID
	BankName              string `gorm:"type:varchar(255);comment:'银行名称，如\"中国工商银行\"'"`                                                      // 银行名称
	BankBranch            string `gorm:"type:varchar(255);comment:'银行支行名称，如\"北京市海淀区支行\"'"`                                              // 银行支行
	BankAccount           string `gorm:"type:varchar(255);comment:'银行账号，收款账户号码'"`                                                            // 银行账号
	BankAccountName       string `gorm:"type:varchar(255);comment:'账户名称，银行卡开户人姓名'"`                                                        // 账户名称
	CryptoAddress         string `gorm:"type:varchar(500);comment:'加密货币钱包地址，用于收款的区块链地址'"`                                            // 加密货币地址
	CryptoNetwork         string `gorm:"type:varchar(50);comment:'网络类型，如\"TRC20\"、\"ERC20\"等区块链网络'"`                                        // 网络类型
	MinLimitAmount        int64  `gorm:"not null;comment:'单笔最小限额，提现金额下限'"`                                                                 // 单笔最小限额
	MaxLimitAmount        int64  `gorm:"not null;comment:'单笔最大限额，提现金额上限'"`                                                                 // 单笔最大限额
	DailyLimitAmount      int64  `gorm:"default:0;comment:'每日提现总额限制，0表示不限制'"`                                                             // 每日限额
	FeePercent            int64  `gorm:"not null;default:0;comment:'手续费百分比'"`                                                                    // 手续费（百分比）
	FixedFee              int64  `gorm:"default:0;comment:'固定手续费金额，每笔提现固定收取的费用'"`                                                    // 固定手续费
	FeeReductionConfig    string `gorm:"type:jsonb;default:'{}';comment:'手续费减免配置，JSON格式，可配置不同条件下的手续费减免规则'"`                   // 手续费减免配置
	RecommendAmountConfig string `gorm:"type:jsonb;default:'{}';comment:'推荐金额配置，JSON格式，前端展示的快捷金额选项'"`                               // 推荐金额配置
	Hint                  string `gorm:"type:text;comment:'提示信息，展示给用户的操作指引，最大1000字'"`                                                 // 提示
	Remark                string `gorm:"type:text;comment:'备注信息，内部管理用'"`                                                                      // 备注
	Sort                  int64  `gorm:"default:0;comment:'排序值，数值越大排序越靠前';index:idx_withdraw_channels_sort"`                               // 排序值
}

// TableName 指定表名
func (*WithdrawChannel) TableName() string {
	return TableNameWithdrawChannel
}
