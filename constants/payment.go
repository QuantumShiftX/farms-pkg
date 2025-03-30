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
