package constants

// TagType represents the type of a tag (system or custom)
type TagType int8

func (t TagType) Int8() int8 {
	return int8(t)
}

const (
	// TagTypeSystem represents a system tag
	TagTypeSystem TagType = 1
	// TagTypeCustom represents a custom user-defined tag
	TagTypeCustom TagType = 2
)

type LandStatus int64

func (t LandStatus) Int64() int64 {
	return int64(t)
}

// 土地状态常量
const (
	_                  LandStatus = iota
	LandStatusUnopened            // 未开垦
	LandStatusOpened              // 已开垦未种植
	LandStatusPlanted             // 已种植
)

type RegisterType int64

func (t RegisterType) Int64() int64 {
	return int64(t)
}

// 注册方式常量
const (
	_                    RegisterType = iota
	RegisterTypeAccount               // 账号注册
	RegisterTypePhone                 // 手机号注册
	RegisterTypeEmail                 // 邮箱注册
	RegisterTypeFacebook              // Facebook注册
	RegisterTypeGoogle                // Google注册
	RegisterTypeLine                  // Line注册
)

type RegisterSource int64

func (t RegisterSource) Int64() int64 {
	return int64(t)
}

// 注册来源常量
const (
	_                       RegisterSource = iota
	RegisterSourceNone                     // 无渠道
	RegisterSourceUnknown                  // 未知渠道
	RegisterSourcePromotion                // 推广注册
	RegisterSourceAdmin                    // 后台添加
	RegisterSourceAPI                      // API注册
	RegisterSourceTikTok                   // TikTok
)

// FreeGiftStatus 新用户免费商品领取状态
type FreeGiftStatus int

const (
	FreeGiftNotClaimed FreeGiftStatus = 1 // 未领取
	FreeGiftClaimed    FreeGiftStatus = 2 // 已领取
)

// ToInt8 返回FreeGiftStatus的int8表示
func (s FreeGiftStatus) ToInt8() int8 {
	return int8(s)
}

// TutorialStatus 新手教程参与状态
type TutorialStatus int

const (
	TutorialNotCompleted TutorialStatus = 1 // 未参加
	TutorialCompleted    TutorialStatus = 2 // 已参加
)

// ToInt8 返回TutorialStatus的int8表示
func (s TutorialStatus) ToInt8() int8 {
	return int8(s)
}

// DefaultAddressStatus 默认存款地址状态
type DefaultAddressStatus int

const (
	NotDefaultAddress DefaultAddressStatus = 1 // 否
	IsDefaultAddress  DefaultAddressStatus = 2 // 是
)

// ToInt8 返回DefaultAddressStatus的int8表示
func (s DefaultAddressStatus) ToInt8() int8 {
	return int8(s)
}

// AddressUsageType 地址用途类型
type AddressUsageType int

const (
	DepositOnly        AddressUsageType = 1 // 仅存款
	WithdrawOnly       AddressUsageType = 2 // 仅提款
	DepositAndWithdraw AddressUsageType = 3 // 存款和提款
)

// ToInt8 返回AddressUsageType的int8表示
func (t AddressUsageType) ToInt8() int8 {
	return int8(t)
}

// WalletType 钱包类型
type WalletType int

const (
	UserWallet   WalletType = 1 // 用户钱包
	AgentWallet  WalletType = 2 // 代理钱包
	RewardWallet WalletType = 3 // 奖励钱包
	USDTWallet   WalletType = 4 // USDT钱包
)

// String 返回WalletType的字符串表示
func (t WalletType) String() string {
	switch t {
	case UserWallet:
		return "用户钱包"
	case AgentWallet:
		return "代理钱包"
	case RewardWallet:
		return "奖励钱包"
	case USDTWallet:
		return "USDT钱包"
	default:
		return "未知钱包类型"
	}
}

// ToInt8 返回WalletType的int8表示
func (t WalletType) ToInt8() int8 {
	return int8(t)
}

// VerifyType 验证方式
type VerifyType int

const (
	VerifyTypePwd        VerifyType = 1 // 密码验证
	VerifyTypeSms        VerifyType = 2 // 短信验证
	VerifyTypeEmail      VerifyType = 3 // 邮件验证
	VerifyTypeGoogleAuth VerifyType = 4 // 谷歌验证
)

func (s VerifyType) Int64() int64 {
	return int64(s)
}

// LoginType 登录方式
type LoginType int

const (
	_                 LoginType = iota
	LoginTypeUsername           // 用户名登录
	LoginTypePhone              // 手机号登录
	LoginTypeEmail              // 邮件登录
	LoginTypeOther              // 第三方登录
)

func (s LoginType) Int64() int64 {
	return int64(s)
}

// RewardType 定义奖励类型枚举
type RewardType int16

const (
	_              RewardType = iota //
	RewardTypeCoin                   // 金币
	RewardTypeItem                   // 道具
	RewardTypeExp                    // 经验
	RewardTypeLand                   // 土地
)

func (t RewardType) Int64() int64 {
	return int64(t)
}

func (t RewardType) Int16() int16 {
	return int16(t)
}
