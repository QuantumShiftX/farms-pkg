package constants

// MessageType 主动推送消息的消息ID
type MessageType int64

const (
	UserPersonalInfoNoticeMsgType   MessageType = 3   // 用户通知结构---用户个人信息变更
	UserNotificationInfoMsgType     MessageType = 800 // 用户通知结构---公告信息body
	UserBalanceNoticeInfoMsgType    MessageType = 801 // 用户通知结构---余额信息body
	UserLandStatusNoticeInfoMsgType MessageType = 802 // 用户通知结构---用户土地变更相关
	UserLandCropCanOpsMsgType       MessageType = 803 // 用户通知结构---用户土地可收获状态
	UserFortuneTreeStatusMsgType    MessageType = 804 // 用户通知结构---用户发财树状态变更数据推送

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
