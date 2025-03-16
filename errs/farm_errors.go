package errs

import "github.com/QuantumShiftX/golib/xerr"

const (
	FarmErrCode xerr.ErrCode = 20000

	FarmNotFoundError   = iota + FarmErrCode // 农场不存在
	FarmCreateError                          // 创建农场失败
	FarmUpdateError                          // 更新农场信息失败
	FarmDeleteError                          // 删除农场失败
	FarmPermissionError                      // 无权操作该农场
	FarmCropError                            // 作物操作错误
	FarmIrrigationError                      // 灌溉系统错误
	FarmInventoryError                       // 库存管理错误
)

var (
	ErrFarmNotFound   = xerr.New(FarmNotFoundError, "农场不存在")
	ErrFarmCreate     = xerr.New(FarmCreateError, "创建农场失败")
	ErrFarmUpdate     = xerr.New(FarmUpdateError, "更新农场信息失败")
	ErrFarmDelete     = xerr.New(FarmDeleteError, "删除农场失败")
	ErrFarmPermission = xerr.New(FarmPermissionError, "无权操作该农场")
	ErrFarmCrop       = xerr.New(FarmCropError, "作物操作错误")
	ErrFarmIrrigation = xerr.New(FarmIrrigationError, "灌溉系统错误")
	ErrFarmInventory  = xerr.New(FarmInventoryError, "库存管理错误")
)
