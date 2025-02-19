// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package pmodel

import "github.com/QuantumShiftX/golib/stores/gormx"

const TableNameAdminLoginLog = "admin_login_logs"

// AdminLoginLog mapped from table <admin_login_logs>
type AdminLoginLog struct {
	ID          int64 `gorm:"column:id;primaryKey;autoIncrement:true;comment:主键ID" json:"id"`      // 主键ID
	LoggedAt    int64  `gorm:"column:logged_at;not null;comment:登录时间" json:"logged_at"`             // 登录时间
	IP          string `gorm:"column:ip;not null;comment:IP地址" json:"ip"`                           // IP地址
	Os          string `gorm:"column:os;not null;comment:操作系统" json:"os"`                           // 操作系统
	DeviceModel string `gorm:"column:device_model;not null;comment:设备型号" json:"device_model"`       // 设备型号
	Result      int64  `gorm:"column:result;not null;default:1;comment:结果：1-成功，2-失败" json:"result"` // 结果：1-成功，2-失败
	LoggedOutAt int64  `gorm:"column:logged_out_at;not null;comment:退出登录时间" json:"logged_out_at"`   // 退出登录时间
	Token       string `gorm:"column:token;not null;comment:toke信息" json:"token"`                   // toke信息
	//
	gormx.OperationBaseModel
	gormx.Model
}

// TableName AdminLoginLog's table name
func (*AdminLoginLog) TableName() string {
	return TableNameAdminLoginLog
}
