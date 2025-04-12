package constants

// CropsType 枚举，表示不同的作物种子
type CropsType int8

// 定义水果常量
const (
	Apple      CropsType = 1 // 苹果
	Orange     CropsType = 2 // 橙子
	Banana     CropsType = 3 // 香蕉
	Cherry     CropsType = 4 // 樱桃
	Avocado    CropsType = 5 // 牛油果
	Strawberry CropsType = 6 // 草莓
	CacaoBean  CropsType = 7 // 可可豆
)

// Map of Fruit to string descriptions
var CropsDescriptions = map[CropsType]string{
	Apple:      "苹果种子",
	Orange:     "橙子种子",
	Banana:     "香蕉种子",
	Cherry:     "樱桃种子",
	Avocado:    "牛油果种子",
	Strawberry: "草莓种子",
	CacaoBean:  "可可豆种子",
}

func (t CropsType) Int8() int8 {
	return int8(t)
}

// ActionType 枚举，表示不同的行为类型
type ActionType int8

// 定义行为类型常量
const (
	Deposit        ActionType = 1 // 存款
	SubDeposit     ActionType = 2 // 下级存款
	Harvest        ActionType = 3 // 收获
	InviteSub      ActionType = 4 // 邀请下级
	LoginDaily     ActionType = 5 // 每日登录
	OnlineDuration ActionType = 6 // 在线时长
)

// Map of ActionType to string descriptions
var ActionTypeDescriptions = map[ActionType]string{
	Deposit:        "存款",
	SubDeposit:     "下级存款",
	Harvest:        "收获",
	InviteSub:      "邀请下级",
	LoginDaily:     "每日登录",
	OnlineDuration: "在线时长",
}

func (t ActionType) Int8() int8 {
	return int8(t)
}

// IsValid 检查 ActionType 是否有效
func (t ActionType) IsValid() bool {
	return t >= Deposit && t <= OnlineDuration
}
