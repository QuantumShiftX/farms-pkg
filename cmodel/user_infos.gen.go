package cmodel

const TableNameUserInfo = "user_info"

// UserInfo 用户信息表
type UserInfo struct {
	ID         string `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	UserID     int64  `gorm:"column:user_id;uniqueIndex" json:"user_id"`
	Username   string `gorm:"column:username" json:"username"`
	AvatarURL  string `gorm:"column:avatar_url" json:"avatar_url"`
	VIPLevelID int64  `gorm:"column:vip_level_id;default:1" json:"vip_level_id"`
	FarmID     int64  `gorm:"column:farm_id" json:"farm_id"`
	LandID     int64  `gorm:"column:land_id" json:"land_id"`
	Coin       int64  `gorm:"column:coin;default:0" json:"coin"`
	UpdatedAt  int64  `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}

// TableName 指定表名
func (*UserInfo) TableName() string {
	return TableNameUserInfo
}
