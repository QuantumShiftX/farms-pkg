package cachekey

import "time"

const (
	LockUserCommonTimeExpire = 10 * time.Second
)

const (
	LockUserFarmsOpsTimeExpire = 10 * time.Second
	// 	LockUserFarmsOpsKey 用户操作农场锁key {user_id} {ops_farm_id}
	LockUserFarmsOpsKey = "{lock}:user_%d_ops_farms_%d"
)

const (
	// 	LockUserFarmsOpsKey 用户操作农场锁key {user_id} {product_id}
	LockUserOpsStoreProductKey = "{lock}:user_%d_ops_product_%d"
)
