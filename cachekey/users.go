package cachekey

/**
 * {cache}标识，redis集群时slot要根据{xxx}落槽
 */

const (
	// CacheVipLevelConfKey 会员等级配置
	CacheVipLevelConfKey = "{cache}:vip_level_rule:%d" // {cache}:vip_level_rule:{id}

)

const (
	// CacheOnlineUsersKey 在线用户key
	CacheOnlineUsersKey    = "{cache_user}:online:users" // {cache}:online:users
	CacheUserOnlineInfoKey = "{cache_user}:online:%d"    // {cache}:user:online:{user_id}

)

const (
	// CacheUserSendCaptchaInfoKey 用户发送验证码
	CacheUserSendCaptchaInfoKey = "{cache}:user:captcha:%s" // {cache}:user:online:{account}

)

const (
	// MerchantExchangeRateInfoKey 用户转换金额缓存
	MerchantExchangeRateInfoKey = "{cache_user_%d}:exchange_%d:ustd" // {user_id}{user_icon}

)

// 常量定义 Redis 键
const (
	UserLoginDailyKey    = "{user_login_%d}:daily:%s" // 用户每日登录记录，{userId}:{date}
	UserLoginStreakKey   = "{user_login_%d}:streak"   // 用户连续登录计数，{userId}
	UserLoginLastDateKey = "{user_login_%d}:last"     // 用户最后登录日期，{userId}
)

const (
	UserLandOperationLogsKey = "{user_%d}:op_land_position_%d:%d:%s" // 用户操作土地位置记录键 {user_用户ID}:op_land_position_位置ID:操作类型:年月日
)
