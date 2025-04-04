package constants

type SettingsConfCategory string

const (
	BasicSettingsCategoryConf    SettingsConfCategory = "basic_settings" // 基础配置
	AgentRebateRulesCategoryConf SettingsConfCategory = "rebate_rules"   // 代理规则
	AgreementCategoryConf        SettingsConfCategory = "agreement"      // 协议
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
)

func (t SettingsConfKey) Str() string {
	return string(t)
}
