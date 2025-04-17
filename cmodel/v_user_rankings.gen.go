package cmodel

const TableNameVUserRanking = "v_user_rankings"

// VUserRanking 用户全局排名模型
type VUserRanking struct {
	UserID    int64  `gorm:"column:user_id" json:"user_id"`       // 用户ID
	Username  string `gorm:"column:username" json:"username"`     // 用户名
	Coin      int64  `gorm:"column:coin" json:"coin"`             // 金币数量
	UpdatedAt int64  `gorm:"column:updated_at" json:"updated_at"` // 更新时间
}

// TableName 设置表名
func (*VUserRanking) TableName() string {
	return "v_user_rankings"
}
