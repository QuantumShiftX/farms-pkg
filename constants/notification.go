package constants

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/logx"
	"regexp"
	"strings"
	"sync"
	"time"
)

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

// NotificationTypeInt 定义整数通知类型
type NotificationTypeInt int64

func (t NotificationTypeInt) Int64() int64 {
	return int64(t)
}

const (
	// NotificationTypeUnspecified 未指定类型
	NotificationTypeUnspecified NotificationTypeInt = 0
	// NotificationTypeDailyGreeting 日常问候类通知
	NotificationTypeDailyGreeting NotificationTypeInt = 1000
	// NotificationTypeOperationTip 游戏操作提示类通知
	NotificationTypeOperationTip NotificationTypeInt = 2000
	// NotificationTypeFriendInfo 好友互动通知
	NotificationTypeFriendInfo NotificationTypeInt = 3000
	// NotificationTypeSystemInfo 系统通知
	NotificationTypeSystemInfo NotificationTypeInt = 4000
)

// NotificationSubTypeInt 定义整数通知子类型
type NotificationSubTypeInt int64

func (t NotificationSubTypeInt) Int64() int64 {
	return int64(t)
}

const (
	// NotificationSubTypeUnspecified 未指定子类型
	NotificationSubTypeUnspecified NotificationSubTypeInt = 0

	// 日常问候子类型 (1000-1999范围)
	// NotificationSubTypeMorning 早上问候
	NotificationSubTypeMorning NotificationSubTypeInt = 1001
	// NotificationSubTypeNoon 中午问候
	NotificationSubTypeNoon NotificationSubTypeInt = 1002
	// NotificationSubTypeEvening 晚上问候
	NotificationSubTypeEvening NotificationSubTypeInt = 1003
	// NotificationSubTypeReturnPlayer 老玩家回归问候
	NotificationSubTypeReturnPlayer NotificationSubTypeInt = 1004

	// 操作提示子类型 (2000-2999范围)
	// NotificationSubTypeCropNeedWater 农作物待浇水
	NotificationSubTypeCropNeedWater NotificationSubTypeInt = 2001
	// NotificationSubTypeCropCanHarvest 农作物待收获
	NotificationSubTypeCropCanHarvest NotificationSubTypeInt = 2002
	// NotificationSubTypeSeedWaitPlant 种子待种植
	NotificationSubTypeSeedWaitPlant NotificationSubTypeInt = 2003
	// NotificationSubTypeCropNeedFertilize 农作物待施肥
	NotificationSubTypeCropNeedFertilize NotificationSubTypeInt = 2004
	// NotificationSubTypeBuySeed 购买种子
	NotificationSubTypeBuySeed NotificationSubTypeInt = 2005

	// 好友信息子类型 (3000-3999范围)
	// NotificationSubTypeFriendRegister 新好友注册
	NotificationSubTypeFriendRegister NotificationSubTypeInt = 3001
	// NotificationSubTypeFriendWatered 好友浇水
	NotificationSubTypeFriendWatered NotificationSubTypeInt = 3002
	// NotificationSubTypeFriendVisit 好友查看农场
	NotificationSubTypeFriendVisit NotificationSubTypeInt = 3003
	// NotificationSubTypeFriendHarvest 好友收获
	NotificationSubTypeFriendHarvest NotificationSubTypeInt = 3004
	// NotificationSubTypeFriendBuySeed 好友购买种子
	NotificationSubTypeFriendBuySeed NotificationSubTypeInt = 3005
	// NotificationSubTypeFriendPlant 好友种植
	NotificationSubTypeFriendPlant NotificationSubTypeInt = 3006

	// 系统信息子类型 (4000-4999范围)
	// NotificationSubTypeDepositSuccess 存款成功
	NotificationSubTypeDepositSuccess NotificationSubTypeInt = 4001
	// NotificationSubTypeWithdrawAvailable 可以提款
	NotificationSubTypeWithdrawAvailable NotificationSubTypeInt = 4002
	// NotificationSubTypeVipUpgrade VIP等级提升
	NotificationSubTypeVipUpgrade NotificationSubTypeInt = 4003
	// NotificationSubTypeRankingBreak 打破排行榜记录
	NotificationSubTypeRankingBreak NotificationSubTypeInt = 4004
)

