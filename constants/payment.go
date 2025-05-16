package constants

const (
	// 货币USDT
	CurrencyUSDT = "USDT"
)

// FundFlowType 表示资金流动的类型
type FundFlowType int8

const (
	_                FundFlowType = iota // 未定义
	FundFlowDeposit                      // 充值
	FundFlowWithdraw                     // 提现
)

var FundFlowTypeStrings = map[FundFlowType]string{
	FundFlowDeposit:  "deposit",  // 充值
	FundFlowWithdraw: "withdraw", // 提现
}

func (f FundFlowType) Int64() int64 {
	return int64(f)
}

func (f FundFlowType) String() string {
	if str, ok := FundFlowTypeStrings[f]; ok {
		return str
	}
	return "unknown"
}

type RechargeChannelType int8

const (
	_                           RechargeChannelType = iota // none
	RechargeChannelTypeUSDT                                // USDT虚拟货币
	RechargeChannelTypeBankCard                            // 银行卡
	RechargeChannelTypeQuickPay                            // 快捷支付
)

func (t RechargeChannelType) Default() RechargeChannelType {
	if t < RechargeChannelTypeUSDT || t > RechargeChannelTypeQuickPay {
		return RechargeChannelTypeUSDT
	}
	return t
}

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

// RechargeOrderStatus 根据原始的订单状态返回合并后的状态
func (t RechargeOrderStatusType) RechargeOrderStatus() OrderStatus {
	switch t {
	case RechargeOrderStatusTypePayWaiting, RechargeOrderStatusTypeExamineRepair, RechargeOrderStatusTypeExamineAgain, RechargeOrderStatusTypeExamineWaiting:
		// 处理中: 待支付、补单审核中、二次审核中、待审核
		return OrderStatusProcessing
	case RechargeOrderStatusTypePaySuccess:
		// 处理成功: 支付成功
		return OrderStatusSuccess
	case RechargeOrderStatusTypePayTimeout, RechargeOrderStatusTypePayFailed, RechargeOrderStatusTypeCancel:
		// 处理失败: 支付超时、支付失败、已取消
		return OrderStatusFailed
	default:
		return OrderStatusFailed // 未知状态
	}
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

// WithdrawOrderStatus 定义提现订单状态的枚举类型
type WithdrawOrderStatus int

// 定义提现订单状态的枚举值
const (
	WithdrawOrderStatusPending           WithdrawOrderStatus = iota + 1 // 1 - 待出款
	WithdrawOrderStatusPendingUnlocked                                  // 2 - 待出款（未锁定）---
	WithdrawOrderStatusPendingLocked                                    // 3 - 待出款（已锁定）---
	WithdrawOrderStatusPendingThirdParty                                // 4 - 待三方付款
	WithdrawOrderStatusPaymentFailed                                    // 5 - 付款失败
	WithdrawOrderStatusRejected                                         // 6 - 已拒绝
	WithdrawOrderStatusCancelled                                        // 7 - 已取消
	WithdrawOrderStatusPaid                                             // 8 - 已付款
	WithdrawOrderStatusForcePaid                                        // 9 - 已强制付款
)

func (s WithdrawOrderStatus) Int64() int64 {
	return int64(s)
}

// WithdrawOrderStatus 根据提现订单状态返回合并后的状态
func (s WithdrawOrderStatus) WithdrawOrderStatus() OrderStatus {
	switch s {
	case WithdrawOrderStatusPending, WithdrawOrderStatusPendingUnlocked, WithdrawOrderStatusPendingLocked, WithdrawOrderStatusPendingThirdParty:
		// 处理中：待出款、待三方付款、待锁定等
		return OrderStatusProcessing
	case WithdrawOrderStatusPaid, WithdrawOrderStatusForcePaid:
		// 处理成功：已付款、已强制付款
		return OrderStatusSuccess
	case WithdrawOrderStatusPaymentFailed, WithdrawOrderStatusRejected, WithdrawOrderStatusCancelled:
		// 处理失败：付款失败、已拒绝、已取消
		return OrderStatusFailed
	default:
		// 未知状态（如果有其他状态可扩展）
		return OrderStatusFailed
	}
}

// String 方法，返回提现订单状态的中文描述
func (s WithdrawOrderStatus) String() string {
	switch s {
	case WithdrawOrderStatusPending:
		return "待出款"
	case WithdrawOrderStatusPendingUnlocked:
		return "待出款（未锁定）"
	case WithdrawOrderStatusPendingLocked:
		return "待出款（已锁定）"
	case WithdrawOrderStatusPendingThirdParty:
		return "待三方付款"
	case WithdrawOrderStatusPaymentFailed:
		return "付款失败"
	case WithdrawOrderStatusRejected:
		return "已拒绝"
	case WithdrawOrderStatusCancelled:
		return "已取消"
	case WithdrawOrderStatusPaid:
		return "已付款"
	case WithdrawOrderStatusForcePaid:
		return "已强制付款"
	default:
		return "未知状态"
	}
}
