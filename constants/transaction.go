package constants

// 定义账变大类枚举
type TxnCategoryType int

const (
	TxnDeposit          TxnCategoryType = iota + 1 // 1: 存款 (充值)
	TxnWithdraw                                    // 2: 提款 (提现)
	TxnPurchase                                    // 3: 购买
	TxnCurrencyExchange                            // 4: 换汇
	TxnRebate                                      // 5: 返佣
	TxnReward                                      // 6: 奖励
	TxnHarvest                                     // 7: 收获
	TxnDepositPromotion                            // 8: 充值奖励
	TxnFundCorrection                              // 9: 资金修正
)

// 定义枚举类型的字符串映射
func (o TxnCategoryType) String() string {
	switch o {
	case TxnDeposit:
		return "存款"
	case TxnWithdraw:
		return "提款"
	case TxnPurchase:
		return "购买"
	case TxnCurrencyExchange:
		return "换汇"
	case TxnRebate:
		return "返佣"
	case TxnReward:
		return "奖励"
	case TxnHarvest:
		return "收获"
	case TxnDepositPromotion:
		return "充值奖励"
	case TxnFundCorrection:
		return "资金修正"
	default:
		return "未知操作"
	}
}

// 定义账变小类枚举
type TxnSubCategoryType int

const (
	TxnSubDeposit          TxnSubCategoryType = iota + 1 // 1: 存款
	TxnSubWithdraw                                       // 2: 提款
	TxnSubPurchase                                       // 3: 购买
	TxnSubCurrencyExchange                               // 4: 换汇
	TxnSubRebate                                         // 5: 返佣
	TxnSubReward                                         // 6: 返佣
	TxnSubHarvest                                        // 7: 返佣
)