// 字符串类型转整数类型的映射
var notificationTypeToInt = map[NotificationType]NotificationTypeInt{
	DailyGreeting: NotificationTypeDailyGreeting,
	OperationTip:  NotificationTypeOperationTip,
	FriendInfo:    NotificationTypeFriendInfo,
	SystemInfo:    NotificationTypeSystemInfo,
}

// 整数类型转字符串类型的映射
var notificationIntToType = map[NotificationTypeInt]NotificationType{
	NotificationTypeDailyGreeting: DailyGreeting,
	NotificationTypeOperationTip:  OperationTip,
	NotificationTypeFriendInfo:    FriendInfo,
	NotificationTypeSystemInfo:    SystemInfo,
}

// 子类型字符串转整数的映射
var notificationSubTypeToInt = map[NotificationType]map[NotificationSubType]NotificationSubTypeInt{
	DailyGreeting: {
		DailyGreetingMorning:      NotificationSubTypeMorning,
		DailyGreetingNoon:         NotificationSubTypeNoon,
		DailyGreetingEvening:      NotificationSubTypeEvening,
		DailyGreetingReturnPlayer: NotificationSubTypeReturnPlayer,
	},
	OperationTip: {
		OperationTipCropNeedWater:     NotificationSubTypeCropNeedWater,
		OperationTipCropCanHarvest:    NotificationSubTypeCropCanHarvest,
		OperationTipSeedWaitPlant:     NotificationSubTypeSeedWaitPlant,
		OperationTipCropNeedFertilize: NotificationSubTypeCropNeedFertilize,
		OperationTipBuySeed:           NotificationSubTypeBuySeed,
	},
	FriendInfo: {
		FriendInfoRegister: NotificationSubTypeFriendRegister,
		FriendInfoWatered:  NotificationSubTypeFriendWatered,
		FriendInfoVisit:    NotificationSubTypeFriendVisit,
		FriendInfoHarvest:  NotificationSubTypeFriendHarvest,
		FriendInfoBuySeed:  NotificationSubTypeFriendBuySeed,
		FriendInfoPlant:    NotificationSubTypeFriendPlant,
	},
	SystemInfo: {
		SystemInfoDepositSuccess:    NotificationSubTypeDepositSuccess,
		SystemInfoWithdrawAvailable: NotificationSubTypeWithdrawAvailable,
		SystemInfoVipUpgrade:        NotificationSubTypeVipUpgrade,
		SystemInfoRankingBreak:      NotificationSubTypeRankingBreak,
	},
}

// 子类型整数转字符串的映射
var notificationSubTypeIntToSubType = map[NotificationSubTypeInt]NotificationSubType{
	// 日常问候子类型
	NotificationSubTypeMorning:      DailyGreetingMorning,
	NotificationSubTypeNoon:         DailyGreetingNoon,
	NotificationSubTypeEvening:      DailyGreetingEvening,
	NotificationSubTypeReturnPlayer: DailyGreetingReturnPlayer,

	// 操作提示子类型
	NotificationSubTypeCropNeedWater:     OperationTipCropNeedWater,
	NotificationSubTypeCropCanHarvest:    OperationTipCropCanHarvest,
	NotificationSubTypeSeedWaitPlant:     OperationTipSeedWaitPlant,
	NotificationSubTypeCropNeedFertilize: OperationTipCropNeedFertilize,
	NotificationSubTypeBuySeed:           OperationTipBuySeed,

	// 好友信息子类型
	NotificationSubTypeFriendRegister: FriendInfoRegister,
	NotificationSubTypeFriendWatered:  FriendInfoWatered,
	NotificationSubTypeFriendVisit:    FriendInfoVisit,
	NotificationSubTypeFriendHarvest:  FriendInfoHarvest,
	NotificationSubTypeFriendBuySeed:  FriendInfoBuySeed,
	NotificationSubTypeFriendPlant:    FriendInfoPlant,

	// 系统信息子类型
	NotificationSubTypeDepositSuccess:    SystemInfoDepositSuccess,
	NotificationSubTypeWithdrawAvailable: SystemInfoWithdrawAvailable,
	NotificationSubTypeVipUpgrade:        SystemInfoVipUpgrade,
	NotificationSubTypeRankingBreak:      SystemInfoRankingBreak,
}

