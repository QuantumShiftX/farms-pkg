package cmodel

const TableNameUserRanking = "mv_user_rankings"

// MvUserRanking 用户全局排名视图模型
type MvUserRanking struct {
	ID         string `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	UserID     int64  `gorm:"column:user_id;primaryKey" json:"user_id"`
	Username   string `gorm:"column:username" json:"username"`
	AvatarURL  string `gorm:"column:avatar_url" json:"avatar_url"`
	VIPLevelID int64  `gorm:"column:vip_level_id" json:"vip_level_id"`
	Coin       int64  `gorm:"column:coin" json:"coin"`
	UpdatedAt  int64  `gorm:"column:updated_at" json:"updated_at"`
}

// TableName 指定表名
func (*MvUserRanking) TableName() string {
	return TableNameUserRanking
}
