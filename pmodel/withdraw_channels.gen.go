package pmodel

import (
	"github.com/QuantumShiftX/golib/stores/gormx"
)

const TableNameWithdrawChannel = "withdraw_channels"

// WithdrawChannel 表示提现渠道子表
type WithdrawChannel struct {
	gormx.OperationBaseModel
	gormx.Model
	Status int64 `gorm:"column:status;not null;default:1;comment:'通道状态：1-启用，2-禁用';index:idx_withdraw_channels_status" json:"status"` // 通道启用状态
	//
	ID                    int64  `gorm:"column:id;primaryKey;autoIncrement;comment:'主键ID'" json:"id"`                                                                                              // 主键ID
	PaymentConfigID       int64  `gorm:"column:payment_config_id;not null;comment:'支付配置ID，关联到payment_configs表的主键';index:idx_withdraw_channels_payment_config_id" json:"payment_config_id"`         // 支付配置ID
	CurrencyCode          string `gorm:"column:currency_code;type:varchar(255);not null;comment:'币种代码，如\"USDT\"、\"CNY\"等';index:idx_withdraw_channels_currency" json:"currency_code"`              // 币种代码
	WithdrawChannelName   string `gorm:"column:withdraw_channel_name;type:varchar(255);not null;comment:'提现通道名称，用于前端展示'" json:"withdraw_channel_name"`                                             // 提现通道名称
	WithdrawChannelCode   string `gorm:"column:withdraw_channel_code;type:varchar(255);not null;comment:'提现通道编码，系统内部唯一标识';index:idx_withdraw_channels_code_currency" json:"withdraw_channel_code"` // 提现通道编码
	WithdrawChannelType   int64  `gorm:"column:withdraw_channel_type;not null;comment:'提现通道类型：1-USDT虚拟货币，2-银行卡，3-快捷支付';index:idx_withdraw_channels_type" json:"withdraw_channel_type"`             // 提现通道类型
	WithdrawChannelLogo   string `gorm:"column:withdraw_channel_logo;type:varchar(500);not null;comment:'提现通道logo图片URL'" json:"withdraw_channel_logo"`                                             // 提现通道logo
	WithdrawMerchantID    int64  `gorm:"column:withdraw_merchant_id;not null;comment:'提现商户配置ID，关联到支付商户表';index:idx_withdraw_channels_merchant_id" json:"withdraw_merchant_id"`                     // 提现商户配置ID
	BankID                int64  `gorm:"column:bank_id;comment:'银行ID，关联到银行信息表';index:idx_withdraw_channels_bank_id" json:"bank_id"`                                                                // 银行ID
	BankName              string `gorm:"column:bank_name;type:varchar(255);comment:'银行名称，如\"中国工商银行\"'" json:"bank_name"`                                                                           // 银行名称
	BankBranch            string `gorm:"column:bank_branch;type:varchar(255);comment:'银行支行名称，如\"北京市海淀区支行\"'" json:"bank_branch"`                                                                   // 银行支行
	BankAccount           string `gorm:"column:bank_account;type:varchar(255);comment:'银行账号，收款账户号码'" json:"bank_account"`                                                                          // 银行账号
	BankAccountName       string `gorm:"column:bank_account_name;type:varchar(255);comment:'账户名称，银行卡开户人姓名'" json:"bank_account_name"`                                                              // 账户名称
	CryptoAddress         string `gorm:"column:crypto_address;type:varchar(500);comment:'加密货币钱包地址，用于收款的区块链地址'" json:"crypto_address"`                                                              // 加密货币地址
	CryptoNetwork         string `gorm:"column:crypto_network;type:varchar(50);comment:'网络类型，如\"TRC20\"、\"ERC20\"等区块链网络'" json:"crypto_network"`                                                   // 网络类型
	MinLimitAmount        int64  `gorm:"column:min_limit_amount;not null;comment:'单笔最小限额，提现金额下限'" json:"min_limit_amount"`                                                                         // 单笔最小限额
	MaxLimitAmount        int64  `gorm:"column:max_limit_amount;not null;comment:'单笔最大限额，提现金额上限'" json:"max_limit_amount"`                                                                         // 单笔最大限额
	DailyLimitAmount      int64  `gorm:"column:daily_limit_amount;default:0;comment:'每日提现总额限制，0表示不限制'" json:"daily_limit_amount"`                                                                  // 每日限额
	FeePercent            int64  `gorm:"column:fee_percent;not null;default:0;comment:'手续费百分比'" json:"fee_percent"`                                                                                // 手续费（百分比）
	FixedFee              int64  `gorm:"column:fixed_fee;default:0;comment:'固定手续费金额，每笔提现固定收取的费用'" json:"fixed_fee"`                                                                                // 固定手续费
	FeeReductionConfig    string `gorm:"column:fee_reduction_config;type:jsonb;default:'{}';comment:'手续费减免配置，JSON格式，可配置不同条件下的手续费减免规则'" json:"fee_reduction_config"`                                // 手续费减免配置
	RecommendAmountConfig string `gorm:"column:recommend_amount_config;type:jsonb;default:'{}';comment:'推荐金额配置，JSON格式，前端展示的快捷金额选项'" json:"recommend_amount_config"`                                // 推荐金额配置
	Hint                  string `gorm:"column:hint;type:text;comment:'提示信息，展示给用户的操作指引，最大1000字'" json:"hint"`                                                                                      // 提示
	Remark                string `gorm:"column:remark;type:text;comment:'备注信息，内部管理用'" json:"remark"`                                                                                               // 备注
	Sort                  int64  `gorm:"column:sort;default:0;comment:'排序值，数值越大排序越靠前';index:idx_withdraw_channels_sort" json:"sort"`                                                               // 排序值
}

// TableName 指定表名
func (*WithdrawChannel) TableName() string {
	return TableNameWithdrawChannel
}
