package constants

type SettingsConfCategory string

const (
	BasicSettingsCategoryConf    SettingsConfCategory = "basic_settings" // 基础配置
	AgentRebateRulesCategoryConf SettingsConfCategory = "rebate_rules"   // 代理规则
	AgreementCategoryConf        SettingsConfCategory = "agreement"      // 协议
	FundsPolicyCategoryConf      SettingsConfCategory = "funds_policy"   // 充值提款
	OperationConfigCategory      SettingsConfCategory = "operation"      // 操作配置
	ExchangeRateCategory         SettingsConfCategory = "exchange_rate"  // 汇率设置
)

func (t SettingsConfCategory) Str() string {
	return string(t)
}

type SettingsConfKey string

const (
	BasicSettingsKeyConf            SettingsConfKey = "basic_settings"
	AgentRebateRulesKeyConf         SettingsConfKey = "rebate_rules"
	AgreementUserAgreementKeyConf   SettingsConfKey = "user_agreement"
	AgreementPrivacyPolicyKeyConf   SettingsConfKey = "privacy_policy"
	AgreementCustomerServiceKeyConf SettingsConfKey = "customer_service"
	FundsDepositPolicyKeyConf       SettingsConfKey = "deposit_policy"
	FundsWithdrawalPolicyKeyConf    SettingsConfKey = "withdrawal_policy"
	OperationAppPackageConfig       SettingsConfKey = "app_package_config" // app包配置
	ExchangeRateCategoryKeyConf     SettingsConfKey = "exchange_rate"      // app包配置

)

func (t SettingsConfKey) Str() string {
	return string(t)
}

// CurrencyType 定义货币类型的枚举类型
type CurrencyType int64

// 定义货币类型的枚举值
const (
	CurrencyTypeFiat     CurrencyType = iota + 1 // 1 - 法定货币
	CurrencyTypeDigital                          // 2 - 数字货币
	CurrencyTypePlatform                         // 3 - 平台币
)

// String 方法，返回货币类型的中文描述
func (c CurrencyType) String() string {
	switch c {
	case CurrencyTypeFiat:
		return "法定货币"
	case CurrencyTypeDigital:
		return "数字货币"
	case CurrencyTypePlatform:
		return "平台币"
	default:
		return "未知货币类型"
	}
}
