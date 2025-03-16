package errs

import "github.com/QuantumShiftX/golib/xerr"

const (
	ConfigErrCode xerr.ErrCode = 90000

	ConfigNotFoundError   = iota + ConfigErrCode // 配置不存在
	ConfigInvalidError                           // 配置无效
	ConfigUpdateError                            // 更新配置失败
	ConfigPermissionError                        // 无权修改配置
	ConfigFormatError                            // 配置格式错误
)

var (
	ErrConfigNotFound   = xerr.New(ConfigNotFoundError, "配置不存在")
	ErrConfigInvalid    = xerr.New(ConfigInvalidError, "配置无效")
	ErrConfigUpdate     = xerr.New(ConfigUpdateError, "更新配置失败")
	ErrConfigPermission = xerr.New(ConfigPermissionError, "无权修改配置")
	ErrConfigFormat     = xerr.New(ConfigFormatError, "配置格式错误")
)
