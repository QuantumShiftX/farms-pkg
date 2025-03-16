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

	CacheSettingAgreementKey = "{cache}:system_agreement:%s" // {cache}:system_setting:{category}

)

/**
 * {cache}标识，redis集群时slot要根据{xxx}落槽
 */

const (
	// SystemVIPLevelKey vip in
	SystemVIPLevelKey = "{cache}:vip_level:info"
)

/**
 * {cache}标识，redis集群时slot要根据{xxx}落槽
 */

const (
	// NotifyTypeCodeLanguageKey 通知缓存（{type_code}:{language_code}）
	NotifyTypeCodeLanguageKey = "{cache}:notify:%s:info:%s"
	// NotifySubTypeCodeLanguageKey 通知缓存（{sub_type_code}:{language_code}）
	NotifySubTypeCodeLanguageKey = "{cache}:notify_sub:%s:info:%s"
)
