// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package pmodel

import "github.com/QuantumShiftX/golib/stores/gormx"

const TableNameUserStorage = "user_storages"

// UserStorage mapped from table <user_storages>
type UserStorage struct {
	ID          int64 `gorm:"column:id;primaryKey;autoIncrement:true;comment:自增主键ID" json:"id"`                                                                                                                             // 自增主键ID
	StorageID   int64 `gorm:"column:storage_id;not null;comment:用户的仓库物品ID" json:"storage_id"`                                                                                                                               // 用户的仓库物品ID
	UserID      int64 `gorm:"column:user_id;not null;index:idx_user_storages_sku,priority:2;index:idx_user_storages_user_id,priority:1;index:idx_user_storages_product_type,priority:2;comment:用户ID，关联到用户表" json:"user_id"` // 用户ID，关联到用户表
	ProductType int16 `gorm:"column:product_type;not null;index:idx_user_storages_product_type,priority:1;comment:商品类型: 1-种子 2-肥料 3-其他道具等" json:"product_type"`                                                             // 商品类型: 1-种子 2-肥料 3-其他道具等
	ProductID   int64 `gorm:"column:product_id;not null;comment:商品ID，关联到对应产品配置表" json:"product_id"`                                                                                                                         // 商品ID，关联到对应产品配置表

	ProductName                       string `gorm:"column:product_name;not null;comment:商品名称快照，保存购买时的名称" json:"product_name"`                                                      // 商品名称快照，保存购买时的名称
	ProductImageURL                   string `gorm:"column:product_image_url;not null;comment:商品图片地址快照，保存购买时的图片URL" json:"product_image_url"`                                       // 商品图片地址快照，保存购买时的图片URL
	ProductQuantity                   int32  `gorm:"column:product_quantity;not null;comment:商品数量，表示用户拥有的该商品数量" json:"product_quantity"`                                            // 商品数量，表示用户拥有的该商品数量
	SkuID                             int64  `gorm:"column:sku_id;not null;index:idx_user_storages_sku,priority:1;comment:商品SKU ID，关联到商品SKU配置表" json:"sku_id"`                      // 商品SKU ID，关联到商品SKU配置表
	SkuName                           string `gorm:"column:sku_name;not null;comment:SKU名称快照，保存购买时的SKU名称" json:"sku_name"`                                                          // SKU名称快照，保存购买时的SKU名称
	Price                             int64  `gorm:"column:price;comment:商品价格快照，保存购买时的价格（单位：ustd）" json:"price"`                                                                    // 商品价格快照，保存购买时的价格（单位：ustd）
	GrowthTime                        int64  `gorm:"column:growth_time;comment:生长时间快照（单位：分钟），仅种子类型有效" json:"growth_time"`                                                           // 生长时间快照（单位：分钟），仅种子类型有效
	CoinReward                        int64  `gorm:"column:coin_reward;comment:成熟后的收益快照（单位：代币），仅种子类型有效" json:"coin_reward"`                                                         // 成熟后的收益快照（单位：代币），仅种子类型有效
	WateringFrequency                 int64  `gorm:"column:watering_frequency;default:1;comment:浇水频率（次/天）" json:"watering_frequency"`                                               // 浇水频率（次/天）
	FertilizingFrequency              int64  `gorm:"column:fertilizing_frequency;default:1;comment:施肥频率（次/天）" json:"fertilizing_frequency"`                                         // 施肥频率（次/天）
	SelfTimeReducedPerWater           int64  `gorm:"column:self_time_reduced_per_water;comment:自己每次浇水减少的时间快照（单位：分钟），仅种子类型有效" json:"self_time_reduced_per_water"`                    // 自己每次浇水减少的时间快照（单位：分钟），仅种子类型有效
	SelfWaterReductionPercentage      int64  `gorm:"column:self_water_reduction_percentage;comment:自己浇水减少时间的百分比快照（0-100），仅种子类型有效" json:"self_water_reduction_percentage"`           // 自己浇水减少时间的百分比快照（0-100），仅种子类型有效
	SelfTimeReducedPerFertilizer      int64  `gorm:"column:self_time_reduced_per_fertilizer;comment:自己每次施肥减少的时间快照（单位：分钟），仅种子类型有效" json:"self_time_reduced_per_fertilizer"`          // 自己每次施肥减少的时间快照（单位：分钟），仅种子类型有效
	SelfFertilizerReductionPercentage int64  `gorm:"column:self_fertilizer_reduction_percentage;comment:自己施肥减少时间的百分比快照（0-100），仅种子类型有效" json:"self_fertilizer_reduction_percentage"` // 自己施肥减少时间的百分比快照（0-100），仅种子类型有效
	FriendWaterTimeReduction          int64  `gorm:"column:friend_water_time_reduction;comment:好友每次浇水减少的时间快照（单位：分钟），仅种子类型有效" json:"friend_water_time_reduction"`                    // 好友每次浇水减少的时间快照（单位：分钟），仅种子类型有效
	FriendWaterPercentage             int64  `gorm:"column:friend_water_percentage;comment:好友浇水减少时间的百分比快照（0-100），仅种子类型有效" json:"friend_water_percentage"`                           // 好友浇水减少时间的百分比快照（0-100），仅种子类型有效
	FriendFertilizerTimeReduction     int64  `gorm:"column:friend_fertilizer_time_reduction;comment:好友每次施肥减少的时间快照（单位：分钟），仅种子类型有效" json:"friend_fertilizer_time_reduction"`          // 好友每次施肥减少的时间快照（单位：分钟），仅种子类型有效
	FriendFertilizerPercentage        int64  `gorm:"column:friend_fertilizer_percentage;comment:好友施肥减少时间的百分比快照（0-100），仅种子类型有效" json:"friend_fertilizer_percentage"`                 // 好友施肥减少时间的百分比快照（0-100），仅种子类型有效
	//
	gormx.Model
}

// TableName UserStorage's table name
func (*UserStorage) TableName() string {
	return TableNameUserStorage
}
