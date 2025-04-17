package cmodel

const TableNameVUserFriendRanking = "v_user_friend_rankings"

// VUserFriendRanking 用户好友排名模型
type VUserFriendRanking struct {
	UserID          int64  `gorm:"column:user_id" json:"user_id"`                     // 用户ID
	FriendID        int64  `gorm:"column:friend_id" json:"friend_id"`                 // 好友ID
	FriendUsername  string `gorm:"column:friend_username" json:"friend_username"`     // 好友用户名
	FriendCoin      int64  `gorm:"column:friend_coin" json:"friend_coin"`             // 好友金币数量
	FriendUpdatedAt int64  `gorm:"column:friend_updated_at" json:"friend_updated_at"` // 好友信息更新时间
}

// TableName 设置表名
func (*VUserFriendRanking) TableName() string {
	return TableNameVUserFriendRanking
}
