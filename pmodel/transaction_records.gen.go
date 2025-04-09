package pmodel

const TableNameTransactionRecord = "transaction_records"

// TransactionRecord 账变记录表
type TransactionRecord struct {
	ID              int64  `gorm:"primaryKey;autoIncrement;column:id"`                           // 主键，自动增长
	UserID          int64  `gorm:"column:user_id;not null"`                                      // 用户ID
	Username        string `gorm:"column:username;not null"`                                     // 用户账号
	OrderNo         string `gorm:"column:order_no;not null"`                                     // 流水编号
	CurrencyCode    string `gorm:"column:currency_code;not null"`                                // 币种编码
	WalletType      int16  `gorm:"column:wallet_type;not null"`                                  // 账变钱包类型（1:用户钱包 2:代理钱包 3:奖励钱包 4:USTD钱包）
	Category        int16  `gorm:"column:category;not null"`                                     // 账变大类交易类型
	SubCategory     int16  `gorm:"column:sub_category"`                                          // 账变小类
	BalancePrevious int64  `gorm:"column:balance_previous;not null"`                             // 变动前余额, 单位:微
	Amount          int64  `gorm:"column:amount;not null"`                                       // 变动金额, 单位:微
	BalanceAfter    int64  `gorm:"column:balance_after;not null"`                                // 变动后余额, 单位:微
	Remark          string `gorm:"column:remark"`                                                // 备注信息
	TransactionTime int64  `gorm:"column:transaction_time;not null"`                             // 交易时间
	RelatedID       string `gorm:"column:related_id;not null"`                                   // 关联ID，如游戏ID，商品ID，活动ID
	MsgID           string `gorm:"column:msg_id"`                                                // 消息队列消息ID
	CreatedAt       int64  `gorm:"column:created_at;not null;default:EXTRACT(epoch FROM now())"` // 创建时间
	UpdatedAt       int64  `gorm:"column:updated_at;not null;default:EXTRACT(epoch FROM now())"` // 更新时间
	DeletedAt       int64  `gorm:"column:deleted_at;not null;default:0"`                         // 删除时间
}

// TableName 设置数据库表名
func (*TransactionRecord) TableName() string {
	return TableNameTransactionRecord
}
