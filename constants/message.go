package constants

// MessageType 主动推送消息的消息ID
type MessageType int64

const (
	UserNotificationInfoMsgType  MessageType = 800 // 用户通知结构---公告信息body
	UserBalanceNoticeInfoMsgType MessageType = 801 // 用户通知结构---余额信息body

)

func (m MessageType) Int64() int64 {
	return int64(m)
}

// PushMessagePathType 表示主动推送消息的path
type PushMessagePathType string

const (
	PushUserMessageType      PushMessagePathType = "/internal/send"      // 推送给指定用户
	PushBroadcastMessageType PushMessagePathType = "/internal/broadcast" // 广播
)

func (p PushMessagePathType) Str() string {
	return string(p)
}
