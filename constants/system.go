package constants

type SettingsConfCategory string

const (
	BasicSettingsCategoryConf    SettingsConfCategory = "basic_settings"
	AgentRebateRulesCategoryConf SettingsConfCategory = "rebate_rules"
)

func (t SettingsConfCategory) Str() string {
	return string(t)
}

type SettingsConfKey string

const (
	BasicSettingsKeyConf    SettingsConfKey = "basic_settings"
	AgentRebateRulesKeyConf SettingsConfKey = "rebate_rules"
)

func (t SettingsConfKey) Str() string {
	return string(t)
}
