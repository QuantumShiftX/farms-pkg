package constants

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
