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
