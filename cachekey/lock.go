package cachekey

import "time"

const (
	LockUserCommonTimeExpire = 10 * time.Second
)

const (
	LockUserWalletTypeChangeTimeExpire = 5 * time.Second
	// 	LockUserFarmsOpsKey 用户操作农场锁key {user_id} {wallet type}
	LockUserWalletTypeChangeKey = "{lock_user_%d}:ops:balance_%d"
	// UserBalanceLogUniqueKey 唯一检查
	UserBalanceLogUniqueKey = "user_balance_log_unique:%s_%d"
	// CacheUserBalanceUserInfoKey 缓存余额变动所需用户信息
	CacheUserBalanceUserInfoKey = "cache_balance_log_user_info:%d"
)

// 用户余额变动响应code 0失败 1 成功 2余额不足
const (
	UserBalanceChangeRespCodeFailed           = 0
	UserBalanceChangeRespCodeSuccess          = 1
	UserBalanceChangeRespCodeNotEnoughBalance = 2
)

const (
	LockUserFarmsOpsTimeExpire = 10 * time.Second
	// 	LockUserFarmsOpsKey 用户操作农场锁key {user_id} {ops_farm_id}
	LockUserFarmsOpsKey = "{lock}:user_%d_ops_farms_%d"
	// 操作发财树
	LockUserFarmsOpsFortuneTreeKey = "{lock}:user_%d_ops_fortune_tree_%d"
)

const (
	// 	LockUserFarmsOpsKey 用户操作农场锁key {user_id} {商店的id}
	LockUserOpsStoreProductKey = "{lock}:user_%d_ops_product_%d"
)

const (
	LockUserRechargeTimeExpire = 5 * time.Second
	// LockUserRechargeKey 用户充值锁:用户id 充值金额
	LockUserRechargeKey = "{lock:user_%d}:recharge_%d"
	// LockUserRechargeNoticeKey 用户充值回调锁:订单ID 充值金额
	LockUserRechargeNoticeKey = "{lock:order_%s}:recharge_%d"
)

const (
	LockASyncStatusTimeExpire = 10 * time.Second
	// LockUserCropStatusCheckKey 用户作物状态锁:用户id
	LockUserCropStatusCheckKey = "{lock:user_%d}:crop_status_check"
	// LockUserTreeStatusCheckKey 用户发财树状态锁:用户id
	LockUserTreeStatusCheckKey = "{lock:user_%d}:tree_status_check"
)
