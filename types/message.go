package types

// Message 代表消息结构
type Message struct {
	MessageID int64   `json:"message_id"` // 消息ID
	UserID    []int64 `json:"user_id"`    // 目标用户ID列表
	Body      []byte  `json:"body"`       // Proto序列化后的数据
}
