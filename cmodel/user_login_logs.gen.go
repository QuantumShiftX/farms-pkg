package cmodel

import (
	"github.com/QuantumShiftX/golib/stores/gormx"
)

const TableNameUserLoginLogs = "user_login_logs"

// UserLoginLog 用户登录日志
type UserLoginLog struct {
	ID            string `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`   // UUID，使用 PostgreSQL 的 UUID 类型
	UserID        int64  `gorm:"default:0;comment:用户id" json:"user_id"`          // 用户 ID
	IP            string `gorm:"comment:登录ip" json:"ip"`                         // 登录 IP
	Domain        string `gorm:"type:VARCHAR(255);comment:登录域名" json:"domain"` // 登录域名
	Area          string `gorm:"comment:登录区域" json:"area"`                     // 登录区域
	DeviceNo      string `gorm:"comment:设备号" json:"device_no"`                  // 设备号
	DeviceModel   string `gorm:"comment:设备型号" json:"device_model"`             // 设备型号
	BrowserFinger string `gorm:"comment:浏览器指纹" json:"browser_finger"`         // 浏览器指纹
	LoginType     int64  `gorm:"default:1;comment:登录方式；1-账号登录；2-邮箱登录；3-手机登录；4-Facebook；5-Google；6-Line" json:"login_type"`
	OS            string `gorm:"type:VARCHAR(255);default:'';comment:操作系统：Windows、iPhone OS、Android、Intel Mac OS等" json:"os"`
	AppVersion    string `gorm:"default:'';comment:应用版本号" json:"app_version"`                                                     // 应用版本号
	System        string `gorm:"type:VARCHAR(255);default:'';comment:操作系统：Windows、iPhone OS、Android、Intel Mac OS等" json:"system"` // 系统
	//
	gormx.Model
}

// TableName 指定表名
func (*UserLoginLog) TableName() string {
	return TableNameUserLoginLogs
}