// 子类型整数与主类型的映射关系
var subTypeIntToMainTypeInt = map[NotificationSubTypeInt]NotificationTypeInt{
	// 日常问候子类型
	NotificationSubTypeMorning:      NotificationTypeDailyGreeting,
	NotificationSubTypeNoon:         NotificationTypeDailyGreeting,
	NotificationSubTypeEvening:      NotificationTypeDailyGreeting,
	NotificationSubTypeReturnPlayer: NotificationTypeDailyGreeting,

	// 操作提示子类型
	NotificationSubTypeCropNeedWater:     NotificationTypeOperationTip,
	NotificationSubTypeCropCanHarvest:    NotificationTypeOperationTip,
	NotificationSubTypeSeedWaitPlant:     NotificationTypeOperationTip,
	NotificationSubTypeCropNeedFertilize: NotificationTypeOperationTip,
	NotificationSubTypeBuySeed:           NotificationTypeOperationTip,

	// 好友信息子类型
	NotificationSubTypeFriendRegister: NotificationTypeFriendInfo,
	NotificationSubTypeFriendWatered:  NotificationTypeFriendInfo,
	NotificationSubTypeFriendVisit:    NotificationTypeFriendInfo,
	NotificationSubTypeFriendHarvest:  NotificationTypeFriendInfo,
	NotificationSubTypeFriendBuySeed:  NotificationTypeFriendInfo,
	NotificationSubTypeFriendPlant:    NotificationTypeFriendInfo,

	// 系统信息子类型
	NotificationSubTypeDepositSuccess:    NotificationTypeSystemInfo,
	NotificationSubTypeWithdrawAvailable: NotificationTypeSystemInfo,
	NotificationSubTypeVipUpgrade:        NotificationTypeSystemInfo,
	NotificationSubTypeRankingBreak:      NotificationTypeSystemInfo,
}

// StringTypeToInt 将字符串类型转换为整数类型
func StringTypeToInt(t NotificationType) NotificationTypeInt {
	if val, ok := notificationTypeToInt[t]; ok {
		return val
	}
	return NotificationTypeUnspecified
}

// IntTypeToString 将整数类型转换为字符串类型
func IntTypeToString(t NotificationTypeInt) NotificationType {
	if val, ok := notificationIntToType[t]; ok {
		return val
	}
	return ""
}

// StringSubTypeToInt 将字符串子类型转换为整数子类型
func StringSubTypeToInt(mainType NotificationType, subType NotificationSubType) NotificationSubTypeInt {
	if mainTypeMap, ok := notificationSubTypeToInt[mainType]; ok {
		if val, ok := mainTypeMap[subType]; ok {
			return val
		}
	}
	return NotificationSubTypeUnspecified
}

// IntSubTypeToString 将整数子类型转换为字符串子类型
func IntSubTypeToString(subTypeInt NotificationSubTypeInt) NotificationSubType {
	if val, ok := notificationSubTypeIntToSubType[subTypeInt]; ok {
		return val
	}
	return ""
}

// GetMainTypeFromSubTypeInt 从子类型整数获取主类型整数
func GetMainTypeFromSubTypeInt(subTypeInt NotificationSubTypeInt) NotificationTypeInt {
	if val, ok := subTypeIntToMainTypeInt[subTypeInt]; ok {
		return val
	}
	return NotificationTypeUnspecified
}

// GetMainTypeStrFromSubTypeInt 从子类型整数获取主类型字符串
func GetMainTypeStrFromSubTypeInt(subTypeInt NotificationSubTypeInt) NotificationType {
	mainTypeInt := GetMainTypeFromSubTypeInt(subTypeInt)
	return IntTypeToString(mainTypeInt)
}

// GetAllTypeMapping 获取所有类型映射关系
func GetAllTypeMapping() map[string]interface{} {
	return map[string]interface{}{
		"notificationTypeToInt":           notificationTypeToInt,
		"notificationIntToType":           notificationIntToType,
		"notificationSubTypeToInt":        notificationSubTypeToInt,
		"notificationSubTypeIntToSubType": notificationSubTypeIntToSubType,
		"subTypeIntToMainTypeInt":         subTypeIntToMainTypeInt,
	}
}

// IsValidNotificationType 判断主类型代码是否合法
func IsValidNotificationType(typeInt int64) bool {
	// 转换为NotificationTypeInt类型
	mainType := NotificationTypeInt(typeInt)

	// 检查是否在预定义的映射中
	exists := IntTypeToString(mainType)
	return exists != "" && mainType != NotificationTypeUnspecified
}

