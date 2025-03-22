package cachekey

/**
 * {cache}标识，redis集群时slot要根据{xxx}落槽
 */

const (
	// NotificationExtraKey 代表通知相关的信息 Redis Key 模板。
	// 该键的格式如下：
	// {cache}:notification_growth_rule:{typeCode}:{subTypeCode}:{userID}:{extraKey}:{dateKey}
	NotificationExtraKey = "{cache}:notification_growth_rule:%s:%s:%d:%s:%s"

	// NotificationKey 代表通知相信息 Redis Key 模板。
	// 该键的格式如下：
	// {cache}:notification_crops:{typeCode}:{subTypeCode}:{userID}:{dateKey}
	NotificationKey = "{cache}:notification_crops:%s:%s:%d:%s"
)
