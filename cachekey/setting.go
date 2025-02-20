package cachekey

/**
 * {cache}标识，redis集群时slot要根据{xxx}落槽
 */

const (
	// CacheGrowthRuleKey 成长规则信息
	CacheGrowthRuleKey = "{cache}:growth_rule:%d" // {cache}:growth_rule:{ActionType}

	// CacheCropsKey 农作物信息
	CacheCropsKey = "{cache}:crops:%d" // {cache}:crops:{id}

	CacheSettingBaseKey = "{cache}:system_setting:%s:%s" // {cache}:system_setting:{category}{key}

	CacheSettingRebateKey = "{cache}:system_rebate:%s:%s" // {cache}:system_setting:{category}{key}

)
