package pmodel

import "github.com/QuantumShiftX/golib/stores/gormx"

const TableNameCustomerService = "customer_services"

// CustomerService 客服表的 GORM 模型
type CustomerService struct {
	ID          int64  `gorm:"primaryKey;autoIncrement;column:id" json:"id"`            // 主键，自动增长
	ServiceType int8   `gorm:"not null;column:service_type" json:"service_type"`        // 客服类型：1:Telegram客服, 2:Whatsapp客服, 3:Line客服, 4:三方客服
	ServiceLink string `gorm:"not null;unique;column:service_link" json:"service_link"` // 客服链接，唯一约束
	ServiceLogo string `gorm:"column:service_logo" json:"service_logo"`                 // 客服图标
	Status      int8   `gorm:"not null;default:1;column:status" json:"status"`          // 是否启用，1:启用，0:禁用
	Remarks     string `gorm:"column:remarks" json:"remarks"`                           // 备注，描述模板的用途或说明
	//
	gormx.OperationBaseModel
	gormx.Model
}

// TableName 定义数据库表名
func (*CustomerService) TableName() string {
	return TableNameCustomerService
}
