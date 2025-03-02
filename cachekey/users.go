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
	CacheOnlineUsersKey    = "{cache}:online:users"   // {cache}:online:users
	CacheUserOnlineInfoKey = "{cache}:user:online:%d" // {cache}:user:online:{user_id}

)

const (
	// CacheUserSendCaptchaInfoKey 用户发送验证码
	CacheUserSendCaptchaInfoKey = "{cache}:user:captcha:%s" // {cache}:user:online:{account}

)
