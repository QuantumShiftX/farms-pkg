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

func (o TxnCategoryType) Int64() int64 {
	return int64(o)
}

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

func (o TxnSubCategoryType) Int64() int64 {
	return int64(o)
}

// TxnDeposit的子类 // 1: 存款
const (
	TxnSubDeposit        TxnSubCategoryType = 1000
	TxnSubSystemDeposit                     = iota + TxnSubDeposit // 1003: 系统充值
	TxnSubServiceDeposit                                           // 1004: 客服充值
)

// TxnWithdraw的子类 // 2: 提款
const (
	TxnSubWithdraw       TxnSubCategoryType = 2000
	TxnSubWithdrawNormal                    = iota + TxnSubWithdraw // 2001: 普通提款
	TxnSubWithdrawFast                                              // 2002: 快速提款
	TxnSubWithdrawAuto                                              // 2003: 自动提款
)

// TxnPurchase的子类 // 3: 购买
const (
	TxnSubPurchase        TxnSubCategoryType = 3000
	TxnSubPurchaseGoods                      = iota + TxnSubPurchase // 3001: 商品购买
	TxnSubPurchaseService                                            // 3002: 服务购买
	TxnSubPurchaseMember                                             // 3003: 会员购买
)

// TxnCurrencyExchange的子类 // 4: 换汇
const (
	TxnSubCurrencyExchange         TxnSubCategoryType = 4000
	TxnSubCurrencyExchangeNormal                      = iota + TxnSubCurrencyExchange // 4001: 常规换汇
	TxnSubCurrencyExchangeDiscount                                                    // 4002: 优惠换汇
)

// TxnRebate的子类 // 5: 返佣
const (
	TxnSubRebate         TxnSubCategoryType = 5000
	TxnSubRebateDirect                      = iota + TxnSubRebate // 5001: 直接返佣
	TxnSubRebateIndirect                                          // 5002: 间接返佣
	TxnSubRebateTeam                                              // 5003: 团队返佣
)

// TxnReward的子类 // 6: 奖励
const (
	TxnSubReward            TxnSubCategoryType = 6000
	TxnSubRewardFortuneTree                    = iota + TxnSubReward // 6001: 发财树奖励
	TxnSubRewardActivity                                             // 6001: 活动奖励
	TxnSubRewardSignIn                                               // 6002: 签到奖励
	TxnSubRewardTask                                                 // 6003: 任务奖励
	TxnSubRewardVIP                                                  // 6004: VIP奖励
)

// TxnHarvest的子类 // 7: 收获
const (
	TxnSubHarvest          TxnSubCategoryType = 7000
	TxnSubHarvestPlanting                     = iota + TxnSubHarvest // 7001: 种植收益
	TxnSubHarvestFixedTerm                                           // 7002: 定期收益
	TxnSubHarvestDividend                                            // 7003: 分红收益
)

// TxnDepositPromotion的子类 // 8: 充值奖励
const (
	TxnSubDepositPromotion          TxnSubCategoryType = 8000
	TxnSubDepositPromotionFirstTime                    = iota + TxnSubDepositPromotion // 8001: 首充奖励
	TxnSubDepositPromotionAmount                                                       // 8002: 充值满送
	TxnSubDepositPromotionLimited                                                      // 8003: 限时充值奖励
)

// TxnFundCorrection的子类 // 9: 资金修正
const (
	TxnSubFundCorrection       TxnSubCategoryType = 9000
	TxnSubFundCorrectionSystem                    = iota + TxnSubFundCorrection // 9001: 系统调整
	TxnSubFundCorrectionManual                                                  // 9002: 人工调整
	TxnSubFundCorrectionError                                                   // 9003: 错误修正
)

// 子类型字符串表示
func (o TxnSubCategoryType) String() string {
	switch o {
	// 存款子类
	case TxnSubSystemDeposit:
		return "系统充值"
	case TxnSubServiceDeposit:
		return "客服充值"

	// 提款子类
	case TxnSubWithdrawNormal:
		return "普通提款"
	case TxnSubWithdrawFast:
		return "快速提款"
	case TxnSubWithdrawAuto:
		return "自动提款"

	// 购买子类
	case TxnSubPurchaseGoods:
		return "商品购买"
	case TxnSubPurchaseService:
		return "服务购买"
	case TxnSubPurchaseMember:
		return "会员购买"

	// 换汇子类
	case TxnSubCurrencyExchangeNormal:
		return "常规换汇"
	case TxnSubCurrencyExchangeDiscount:
		return "优惠换汇"

	// 返佣子类
	case TxnSubRebateDirect:
		return "直接返佣"
	case TxnSubRebateIndirect:
		return "间接返佣"
	case TxnSubRebateTeam:
		return "团队返佣"

	// 奖励子类
	case TxnSubRewardFortuneTree:
		return "发财树奖励"
	case TxnSubRewardActivity:
		return "活动奖励"
	case TxnSubRewardSignIn:
		return "签到奖励"
	case TxnSubRewardTask:
		return "任务奖励"
	case TxnSubRewardVIP:
		return "VIP奖励"

	// 收获子类
	case TxnSubHarvestPlanting:
		return "种植收益"
	case TxnSubHarvestFixedTerm:
		return "定期收益"
	case TxnSubHarvestDividend:
		return "分红收益"

	// 充值奖励子类
	case TxnSubDepositPromotionFirstTime:
		return "首充奖励"
	case TxnSubDepositPromotionAmount:
		return "充值满送"
	case TxnSubDepositPromotionLimited:
		return "限时充值奖励"

	// 资金修正子类
	case TxnSubFundCorrectionSystem:
		return "系统调整"
	case TxnSubFundCorrectionManual:
		return "人工调整"
	case TxnSubFundCorrectionError:
		return "错误修正"

	default:
		return "未知子类"
	}
}
