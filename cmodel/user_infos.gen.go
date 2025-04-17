package cmodel

const TableNameUserInfo = "user_infos"

// UserInfo 用户信息表
type UserInfo struct {
	ID        string `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	UserID    int64  `gorm:"column:user_id;uniqueIndex" json:"user_id"`
	Username  string `gorm:"column:username" json:"username"`
	FarmID    int64  `gorm:"column:farm_id" json:"farm_id"`
	LandID    int64  `gorm:"column:land_id" json:"land_id"`
	Coin      int64  `gorm:"column:coin;default:0" json:"coin"`
	Status    int64  `gorm:"column:status;default:1" json:"status"`
	UpdatedAt int64  `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
	Sign      int8   `gorm:"column:sign;comment:'标记字段：1 表示新增行，-1 表示删除行'"` // 标记字段
	Version   int64  `gorm:"column:version;comment:'版本号(Unix时间戳，微秒)'"`    // 版本号
}

// TableName 指定表名
func (*UserInfo) TableName() string {
	return TableNameUserInfo
}
