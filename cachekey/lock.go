package cachekey

import "time"

const (
	LockUserCommonTimeExpire = 10 * time.Second
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