// IsValidNotificationSubType 判断子类型代码是否合法
func IsValidNotificationSubType(subTypeInt int64) bool {
	// 转换为NotificationSubTypeInt类型
	subType := NotificationSubTypeInt(subTypeInt)

	// 检查是否在预定义的映射中
	exists := IntSubTypeToString(subType)
	return exists != "" && subType != NotificationSubTypeUnspecified
}

// IsValidNotificationTypePair 判断主类型和子类型的组合是否合法
func IsValidNotificationTypePair(mainTypeInt, subTypeInt int64) bool {
	// 转换为对应的类型
	mainType := NotificationTypeInt(mainTypeInt)
	subType := NotificationSubTypeInt(subTypeInt)

	// 先检查主类型和子类型各自是否合法
	if !IsValidNotificationType(mainTypeInt) || !IsValidNotificationSubType(subTypeInt) {
		return false
	}

	// 获取子类型对应的主类型
	expectedMainType := GetMainTypeFromSubTypeInt(subType)

	// 判断子类型是否属于指定的主类型
	return expectedMainType == mainType
}

// GetNotificationTypeNamePair 获取主类型和子类型的名称对，如果不合法返回空字符串
func GetNotificationTypeNamePair(mainTypeInt, subTypeInt int64) (mainTypeName, subTypeName string) {
	// 判断组合是否合法
	if !IsValidNotificationTypePair(mainTypeInt, subTypeInt) {
		return "", ""
	}

	// 转换为对应的类型
	mainType := NotificationTypeInt(mainTypeInt)
	subType := NotificationSubTypeInt(subTypeInt)

	// 获取类型名称
	mainTypeStr := IntTypeToString(mainType)
	subTypeStr := IntSubTypeToString(subType)

	return mainTypeStr.ToStr(), subTypeStr.ToStr()
}

// TemplateEngine 根据不同类型替换文本中的变量
// 支持各种数据类型的变量替换，包括字符串、数字、布尔值、列表等
func TemplateEngine(template string, variables map[string]interface{}) string {
	// 如果模板为空或没有变量，直接返回
	if template == "" || len(variables) == 0 {
		return template
	}

	// 快速检查是否包含变量标记，避免不必要的正则处理
	if !strings.Contains(template, "${") {
		return template
	}

	// 编译正则表达式（预编译提高性能）
	var reOnce sync.Once
	var re *regexp.Regexp
	reOnce.Do(func() {
		re = regexp.MustCompile(`\${([^}]+)}`)
	})

	// 替换所有匹配的变量
	result := re.ReplaceAllStringFunc(template, func(match string) string {
		// 提取变量名（去掉 ${ 和 }）
		varName := match[2 : len(match)-1]

		// 从变量映射中获取值
		value, exists := variables[varName]
		if !exists {
			// 如果变量不存在，保留原始变量表达式
			return match
		}

		// 根据变量类型进行适当的转换
		var resultStr string
		switch v := value.(type) {
		case string:
			resultStr = v
		case int, int32, int64, uint, uint32, uint64:
			resultStr = fmt.Sprintf("%d", v)
		case float32, float64:
			resultStr = fmt.Sprintf("%.2f", v)
		case bool:
			resultStr = fmt.Sprintf("%t", v)
		case []string:
			resultStr = strings.Join(v, ", ")
		case []interface{}:
			// 处理复杂类型的数组
			var items []string
			for _, item := range v {
				items = append(items, fmt.Sprintf("%v", item))
			}
			resultStr = strings.Join(items, ", ")
		case map[string]interface{}:
			// 处理嵌套对象，转换为JSON字符串
			jsonStr, err := jsonx.Marshal(v)
			if err != nil {
				logx.Errorf("转换JSON失败: %v", err)
				return match
			}
			resultStr = string(jsonStr)
		case time.Time:
			// 时间类型特殊处理
			resultStr = v.Format(time.DateTime)
		case nil:
			// 空值处理
			resultStr = ""
		default:
			// 其他类型尝试使用默认字符串表示
			resultStr = fmt.Sprintf("%v", v)
		}

		return resultStr
	})

	return result
}

// 定义状态枚举
type AnnouncementDisplayStatus int

const (
	Pending    AnnouncementDisplayStatus = iota + 1 // 待展示
	Displaying                                      // 展示中
	Expired                                         // 已过期
)

func (t AnnouncementDisplayStatus) Int8() int8 {
	return int8(t)
}
