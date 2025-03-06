package cmodel

const TableNameUserFriend = "user_friends"

// UserFriend 好友关系表
type UserFriend struct {
	ID        string `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	UserID    int64  `gorm:"column:user_id;index:idx_user_friend" json:"user_id"`
	FriendID  int64  `gorm:"column:friend_id;index:idx_user_friend" json:"friend_id"`
	CreatedAt int64  `gorm:"column:created_at;autoCreateTime" json:"created_at"`
}

// TableName 指定表名
func (*UserFriend) TableName() string {
	return TableNameUserFriend
}
