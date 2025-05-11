package pmodel

import "github.com/QuantumShiftX/golib/stores/gormx"

// 代付记录表
const TableNameWithdrawPaymentRecord = "withdraw_payment_records"

// WithdrawPaymentRecord mapped from table <withdraw_payment_records>
type WithdrawPaymentRecord struct {
	ID                 int64  `gorm:"column:id;primaryKey;autoIncrement:true;comment:自增主键" json:"id"`                                                                                              // 自增主键
	OrderNumber        string `gorm:"column:order_number;type:varchar(500);not null;index:idx_withdraw_payment_records_order_number,priority:1;comment:订单号" json:"order_number"`                   // 订单号
	ThirdOrderNumber   string `gorm:"column:third_order_number;type:varchar(500);comment:三方订单号" json:"third_order_number"`                                                                         // 三方订单号
	WithdrawMerchantID int64  `gorm:"column:withdraw_merchant_id;not null;comment:三方代付商户ID" json:"withdraw_merchant_id"`                                                                           // 三方代付商户ID
	WithdrawTypeID     int64  `gorm:"column:withdraw_type_id;not null;comment:提现方式" json:"withdraw_type_id"`                                                                                       // 提现方式
	Status             int64  `gorm:"column:status;not null;index:idx_withdraw_payment_records_user_id_status,priority:2;comment:代付状态 1-订单提交成功 2-付款失败 3-已付款 4-已强制出款 5-已强制取消 6-强制失败" json:"status"` // 代付状态 1-订单提交成功 2-付款失败 3-已付款 4-已强制出款 5-已强制取消 6-强制失败
	Remark             string `gorm:"column:remark;type:varchar(500);comment:备注" json:"remark"`                                                                                                    // 备注
	Req                string `gorm:"column:req;type:text;comment:请求参数" json:"req"`                                                                                                                // 请求参数
	Resp               string `gorm:"column:resp;type:text;comment:响应参数" json:"resp"`                                                                                                              // 响应参数
	UserID             int64  `gorm:"column:user_id;index:idx_withdraw_payment_records_user_id_status,priority:1;comment:用户ID" json:"user_id"`                                                     // 用户ID
	Amount             int64  `gorm:"column:amount;comment:提现金额" json:"amount"`                                                                                                                    // 提现金额
	CurrencyCode       string `gorm:"column:currency_code;type:varchar(500);comment:币种" json:"currency_code"`                                                                                      // 币种
	gormx.OperationBaseModel
	gormx.Model
}

// TableName WithdrawPaymentRecord's table name
func (*WithdrawPaymentRecord) TableName() string {
	return TableNameWithdrawPaymentRecord
}
