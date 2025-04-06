package constants

// 定义账变大类枚举
type TxnCategoryType int

const (
	TxnDeposit          TxnCategoryType = iota + 1 // 1: 存款
	TxnWithdraw                                    // 2: 提款
	TxnPayment                                     // 3: 收款
	TxnPurchase                                    // 4: 购买
	TxnCurrencyExchange                            // 5: 换汇
	TxnRebate                                      // 6: 返佣
)

// 定义枚举类型的字符串映射
func (o TxnCategoryType) String() string {
	switch o {
	case TxnDeposit:
		return "存款"
	case TxnWithdraw:
		return "提款"
	case TxnPayment:
		return "收款"
	case TxnPurchase:
		return "购买"
	case TxnCurrencyExchange:
		return "换汇"
	case TxnRebate:
		return "返佣"
	default:
		return "未知操作"
	}
}

// 定义账变小类枚举
type TxnSubCategoryType int

const (
	TxnSubDeposit          TxnSubCategoryType = iota + 1 // 1: 存款
	TxnSubWithdraw                                       // 2: 提款
	TxnSubPayment                                        // 3: 收款
	TxnSubPurchase                                       // 4: 购买
	TxnSubCurrencyExchange                               // 5: 换汇
	TxnSubRebate                                         // 6: 返佣
)
