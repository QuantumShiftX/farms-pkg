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

func (t StatusEnabledType) Int8() int8 {
	return int8(t)
}

func (t StatusEnabledType) Default() StatusEnabledType {
	if !(t == StatusEnabled || t == StatusDisabled) {
		return StatusEnabled
	}
	return t
}

// 定义在线离线状态的枚举类型
type UserStatus int

const (
	UserStatusOnline  UserStatus = iota + 1 // 在线状态 1
	UserStatusOffline                       // 离线状态 2
)

func (t UserStatus) Int64() int64 {
	return int64(t)
}

func (t UserStatus) Int8() int8 {
	return int8(t)
}

// 定义是否可操作的枚举
type OperableEnum int

const (
	// 可操作
	Operable OperableEnum = 1

	// 不可操作
	UnOperable OperableEnum = 2
)

// 转换枚举为字符串的映射（可选）
var OperableEnumStrings = map[OperableEnum]string{
	Operable:   "可操作",
	UnOperable: "不可操作",
}

func (o OperableEnum) Int64() int64 {
	return int64(o)
}

// 示例函数：输出是否可操作
func (o OperableEnum) String() string {
	return OperableEnumStrings[o]
}
