package cachekey

/**
 * {cache}标识，redis集群时slot要根据{xxx}落槽
 */

const (
	// CacheVipLevelConfKey 会员等级配置
	CacheVipLevelConfKey = "{cache}:vip_level_rule:%d" // {cache}:vip_level_rule:{id}

)
