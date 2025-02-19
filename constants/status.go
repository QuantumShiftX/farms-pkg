package constants

// StatusEnabledType 通用status状态
type StatusEnabledType int64

const (
	_ StatusEnabledType = iota //
	// StatusEnabled 启用
	StatusEnabled
	// StatusDisabled 禁用
	StatusDisabled
)

func (t StatusEnabledType) Int64() int64 {
	return int64(t)
}

func (t StatusEnabledType) Default() StatusEnabledType {
	if !(t == StatusEnabled || t == StatusDisabled) {
		return StatusEnabled
	}
	return t
}
