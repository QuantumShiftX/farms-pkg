package pmodel

import "github.com/QuantumShiftX/golib/stores/gormx"

// 用户提现账户配置表
const TableNameUserWithdrawAccount = "user_withdraw_accounts"

// UserWithdrawAccount mapped from table <user_withdraw_accounts>
type UserWithdrawAccount struct {
	ID           int64  `gorm:"column:id;primaryKey;autoIncrement:true;comment:主键" json:"id"`                                                                                      // 主键
	UserID       int64  `gorm:"column:user_id;not null;default:0;index:idx_user_withdraw_accounts_user_id_status,priority:1;comment:用户id" json:"user_id"`                          // 用户id
	CurrencyCode string `gorm:"column:currency_code;type:varchar(64);not null;default:'';comment:币种编码" json:"currency_code"`                                                       // 币种编码
	BankID       int64  `gorm:"column:bank_id;type:smallint;not null;default:0;comment:银行id(bank_configs表id)" json:"bank_id"`                                                      // 银行id(bank_configs表id)
	Account      string `gorm:"column:account;type:varchar(1024);not null;default:'';comment:提现账号/地址" json:"account"`                                                              // 提现账号/地址
	RealName     string `gorm:"column:real_name;type:varchar(128);not null;default:'';comment:真实姓名" json:"real_name"`                                                              // 真实姓名
	Identity     string `gorm:"column:identity;type:varchar(128);not null;default:'';comment:身份证号" json:"identity"`                                                                // 身份证号
	AccountType  int8   `gorm:"column:account_type;type:smallint;not null;default:0;comment:账号类型 1:手机号 2:钱包id" json:"account_type"`                                                //账号类型 1:手机号 2:银行卡 3:PIX 4:电子钱包
	CountryCode  string `gorm:"column:country_code;type:smallint;not null;default:'';comment:国家/地区id" json:"country_code"`                                                         // 国家/地区id
	IFSC         string `gorm:"column:ifsc;type:varchar(1024);not null;default:'';comment:开户行/IFSC" json:"ifsc"`                                                                   // 开户行/IFSC
	Branch       string `gorm:"column:branch;type:varchar(1024);not null;default:'';comment:支行名称" json:"branch"`                                                                   // 支行名称
	Remark       string `gorm:"column:remark;type:varchar(1024);not null;default:'';comment:后台备注" json:"remark"`                                                                   // 后台备注
	Status       int8   `gorm:"column:status;type:smallint;not null;default:1;index:idx_user_withdraw_accounts_user_id_status,priority:2;comment:启用/停用状态 1:启用 2:停用" json:"status"` // 启用/停用状态 1:启用 2:停用
	IsDefault    int8   `gorm:"column:is_default;type:smallint;not null;default:2;,priority:2;comment:是否默认账户 1:是 2:否" json:"is_default"`                                           // 是否默认账户 1:是 2:否
	VerifyStatus int8   `gorm:"column:verify_status;type:smallint;not null;default:3;,priority:2;comment:验证状态 1:已验证 2:验证失败 3:未验证" json:"verify_status"`                            // 验证状态 1:已验证 2:验证失败 3:未验证
	gormx.OperationBaseModel
	gormx.Model
}

// TableName UserWithdrawAccount's table name
func (*UserWithdrawAccount) TableName() string {
	return TableNameUserWithdrawAccount
}
