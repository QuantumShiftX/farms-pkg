package pmodel

import "github.com/QuantumShiftX/golib/stores/gormx"

const TableNameAgentRelationship = "agent_relationships"

// AgentRelationship 代理层级关系表
type AgentRelationship struct {
	gormx.Model
	//
	ID         int64  `gorm:"column:id;primaryKey;autoIncrement" json:"id"`         // 主键ID
	AgentID    int64  `gorm:"column:agent_id;not null;index" json:"agent_id"`       // 代理ID
	ParentID   int64  `gorm:"column:parent_id;not null;index" json:"parent_id"`     // 上级代理ID
	AgentLevel int64  `gorm:"column:agent_level;not null;index" json:"agent_level"` // 层级关系：1表示直接上级，2表示上上级，依此类推
	Agent      *Agent `gorm:"foreignKey:AgentID;references:AgentID" json:"agent"`   // 关联代理
	Parent     *Agent `gorm:"foreignKey:ParentID;references:AgentID" json:"parent"` // 关联上级代理
}

// TableName 指定表名
func (*AgentRelationship) TableName() string {
	return TableNameAgentRelationship
}
