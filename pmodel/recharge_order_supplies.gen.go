package pmodel

import "github.com/QuantumShiftX/golib/stores/gormx"

const TableNameRechargeOrderSupply = "recharge_order_supplies"

// RechargeOrderSupply 映射自表 <recharge_order_supplies>
type RechargeOrderSupply struct {
	ID                       int64   `gorm:"column:id;primaryKey;autoIncrement:true;comment:自增主键" json:"id"`                                                                         // 自增主键
	OrderNumber              string  `gorm:"column:order_number;not null;index:idx_recharge_order_supplies_order_number;comment:原订单号" json:"order_number"`                           // 原订单号
	OrderSupplyNumber        string  `gorm:"column:order_supply_number;not null;uniqueIndex:idx_recharge_order_supplies_order_supply_number;comment:补单号" json:"order_supply_number"` // 补单号
	Status                   int8    `gorm:"column:status;not null;index:idx_recharge_order_supplies_status,idx_recharge_order_supplies_user_status;comment:状态" json:"status"`       // 状态
	UserCurrencyCode         string  `gorm:"column:user_currency_code;not null;size:10;comment:用户货币类型" json:"user_currency_code"`                                                    // 用户货币类型
	UserID                   int64   `gorm:"column:user_id;not null;index:idx_recharge_order_supplies_user_id,idx_recharge_order_supplies_user_status;comment:用户ID" json:"user_id"`  // 用户ID
	Username                 string  `gorm:"column:username;not null;size:60;comment:用户名" json:"username"`                                                                           // 用户名
	BeforeOrderAmount        int64   `gorm:"column:before_order_amount;not null;comment:原订单订单金额" json:"before_order_amount"`                                                         // 原订单订单金额
	BeforeGiftAmount         int64   `gorm:"column:before_gift_amount;not null;comment:原订单赠送金额" json:"before_gift_amount"`                                                           // 原订单赠送金额
	BeforeRealAmount         int64   `gorm:"column:before_real_amount;not null;comment:原订单总上分金额" json:"before_real_amount"`                                                          // 原订单总上分金额
	BeforeFeeAmount          int64   `gorm:"column:before_fee_amount;default:0;comment:补单前手续费" json:"before_fee_amount"`                                                             // 补单前手续费
	OrderAmount              int64   `gorm:"column:order_amount;not null;comment:补单后订单金额" json:"order_amount"`                                                                       // 补单后订单金额
	GiftAmount               int64   `gorm:"column:gift_amount;not null;comment:补单后赠送金额" json:"gift_amount"`                                                                         // 补单后赠送金额
	RealAmount               int64   `gorm:"column:real_amount;not null;comment:补单后总上分金额" json:"real_amount"`                                                                        // 补单后总上分金额
	FeeAmount                int64   `gorm:"column:fee_amount;default:0;comment:补单后手续费" json:"fee_amount"`                                                                           // 补单后手续费
	ReasonType               int8    `gorm:"column:reason_type;comment:补单原因类型" json:"reason_type"`                                                                                   // 补单原因类型
	BeforeRechargeChannelID  int64   `gorm:"column:before_recharge_channel_id;not null;comment:原订单充值通道id" json:"before_recharge_channel_id"`                                         // 原订单充值通道id
	BeforeRechargeMerchantID int64   `gorm:"column:before_recharge_merchant_id;not null;comment:原订单充值商户id" json:"before_recharge_merchant_id"`                                       // 原订单充值商户id
	BeforeRealName           string  `gorm:"column:before_real_name;size:64;comment:原订单真实姓名" json:"before_real_name"`                                                                // 原订单真实姓名
	BeforeReceiveAccount     string  `gorm:"column:before_receive_account;size:128;comment:原订单收款账号" json:"before_receive_account"`                                                   // 原订单收款账号
	RechargeChannelID        int64   `gorm:"column:recharge_channel_id;not null;comment:补单后充值通道id" json:"recharge_channel_id"`                                                       // 补单后充值通道id
	RechargeMerchantID       int64   `gorm:"column:recharge_merchant_id;not null;comment:补单后充值商户id" json:"recharge_merchant_id"`                                                     // 补单后充值商户id
	RealName                 string  `gorm:"column:real_name;size:64;comment:补单后真实姓名" json:"real_name"`                                                                              // 补单后真实姓名
	ReceiveAccount           string  `gorm:"column:receive_account;size:128;comment:补单后收款账号" json:"receive_account"`                                                                 // 补单后收款账号
	AuditMulti               float64 `gorm:"column:audit_multi;type:numeric(10,2);not null;comment:审核倍率" json:"audit_multi"`                                                         // 审核倍率
	Applicant                string  `gorm:"column:applicant;not null;size:60;index:idx_recharge_order_supplies_applicant;comment:申请人" json:"applicant"`                             // 申请人
	ApplyTime                int64   `gorm:"column:apply_time;not null;comment:申请时间" json:"apply_time"`                                                                              // 申请时间
	Reviewer                 string  `gorm:"column:reviewer;not null;size:60;index:idx_recharge_order_supplies_reviewer;comment:审核人" json:"reviewer"`                                // 审核人
	ReviewTime               int64   `gorm:"column:review_time;not null;comment:审核时间" json:"review_time"`                                                                            // 审核时间
	Remark                   string  `gorm:"column:remark;size:1000;comment:备注" json:"remark"`                                                                                       // 备注
	TransferVoucher          string  `gorm:"column:transfer_voucher;size:255;comment:转账凭证" json:"transfer_voucher"`                                                                  // 转账凭证
	RechargeInfo             string  `gorm:"column:recharge_info;type:jsonb;default:'{}';comment:充值信息" json:"recharge_info"`                                                         // 充值信息
	//
	gormx.OperationBaseModel
	gormx.Model
}

// TableName RechargeOrderSupply 表名
func (*RechargeOrderSupply) TableName() string {
	return TableNameRechargeOrderSupply
}
