package pmodel

import "github.com/QuantumShiftX/golib/stores/gormx"

const TableNameWalletManagement = "wallet_managements"

// WalletManagement 钱包管理表 - 用于管理平台的收款钱包和提款钱包
type WalletManagement struct {
	ID                   int64  `gorm:"column:id;primaryKey;autoIncrement" json:"id"`                                                          // 主键ID，自增
	WalletType           int8   `gorm:"column:wallet_type;not null" json:"wallet_type"`                                                        // 钱包类型：1-收款钱包 2-打款钱包
	Network              string `gorm:"column:network;size:20;not null" json:"network"`                                                        // 网络类型，如TRC20、ERC20等
	WalletAddress        string `gorm:"column:wallet_address;size:255;not null;uniqueIndex:addr_idx,where:deleted_at=0" json:"wallet_address"` // 钱包地址
	PrivateKey           string `gorm:"column:private_key;size:255" json:"private_key"`                                                        // 私钥，加密存储
	UsdtBalance          int64  `gorm:"column:usdt_balance;default:0" json:"usdt_balance"`                                                     // USDT余额
	TokenBalance         int64  `gorm:"column:token_balance;default:0" json:"token_balance"`                                                   // 平台币余额
	TokenName            string `gorm:"column:token_name;size:20" json:"token_name"`                                                           // 平台币名称，如TRX
	AutoCollect          int8   `gorm:"column:auto_collect;default:1" json:"auto_collect"`                                                     // 是否自动归集
	AutoCollectThreshold int64  `gorm:"column:auto_collect_threshold" json:"auto_collect_threshold"`                                           // 自动归集阈值(USDT)
	Status               int8   `gorm:"column:status;default:1;index" json:"status"`                                                           // 状态：1-正常，2-禁用
	Remarks              string `gorm:"column:remarks;type:text" json:"remarks"`                                                               // 备注
	//
	gormx.OperationBaseModel
	gormx.Model
}

// TableName 设置WalletManagement的表名
func (*WalletManagement) TableName() string {
	return TableNameWalletManagement
}
