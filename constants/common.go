package constants

const Trace = "trace_id"

// Endpoint 终端
type Endpoint int

const (
	_                  Endpoint = iota
	EndpointPC                  // PC
	EndpointAndroidH5           // Android_H5
	EndpointIOSH5               // IOS_H5
	EndpointAndroidApp          // Android_APP
	EndpointIOSApp              // IOS_APP
)

// TimeSearchType 0-全部，1-今日，2-昨日，3-本周，4-上周，5-本月，6-上月，7-自定义
type TimeSearchType int

const (
	TimeSearchTypeAll       TimeSearchType = iota // ALL
	TimeSearchTypeToday                           // 今天
	TimeSearchTypeYesterday                       // 昨天
	TimeSearchTypeThisWeek                        // 本周
	TimeSearchTypeLastWeek                        // 上周
	TimeSearchTypeThisMonth                       // 本月
	TimeSearchTypeLastMonth                       // 上月
	TimeSearchTypeCustom                          // 自定义
)
