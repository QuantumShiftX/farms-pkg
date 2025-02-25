package cmodel

import (
	"github.com/QuantumShiftX/golib/stores/gormx"
	"github.com/google/uuid"
)

const TableNameUserRegisterLog = "user_register_logs"

// UserRegisterLog 用户注册日志
type UserRegisterLog struct {
	ID            uuid.UUID `gorm:"type:uuid;default:generateUUIDv4();primary_key" json:"id"`
	UserID        int64     `gorm:"column:user_id;comment:用户id" json:"user_id"`
	IP            string    `gorm:"column:ip;comment:注册ip" json:"ip"`
	Domain        string    `gorm:"column:domain;default:'';comment:注册域名" json:"domain"`
	Link          string    `gorm:"column:link;default:'';comment:注册地址" json:"link"`
	DeviceNo      string    `gorm:"column:device_no;default:'';comment:设备号" json:"device_no"`
	DeviceModel   string    `gorm:"column:device_model;default:'';comment:设备型号" json:"device_model"`
	BrowserFinger string    `gorm:"column:browser_finger;default:'';comment:浏览器指纹" json:"browser_finger"`
	Area          string    `gorm:"column:area;default:'';comment:注册地区" json:"area"`
	PkgID         int64     `gorm:"column:pkg_id;comment:注册包id" json:"pkg_id"`
	PkgName       string    `gorm:"column:pkg_name;comment:注册来源" json:"pkg_name"`
	Type          int64     `gorm:"column:type;default:3;comment:注册方式：1-账号；2-手机号；3-邮箱；4-Facebook；5-Google；6-line" json:"type"`
	Source        int64     `gorm:"column:source;default:1;comment:注册来源：1-无渠道;2-未知渠道;3-推广注册;4-后台添加;5-API注册;6-TikTok" json:"source"`
	//
	gormx.Model
}

// TableName 指定表名
func (*UserRegisterLog) TableName() string {
	return TableNameUserRegisterLog
}
