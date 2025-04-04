package constants

const (
	// SuperAdmin 超级管理员标识
	SuperAdmin = "super_admin"
)

// 定义客服类型枚举
type CustomerServiceType int8

const (
	TelegramService   CustomerServiceType = iota + 1 // Telegram 客服
	WhatsappService                                  // Whatsapp 客服
	LineService                                      // Line 客服
	ThirdPartyService                                // 三方客服
)

func (c CustomerServiceType) Int8() int8 {
	return int8(c)
}

// 转换枚举值为字符串
func (c CustomerServiceType) String() string {
	switch c {
	case TelegramService:
		return "Telegram客服"
	case WhatsappService:
		return "Whatsapp客服"
	case LineService:
		return "Line客服"
	case ThirdPartyService:
		return "三方客服"
	default:
		return "未知客服类型"
	}
}

// WalletType 钱包类型常量
const (
	WalletTypeReceive int8 = 1 // 收款钱包
	WalletTypePayment int8 = 2 // 打款钱包
)

// TransactionType 交易类型常量
const (
	TransactionTypeDeposit    int8 = 1 // 入款
	TransactionTypeWithdrawal int8 = 2 // 出款
	TransactionTypeCollect    int8 = 3 // 归集
)

// TransactionStatus 交易状态常量
const (
	TransactionStatusSuccess    int8 = 1 // 成功
	TransactionStatusFailed     int8 = 2 // 失败
	TransactionStatusProcessing int8 = 3 // 处理中
)
