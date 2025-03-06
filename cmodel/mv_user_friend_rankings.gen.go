package cmodel

const TableNameUserFriendRanking = "mv_user_friend_rankings"

// MvUserFriendRanking 用户好友排名视图模型
type MvUserFriendRanking struct {
	ID         string `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	UserID     int64  `gorm:"column:user_id;primaryKey;index:idx_user_friend" json:"user_id"`
	FriendID   int64  `gorm:"column:friend_id;primaryKey;index:idx_user_friend" json:"friend_id"`
	Username   string `gorm:"column:username" json:"username"`
	AvatarURL  string `gorm:"column:avatar_url" json:"avatar_url"`
	VIPLevelID int64  `gorm:"column:vip_level_id" json:"vip_level_id"`
	Coin       int64  `gorm:"column:coin" json:"coin"`
	UpdatedAt  int64  `gorm:"column:updated_at" json:"updated_at"`
}

// TableName 指定表名
func (*MvUserFriendRanking) TableName() string {
	return TableNameUserFriendRanking
}
