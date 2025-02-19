package constants

// ResultType 通用result
type ResultType int64

const (
	_ ResultType = iota //
	// ResultSuccess 成功
	ResultSuccess
	// ResultFail 失败
	ResultFail
)

func (t ResultType) Int64() int64 {
	return int64(t)
}
