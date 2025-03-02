package cachekey

import "time"

/**
 * {cache}标识，redis集群时slot要根据{xxx}落槽
 */

const (
	CacheCommonKeyExpire = time.Minute * 120

	CacheCaptchaKeyExpire = time.Minute * 15
)
