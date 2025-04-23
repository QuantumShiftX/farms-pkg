package pmodel

import "github.com/QuantumShiftX/golib/stores/gormx"

const TableNameAppPackage = "app_packages"

type AppPackage struct {
	gormx.OperationBaseModel
	gormx.Model

	ID                   int64  `gorm:"primaryKey;column:id;autoIncrement;comment:自增主键" json:"id"`
	Icon                 string `gorm:"column:icon;type:varchar(255);comment:图标" json:"icon"`
	SystemType           int64  `gorm:"column:system_type;not null;default:1;comment:系统类型 默认为1 (1: 安卓, 2: iOS)" json:"system_type"`
	AppName              string `gorm:"column:app_name;type:varchar(255);comment:应用名称" json:"app_name"`
	AppPackageName       string `gorm:"column:app_package_name;type:varchar(255);comment:应用包名" json:"app_package_name"`
	QuickAppDownloadURL  string `gorm:"column:quick_app_download_url;type:varchar(255);comment:极速APP下载地址" json:"quick_app_download_url"`
	NativeAppDownloadURL string `gorm:"column:native_app_download_url;type:varchar(255);comment:原生APP下载地址" json:"native_app_download_url"`
	Version              string `gorm:"column:version;type:varchar(255);comment:版本号" json:"version"`
	InternalVersion      int64  `gorm:"column:internal_version;not null;comment:内部版本号" json:"internal_version"`
	UpdateLog            string `gorm:"column:update_log;type:varchar(100);comment:更新日志" json:"update_log"`
	Remark               string `gorm:"column:remark;type:varchar(200);comment:备注" json:"remark"`
	QuickAppPackageSize  int64  `gorm:"column:quick_app_package_size;default:0;comment:极速App包大小" json:"quick_app_package_size"`
	NativeAppPackageSize int64  `gorm:"column:native_app_package_size;default:0;comment:原生App包大小" json:"native_app_package_size"`
	IsForceUpdate        int64  `gorm:"column:is_force_update;not null;default:2;comment:是否强制更新 (1: 是, 2: 否)" json:"is_force_update"`
	Status               int64  `gorm:"column:status;not null;default:1;comment:状态 (1: 启动, 2: 禁用) 默认为1" json:"status"`
}

// TableName 设置表名
func (*AppPackage) TableName() string {
	return TableNameAppPackage
}

type AppPackageConfig struct {
	QuickAppStatus  int64 `json:"quick_app_status"`  // 极速App开关
	NativeAppStatus int64 `json:"native_app_status"` // 原生App开关
}

func (*AppPackageConfig) Default() *AppPackageConfig {
	return &AppPackageConfig{
		QuickAppStatus:  1,
		NativeAppStatus: 1,
	}
}
