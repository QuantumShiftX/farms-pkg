package pmodel

const TableNameRechargeOrder = "recharge_orders"

// RechargeOrder 映射自表 <recharge_orders>
type RechargeOrder struct {
	ID                           int64   `gorm:"column:id;primaryKey;autoIncrement:true;comment:主键" json:"id"`                                                                                                                       // 主键
	OrderNumber                  string  `gorm:"column:order_number;not null;uniqueIndex:idx_recharge_orders_order_number;comment:订单编号" json:"order_number"`                                                                         // 订单编号
	ThirdOrderNumber             string  `gorm:"column:third_order_number;not null;index:idx_recharge_orders_third_order;comment:第三方订单编号" json:"third_order_number"`                                                                 // 第三方订单编号
	UserID                       int64   `gorm:"column:user_id;not null;index:idx_recharge_orders_user_status,idx_recharge_orders_user_status_success,idx_recharge_orders_user_type_status;comment:用户id" json:"user_id"`             // 用户id
	Username                     string  `gorm:"column:username;not null;size:60;comment:用户名" json:"username"`                                                                                                                       // 用户名
	RechargeCurrencyCode         string  `gorm:"column:recharge_currency_code;not null;size:10;comment:充值通道货币代码" json:"recharge_currency_code"`                                                                                      // 充值通道货币代码
	UserCurrencyCode             string  `gorm:"column:user_currency_code;not null;size:10;comment:用户货币代码" json:"user_currency_code"`                                                                                                // 用户货币代码
	ExchangeRate                 string  `gorm:"column:exchange_rate;size:25;default:1:1;comment:货币兑换比例" json:"exchange_rate"`                                                                                                       // 货币兑换比例
	RechargeAmount               int64   `gorm:"column:recharge_amount;not null;comment:充值金额（通道充值币种）" json:"recharge_amount"`                                                                                                        // 充值金额（通道充值币种）
	OrderAmount                  int64   `gorm:"column:order_amount;not null;comment:订单金额（会员/用户币种）" json:"order_amount"`                                                                                                             // 订单金额（会员/用户币种）
	GiftAmount                   int64   `gorm:"column:gift_amount;not null;default:0;comment:赠送金额（会员/用户币种）" json:"gift_amount"`                                                                                                     // 赠送金额（会员/用户币种）
	RealAmount                   int64   `gorm:"column:real_amount;not null;comment:总上分金额（会员/用户币种）" json:"real_amount"`                                                                                                              // 总上分金额（会员/用户币种）
	FeeAmount                    int64   `gorm:"column:fee_amount;not null;default:0;comment:手续费（会员/用户币种）" json:"fee_amount"`                                                                                                        // 手续费（会员/用户币种）
	OrderType                    int8    `gorm:"column:order_type;not null;index:idx_recharge_orders_user_type_status;comment:充值订单类型" json:"order_type"`                                                                             // 充值订单类型
	OrderStatus                  int8    `gorm:"column:order_status;not null;index:idx_recharge_orders_user_status,idx_recharge_orders_user_status_success,idx_recharge_orders_user_type_status;comment:充值订单状态" json:"order_status"` // 充值订单状态
	OrderTag                     int8    `gorm:"column:order_tag;not null;default:0;comment:充值订单标签" json:"order_tag"`                                                                                                                // 充值订单标签
	OrderSuccessAt               int64   `gorm:"column:order_success_at;index:idx_recharge_orders_user_status_success;comment:订单支付成功时间" json:"order_success_at"`                                                                     // 订单支付成功时间
	ReceiveAccount               string  `gorm:"column:receive_account;size:128;comment:收款账户/地址" json:"receive_account"`                                                                                                             // 收款账户/地址
	ReceiveName                  string  `gorm:"column:receive_name;size:64;comment:收款人姓名" json:"receive_name"`                                                                                                                      // 收款人姓名
	RechargeInfo                 string  `gorm:"column:recharge_info;type:jsonb;comment:充值信息" json:"recharge_info"`                                                                                                                  // 充值信息
	RechargeCategoryID           int64   `gorm:"column:recharge_category_id;not null;comment:充值大类id" json:"recharge_category_id"`                                                                                                    // 充值大类id
	RechargeMerchantID           int64   `gorm:"column:recharge_merchant_id;not null;index:idx_recharge_orders_merchant_channel;comment:充值商户id" json:"recharge_merchant_id"`                                                         // 充值商户id
	RechargeChannelID            int64   `gorm:"column:recharge_channel_id;not null;index:idx_recharge_orders_merchant_channel;comment:充值通道id" json:"recharge_channel_id"`                                                           // 充值通道id
	RechargeCount                int64   `gorm:"column:recharge_count;not null;default:0;comment:充值次数" json:"recharge_count"`                                                                                                        // 充值次数
	ProxyRechargeCustomerService string  `gorm:"column:proxy_recharge_customer_service;size:60;comment:代充客服" json:"proxy_recharge_customer_service"`                                                                                 // 代充客服
	TransferVoucher              string  `gorm:"column:transfer_voucher;size:255;comment:转账凭证" json:"transfer_voucher"`                                                                                                              // 转账凭证
	FrontRemark                  string  `gorm:"column:front_remark;size:255;comment:前台备注" json:"front_remark"`                                                                                                                      // 前台备注
	BackRemark                   string  `gorm:"column:back_remark;size:255;comment:后台备注" json:"back_remark"`                                                                                                                        // 后台备注
	ErrorMsg                     string  `gorm:"column:error_msg;size:255;comment:错误详情" json:"error_msg"`                                                                                                                            // 错误详情
	Utr                          string  `gorm:"column:utr;size:64;comment:UTR" json:"utr"`                                                                                                                                          // UTR
	RealName                     string  `gorm:"column:real_name;size:64;comment:充值用户真实姓名" json:"real_name"`                                                                                                                         // 充值用户真实姓名
	AuditMulti                   float64 `gorm:"column:audit_multi;type:numeric(10,2);default:1.00;comment:稽核倍数" json:"audit_multi"`                                                                                                 // 稽核倍数
	BankID                       int64   `gorm:"column:bank_id;comment:银行id" json:"bank_id"`                                                                                                                                         // 银行id
	IP                           string  `gorm:"column:ip;size:60;comment:操作IP" json:"ip"`                                                                                                                                           // 操作IP
	RechargeCountType            int8    `gorm:"column:recharge_count_type;default:0;comment:充值次数类型" json:"recharge_count_type"`                                                                                                     // 充值次数类型
	OperatorID                   int64   `gorm:"column:operator_id;not null;default:0;comment:操作员ID" json:"operator_id"`                                                                                                             // 操作员ID
	Operator                     string  `gorm:"column:operator;size:255;comment:操作员姓名" json:"operator"`                                                                                                                             // 操作员姓名
	OperationTime                int64   `gorm:"column:operation_time;not null;default:0;comment:操作时间" json:"operation_time"`                                                                                                        // 操作时间
	CreatedAt                    int64   `gorm:"column:created_at;not null;default:0;index:idx_recharge_orders_created_at;comment:创建时间" json:"created_at"`                                                                           // 创建时间
	UpdatedAt                    int64   `gorm:"column:updated_at;not null;default:0;comment:更新时间" json:"updated_at"`                                                                                                                // 更新时间
	DeletedAt                    int64   `gorm:"column:deleted_at;not null;default:0;comment:删除时间" json:"deleted_at"`                                                                                                                // 删除时间
}

// TableName RechargeOrder 表名
func (*RechargeOrder) TableName() string {
	return TableNameRechargeOrder
}
