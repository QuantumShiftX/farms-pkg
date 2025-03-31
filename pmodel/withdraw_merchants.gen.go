package pmodel

import "github.com/QuantumShiftX/golib/stores/gormx"

// 三方代付平台表
const TableNameWithdrawMerchant = "withdraw_merchants"

// WithdrawMerchant mapped from table <withdraw_merchants>
type WithdrawMerchant struct {
	ID                    int64  `gorm:"column:id;primaryKey;autoIncrement:true;comment:自增主键" json:"id"`                                              // 自增主键
	Name                  string `gorm:"column:name;type:varchar(255);not null;comment:三方代付平台名称" json:"name"`                                         // 三方代付平台名称
	CurrencyCode          string `gorm:"column:currency_code;type:varchar(255);not null;comment:币种编码" json:"currency_code"`                           // 币种编码
	PaymentConfigID       int64  `gorm:"column:payment_config_id;type:integer;not null;comment:三方支付/代付id" json:"payment_config_id"`                   // 三方支付/代付id
	MerchantNo            string `gorm:"column:merchant_no;type:varchar(255);not null;comment:商户号" json:"merchant_no"`                                // 商户号
	Key                   string `gorm:"column:key;type:varchar(255);not null;comment:密钥" json:"key"`                                                 // 密钥
	ExternalInfo          string `gorm:"column:external_info;type:text;not null;comment:扩展信息" json:"external_info"`                                   // 扩展信息
	WithdrawOrderPayURL   string `gorm:"column:withdraw_order_pay_url;type:varchar(255);not null;comment:下单地址(提现下单地址)" json:"withdraw_order_pay_url"` // 下单地址(提现下单地址)
	WithdrawOrderQueryURL string `gorm:"column:withdraw_order_query_url;type:varchar(255);not null;comment:订单查询地址" json:"withdraw_order_query_url"`   // 订单查询地址
	BalanceQueryURL       string `gorm:"column:balance_query_url;type:varchar(255);not null;comment:余额查询地址" json:"balance_query_url"`                 // 余额查询地址
	SuccessTag            string `gorm:"column:success_tag;type:varchar(255);default:success;comment:提现成功标识(回调成功标识)" json:"success_tag"`              // 提现成功标识(回调成功标识)
	CallbackAddress       string `gorm:"column:callback_address;type:varchar(255);comment:回调IP地址(多个用逗号分隔)" json:"callback_address"`                   // 回调IP地址(多个用逗号分隔)
	WithdrawTypeIDs       string `gorm:"column:withdraw_type_ids;type:jsonb;comment:提现方式Id's" json:"withdraw_type_ids"`                               // 提现方式Id's
	LimitAmountLower      int64  `gorm:"column:limit_amount_lower;not null;comment:最低提现金额" json:"limit_amount_lower"`                                 // 最低提现金额
	LimitAmountUpper      int64  `gorm:"column:limit_amount_upper;not null;comment:最高提现金额" json:"limit_amount_upper"`                                 // 最高提现金额
	Status                int8   `gorm:"column:status;type:smallint;not null;comment:状态(1:启用,2:禁用)" json:"status"`                                    // 状态(1:启用,2:禁用)
	Remark                string `gorm:"column:remark;type:varchar(1000);not null;comment:备注" json:"remark"`                                          // 备注
	gormx.OperationBaseModel
	gormx.Model
}

// TableName WithdrawMerchant's table name
func (*WithdrawMerchant) TableName() string {
	return TableNameWithdrawMerchant
}
