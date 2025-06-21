package cachekey

import "time"

/**
 * {cache}标识，redis集群时slot要根据{xxx}落槽
 */

const (
	// CacheGrowthRuleKey 成长规则信息
	CacheGrowthRuleKey = "{cache}:growth_rule:%d" // {cache}:growth_rule:{ActionType}

	// CacheCropsKey 农作物信息
	CacheCropsKey = "{cache}:crops:%d" // {cache}:crops:{id}

	// CacheSettingBaseKey 系统设置信息[通用key system_setting]
	CacheSettingBaseKey = "{cache}:system_setting:%s:%s" // {cache}:system_setting:{category}{key}

	// CacheSettingRebateKey 推广返佣信息
	CacheSettingRebateKey = "{cache}:system_rebate:%s:%s" // {cache}:system_setting:{category}{key}

	// CacheSettingAgreementKey 协议信息
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

/**
 * {cache}标识，redis集群时slot要根据{xxx}落槽
 */

const (
	// SystemOssSignedURLTimeExpire 图片签名过期时间
	SystemOssSignedURLTimeExpire = 24*7*time.Hour - 60*time.Minute
	// SystemOssSignedURLKey 图片签名
	SystemOssSignedURLKey = "{oss:signed:url:%s}"
)

/**
 * {cache}标识，redis集群时slot要根据{xxx}落槽
 */

const (
	// SystemsRateTimeExpire 过期时间
	SystemsRateTimeExpire = 1 * time.Minute
	// SystemsRateBaseKey 基准汇率 官方/参考汇率
	SystemsRateBaseKey = "{rate:base:coins_usdt}"
	// SystemsRateMerchantKey 各币商汇率
	SystemsRateMerchantKey = "{rate:merchant:%d:coins_usdt}"
)
