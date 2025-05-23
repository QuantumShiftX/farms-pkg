package pmodel

import "github.com/QuantumShiftX/golib/stores/gormx"

const TableNameWithdrawOrder = "withdraw_orders"

// WithdrawOrder 映射自表 <withdraw_orders>
type WithdrawOrder struct {
	ID                   int64  `gorm:"column:id;primaryKey;autoIncrement:true;comment:自增主键" json:"id"`                                                                                                                              // 自增主键
	OrderNumber          string `gorm:"column:order_number;not null;uniqueIndex:idx_withdraw_orders_order_number;comment:订单号" json:"order_number"`                                                                                   // 订单号
	ThirdOrderNumber     string `gorm:"column:third_order_number;index:idx_withdraw_orders_third_order_number;comment:三方订单号" json:"third_order_number"`                                                                              // 三方订单号
	UserID               int64  `gorm:"column:user_id;not null;index:idx_withdraw_orders_user_status_success;comment:用户id" json:"user_id"`                                                                                           // 用户id
	Username             string `gorm:"column:username;not null;size:60;comment:用户名" json:"username"`                                                                                                                                // 用户名
	UserCurrencyCode     string `gorm:"column:user_currency_code;not null;size:10;comment:用户货币类型" json:"user_currency_code"`                                                                                                         // 用户货币类型
	WithdrawCurrencyCode string `gorm:"column:withdraw_currency_code;not null;size:10;index:idx_withdraw_orders_status_success_currency;comment:提现货币类型" json:"withdraw_currency_code"`                                               // 提现货币类型
	PaymentConfigID      int64  `gorm:"column:payment_config_id;not null;comment:支付方式ID" json:"payment_config_id"`                                                                                                                   // 支付方式ID
	WithdrawAmount       int64  `gorm:"column:withdraw_amount;not null;comment:提现金额" json:"withdraw_amount"`                                                                                                                         // 提现金额
	ReceivedAmount       int64  `gorm:"column:received_amount;not null;comment:预计到账金额" json:"received_amount"`                                                                                                                       // 预计到账金额
	FeeAmount            int64  `gorm:"column:fee_amount;not null;default:0;comment:手续费" json:"fee_amount"`                                                                                                                          // 手续费
	ActualReceivedAmount int64  `gorm:"column:actual_received_amount;not null;comment:实际到账金额" json:"actual_received_amount"`                                                                                                         // 实际到账金额
	WithdrawMerchantID   int64  `gorm:"column:withdraw_merchant_id;not null;index:idx_withdraw_orders_merchant_id,idx_withdraw_orders_status_merchant;comment:三方代付商户ID" json:"withdraw_merchant_id"`                                 // 三方代付商户ID
	WithdrawChannelID    int64  `gorm:"column:withdraw_channel_id;not null;comment:提现通道id" json:"withdraw_channel_id"`                                                                                                               // 提现通道id
	RealName             string `gorm:"column:real_name;not null;size:64;comment:真实姓名" json:"real_name"`                                                                                                                             // 真实姓名
	Account              string `gorm:"column:account;not null;size:128;comment:收款账号" json:"account"`                                                                                                                                // 收款账号
	BankID               int64  `gorm:"column:bank_id;comment:银行Id" json:"bank_id"`                                                                                                                                                  // 银行Id
	Ifsc                 string `gorm:"column:ifsc;size:64;comment:IFSC代码" json:"ifsc"`                                                                                                                                              // IFSC代码
	Utr                  string `gorm:"column:utr;size:64;comment:UTR参考号" json:"utr"`                                                                                                                                                // UTR参考号
	OrderStatus          int8   `gorm:"column:order_status;not null;index:idx_withdraw_orders_status_success_currency,idx_withdraw_orders_user_status_success,idx_withdraw_orders_status_merchant;comment:订单状态" json:"order_status"` // 订单状态 1-待出款 2-待出款（未锁定） 3-待出款（已锁定） 4-待三方付款 5-付款失败 6-已拒绝 7-已取消 8-已付款 9-已强制付款';
	RemarkBack           string `gorm:"column:remark_back;size:1000;comment:后台备注" json:"remark_back"`                                                                                                                                // 后台备注
	RemarkFront          string `gorm:"column:remark_front;size:1000;comment:前台备注" json:"remark_front"`                                                                                                                              // 前台备注
	IP                   string `gorm:"column:ip;not null;size:60;comment:提现时IP" json:"ip"`                                                                                                                                          // 提现时IP
	SuccessAt            int64  `gorm:"column:success_at;default:0;index:idx_withdraw_orders_status_success_currency,idx_withdraw_orders_user_status_success;comment:提现成功时间" json:"success_at"`                                      // 提现成功时间
	LockStatus           int8   `gorm:"column:lock_status;default:0;index:idx_withdraw_orders_lock_status;comment:锁定状态 1-未锁定 2-已锁定" json:"lock_status"`                                                                              // 锁定状态
	RiskCheckStatus      int8   `gorm:"column:risk_check_status;default:1;comment:风控检查状态 1-未检查 2-检查通过 3-检查未通过'" json:"risk_check_status"`                                                                                            // 风控检查状态
	Rate                 int64  `gorm:"column:rate;comment:汇率" json:"rate"`                                                                                                                                                          // 汇率
	AutoWithdraw         int8   `gorm:"column:auto_withdraw;default:2;comment:是否自动审核 1-是 2-否" json:"auto_withdraw"`                                                                                                                  // 是否自动审核
	UserRiskInfo         string `gorm:"column:user_risk_info;type:jsonb;default:'{}';comment:提现时用户风控信息" json:"user_risk_info"`                                                                                                       // 提现时用户风控信息
	ErrorMsg             string `gorm:"column:error_msg;size:255;comment:错误信息" json:"error_msg"`                                                                                                                                     // 错误信息
	TransferVoucher      string `gorm:"column:transfer_voucher;size:255;comment:转账凭证" json:"transfer_voucher"`                                                                                                                       // 转账凭证
	gormx.OperationBaseModel
	gormx.Model
}

// TableName WithdrawOrder 表名
func (*WithdrawOrder) TableName() string {
	return TableNameWithdrawOrder
}
