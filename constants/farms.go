package constants

import (
	"github.com/QuantumShiftX/farms-proto/proto-gen-go/user/v1"
)

// StoreProductType 存储商品类型
type StoreProductType int64

const (
	_ StoreProductType = iota //
	// StoreProductTypeSeed 1-种子
	StoreProductTypeSeed
	// StoreProductTypeFertilizer 2-肥料
	StoreProductTypeFertilizer
)

func (t StoreProductType) Int64() int64 {
	return int64(t)
}

func (t StoreProductType) Int8() int8 {
	return int8(t)
}

// 生长阶段百分比阈值
const (
	SeedThreshold      = 1000  // 10.00%
	SproutThreshold    = 2500  // 25.00%
	GrowthThreshold    = 4500  // 45.00%
	FloweringThreshold = 6500  // 65.00%
	MatureThreshold    = 8500  // 85.00%
	MaxPercentage      = 10000 // 100.00%
)

// 默认生长时间
const (
	DefaultGrowthTime = 864000 // 默认10天（秒）
)

// OperationTypeMap 将protobuf枚举映射到数据库操作的int64
var OperationTypeMap = map[v1.FarmOperationType]int64{
	v1.FarmOperationType_PLANTING:    1,
	v1.FarmOperationType_WATERING:    2,
	v1.FarmOperationType_FERTILIZING: 3,
	v1.FarmOperationType_HARVESTING:  4,
}

// FarmOperationType 操作类型
type FarmOperationType int64

const (
	_ FarmOperationType = iota //
	// FarmOperationTypeWater  浇水
	FarmOperationTypeWater
	// FarmOperationTypeFertilizer 施肥
	FarmOperationTypeFertilizer
)

func (t FarmOperationType) Int64() int64 {
	return int64(t)
}

func (t FarmOperationType) Int8() int8 {
	return int8(t)
}
