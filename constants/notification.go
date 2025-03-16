package constants

// NotificationType 定义通知类型的枚举
type NotificationType string

func (t NotificationType) ToStr() string {
	return string(t)
}

const (
	// DailyGreeting 日常问候类通知
	DailyGreeting NotificationType = "daily_greeting"
	// OperationTip 游戏操作提示类通知
	OperationTip NotificationType = "operation_tip"
	// FriendInfo 好友互动通知
	FriendInfo NotificationType = "friend_info"
	// SystemInfo 系统通知
	SystemInfo NotificationType = "system_info"
)

// NotificationSubType 定义通知子类型的枚举
type NotificationSubType string

func (t NotificationSubType) ToStr() string {
	return string(t)
}

// DailyGreeting 子类型
const (
	// DailyGreetingMorning 早上问候
	DailyGreetingMorning NotificationSubType = "morning"
	// DailyGreetingNoon 中午问候
	DailyGreetingNoon NotificationSubType = "noon"
	// DailyGreetingEvening 晚上问候
	DailyGreetingEvening NotificationSubType = "evening"
	// DailyGreetingReturnPlayer 老玩家回归问候
	DailyGreetingReturnPlayer NotificationSubType = "return_player"
)

// OperationTip 子类型
const (
	// OperationTipCropNeedWater 农作物待浇水
	OperationTipCropNeedWater NotificationSubType = "crop_need_water"
	// OperationTipCropCanHarvest 农作物待收获
	OperationTipCropCanHarvest NotificationSubType = "crop_can_harvest"
	// OperationTipSeedWaitPlant 种子待种植
	OperationTipSeedWaitPlant NotificationSubType = "seed_wait_plant"
	// OperationTipCropNeedFertilize 农作物待施肥
	OperationTipCropNeedFertilize NotificationSubType = "crop_need_fertilize"
	// OperationTipBuySeed 购买种子
	OperationTipBuySeed NotificationSubType = "buy_seed"
)

// FriendInfo 子类型
const (
	// FriendInfoRegister 新好友注册
	FriendInfoRegister NotificationSubType = "friend_register"
	// FriendInfoWatered 好友浇水
	FriendInfoWatered NotificationSubType = "friend_watered"
	// FriendInfoVisit 好友查看农场
	FriendInfoVisit NotificationSubType = "friend_visit"
	// FriendInfoHarvest 好友收获
	FriendInfoHarvest NotificationSubType = "friend_harvest"
	// FriendInfoBuySeed 好友购买种子
	FriendInfoBuySeed NotificationSubType = "friend_buy_seed"
	// FriendInfoPlant 好友种植
	FriendInfoPlant NotificationSubType = "friend_plant"
)

// SystemInfo 子类型
const (
	// SystemInfoDepositSuccess 存款成功
	SystemInfoDepositSuccess NotificationSubType = "deposit_success"
	// SystemInfoWithdrawAvailable 可以提款
	SystemInfoWithdrawAvailable NotificationSubType = "withdraw_available"
	// SystemInfoVipUpgrade VIP等级提升
	SystemInfoVipUpgrade NotificationSubType = "vip_upgrade"
	// SystemInfoRankingBreak 打破排行榜记录
	SystemInfoRankingBreak NotificationSubType = "ranking_break"
)

// 获取所有通知类型
func GetAllNotificationTypes() []NotificationType {
	return []NotificationType{
		DailyGreeting,
		OperationTip,
		FriendInfo,
		SystemInfo,
	}
}

// 获取指定类型的所有子类型
func GetAllSubTypesByType(notificationType NotificationType) []NotificationSubType {
	switch notificationType {
	case DailyGreeting:
		return []NotificationSubType{
			DailyGreetingMorning,
			DailyGreetingNoon,
			DailyGreetingEvening,
			DailyGreetingReturnPlayer,
		}
	case OperationTip:
		return []NotificationSubType{
			OperationTipCropNeedWater,
			OperationTipCropCanHarvest,
			OperationTipSeedWaitPlant,
			OperationTipCropNeedFertilize,
			OperationTipBuySeed,
		}
	case FriendInfo:
		return []NotificationSubType{
			FriendInfoRegister,
			FriendInfoWatered,
			FriendInfoVisit,
			FriendInfoHarvest,
			FriendInfoBuySeed,
			FriendInfoPlant,
		}
	case SystemInfo:
		return []NotificationSubType{
			SystemInfoDepositSuccess,
			SystemInfoWithdrawAvailable,
			SystemInfoVipUpgrade,
			SystemInfoRankingBreak,
		}
	default:
		return []NotificationSubType{}
	}
}
