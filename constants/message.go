package constants

// PushMessagePathType 表示主动推送消息的path
type PushMessagePathType string

const (
	PushUserMessageType      PushMessagePathType = "/internal/send"      // 推送给指定用户
	PushBroadcastMessageType PushMessagePathType = "/internal/broadcast" // 广播
)

func (p PushMessagePathType) Str() string {
	return string(p)
}
