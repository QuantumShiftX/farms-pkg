package constants

// RechargeOrderStatusType 充值订单状态
type RechargeOrderStatusType int8

const (
	_                                     RechargeOrderStatusType = iota
	RechargeOrderStatusTypePayWaiting                             // 待支付
	RechargeOrderStatusTypePaySuccess                             // 支付成功
	RechargeOrderStatusTypePayTimeout                             // 支付超时
	RechargeOrderStatusTypePayFailed                              // 支付失败
	RechargeOrderStatusTypeExamineRepair                          // 补单审核中
	RechargeOrderStatusTypeExamineAgain                           // 二次审核中
	RechargeOrderStatusTypeExamineWaiting                         // 待审核
	RechargeOrderStatusTypeCancel                                 // 已取消
)

func (t RechargeOrderStatusType) Int8() int8 {
	return int8(t)
}

// RechargeOrderType 充值订单类型
type RechargeOrderType int8

const (
	_                         RechargeOrderType = iota
	RechargeOrderTypeOnline                     // 在线充值
	RechargeOrderTypeTransfer                   // 转账充值
	RechargeOrderTypeProxy                      // 客服代充
)

func (t RechargeOrderType) Int8() int8 {
	return int8(t)
}

func (t RechargeOrderType) String() string {
	return [...]string{"", "在线充值", "转账充值", "客服代充"}[t]
}

// RechargeOrderTagType 充值订单标签
type RechargeOrderTagType int8

const (
	_                                             RechargeOrderTagType = iota // none
	RechargeOrderTagTypeDepositConfirm                                        // 确认入款
	RechargeOrderTagTypeDepositCancel                                         // 取消入款
	RechargeOrderTagTypeExamineRepairReviewRefuse                             // 补单审核拒绝
	RechargeOrderTagTypeExamineRepairReviewPass                               // 补单审核通过
	RechargeOrderTagTypeExamineRepairNoneReview                               // 无审核🙅‍单
	RechargeOrderTagTypeExamineAgainReviewPass                                // 二次审核通过
	RechargeOrderTagTypeExamineAgainReviewRefuse                              // 二次审核拒绝
	RechargeOrderTagTypePullFailed                                            // 拉单失败
	RechargeOrderTagTypeExamineRepairPass                                     // 补单通过(客服代充用)
	RechargeOrderTagTypeExamineRepairRefuse                                   // 补单拒绝(客服代充用)
	RechargeOrderTagTypeExamineDepositPass                                    // 入款审核通过(客服代充用)
	RechargeOrderTagTypeExamineDepositRefuse                                  // 入款审核拒绝(客服代充用)
)

func (t RechargeOrderTagType) Int8() int8 {
	return int8(t)
}

// 定义账号类型枚举 （用户提现）
type AccountType int

const (
	PhoneNumber AccountType = iota + 1 // 手机号
	BankCard                           // 银行卡
	PIX                                // PIX
	EWallet                            // 电子钱包
)

func (t AccountType) Int8() int8 {
	return int8(t)
}

// OrderStatus 定义订单状态的枚举类型
type OrderStatus int

// 定义订单状态的枚举值
const (
	OrderStatusProcessing OrderStatus = iota + 1 // 1 - 处理中
	OrderStatusSuccess                           // 2 - 处理成功
	OrderStatusFailed                            // 3 - 处理失败
)

func (s OrderStatus) Int64() int64 {
	return int64(s)
}

// String 方法，返回订单状态的中文描述
func (s OrderStatus) String() string {
	switch s {
	case OrderStatusProcessing:
		return "处理中"
	case OrderStatusSuccess:
		return "处理成功"
	case OrderStatusFailed:
		return "处理失败"
	default:
		return "未知状态"
	}
}

// DepositOrderStatus 定义充值订单状态的枚举类型
type DepositOrderStatus int

// 定义充值订单状态的枚举值
const (
	DepositOrderStatusPending           DepositOrderStatus = iota + 1 // 1 - 待出款
	DepositOrderStatusPendingUnlocked                                 // 2 - 待出款（未锁定）
	DepositOrderStatusPendingLocked                                   // 3 - 待出款（已锁定）
	DepositOrderStatusPendingThirdParty                               // 4 - 待三方付款
	DepositOrderStatusPaymentFailed                                   // 5 - 付款失败
	DepositOrderStatusRejected                                        // 6 - 已拒绝
	DepositOrderStatusCancelled                                       // 7 - 已取消
	DepositOrderStatusPaid                                            // 8 - 已付款
	DepositOrderStatusForcePaid                                       // 9 - 已强制付款
)

func (s DepositOrderStatus) Int64() int64 {
	return int64(s)
}

// String 方法，返回充值订单状态的中文描述
func (s DepositOrderStatus) String() string {
	switch s {
	case DepositOrderStatusPending:
		return "待出款"
	case DepositOrderStatusPendingUnlocked:
		return "待出款（未锁定）"
	case DepositOrderStatusPendingLocked:
		return "待出款（已锁定）"
	case DepositOrderStatusPendingThirdParty:
		return "待三方付款"
	case DepositOrderStatusPaymentFailed:
		return "付款失败"
	case DepositOrderStatusRejected:
		return "已拒绝"
	case DepositOrderStatusCancelled:
		return "已取消"
	case DepositOrderStatusPaid:
		return "已付款"
	case DepositOrderStatusForcePaid:
		return "已强制付款"
	default:
		return "未知状态"
	}
}
