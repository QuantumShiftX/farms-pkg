package types

import (
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// Message 代表消息结构
type Message struct {
	MessageID int64   `json:"message_id"` // 消息ID
	UserIDs   []int64 `json:"user_ids"`   // 目标用户ID列表
	Body      []byte  `json:"body"`       // Proto序列化后的数据
}

// ProtoMessage 实现proto.Message接口的标记方法
func (m *Message) ProtoMessage() {}

// ProtoReflect 返回protoreflect.Message接口，实现Protocol Buffers反射
func (m *Message) ProtoReflect() protoreflect.Message {
	// 创建一个实现了protoreflect.Message接口的包装器类型
	return &messageReflect{
		message: m,
	}
}

// messageReflect 是一个实现了protoreflect.Message的包装器类型
type messageReflect struct {
	protoreflect.Message
	message *Message
}

// Descriptor 返回消息的描述符
func (mr *messageReflect) Descriptor() protoreflect.MessageDescriptor {
	// 实际项目中应该返回真实的描述符
	// 这里返回nil仅作为示例，真实实现需要生成适当的描述符
	return nil
}

// Type 返回消息的类型
func (mr *messageReflect) Type() protoreflect.MessageType {
	return nil
}

// New 创建一个新的空消息
func (mr *messageReflect) New() protoreflect.Message {
	return (&Message{}).ProtoReflect()
}

// Interface 返回原始消息
func (mr *messageReflect) Interface() protoreflect.ProtoMessage {
	return mr.message
}

// Range 遍历所有字段
func (mr *messageReflect) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	// 为每个字段调用函数f
	// 这里需要为每个字段创建适当的描述符和值
	// 简化实现，实际项目中需要更完整的实现
}

// Has 检查字段是否存在
func (mr *messageReflect) Has(fd protoreflect.FieldDescriptor) bool {
	// 根据字段描述符检查字段是否存在
	return false
}

// Clear 清除字段
func (mr *messageReflect) Clear(fd protoreflect.FieldDescriptor) {
	// 根据字段描述符清除字段
}

// Get 获取字段值
func (mr *messageReflect) Get(fd protoreflect.FieldDescriptor) protoreflect.Value {
	// 根据字段描述符获取字段值
	return protoreflect.ValueOfString("")
}

// Set 设置字段值
func (mr *messageReflect) Set(fd protoreflect.FieldDescriptor, v protoreflect.Value) {
	// 根据字段描述符设置字段值
}

// Mutable 获取可变字段
func (mr *messageReflect) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	// 获取可变字段
	return protoreflect.ValueOfString("")
}

// 序列化 Message 为 Proto 格式
func (m *Message) Serialize() ([]byte, error) {
	return proto.Marshal(m)
}

// 反序列化 Proto 数据到 Message
func Deserialize(data []byte) (*Message, error) {
	var msg Message
	err := proto.Unmarshal(data, &msg)
	return &msg, err
}
