package vars

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/logx"
	"regexp"
	"strings"
	"sync"
	"time"
)

// TemplateEngineVarType 定义变量替换类型
type TemplateEngineVarType string

func (t TemplateEngineVarType) ToStr() string {
	return string(t)
}

const (
	// UserID user id
	UserID TemplateEngineVarType = "user_id"
	// LevelName vip等级名称
	LevelName TemplateEngineVarType = "level_name"
	// FarmName 农场名称
	FarmName TemplateEngineVarType = "farm_name"
	// Username 用户名称
	Username TemplateEngineVarType = "username"
	// Nickname 用户昵称
	Nickname TemplateEngineVarType = "nickname"
	// 未来可以在此处添加更多变量类型
)

// TemplateEngineVal 模板变量值结构体
type TemplateEngineVal struct {
	UserID    int64  `json:"user_id"`
	LevelName string `json:"level_name"`
	FarmName  string `json:"farm_name"`
	Username  string `json:"username"`
	Nickname  string `json:"nickname"`
	// 支持自定义扩展字段
	ExtFields map[string]interface{} `json:"ext_fields,omitempty"`
}

// 使用懒加载模式初始化正则表达式
var (
	templateRegex     *regexp.Regexp
	templateRegexOnce sync.Once
)

// getTemplateRegex 获取懒加载的正则表达式
func getTemplateRegex() *regexp.Regexp {
	templateRegexOnce.Do(func() {
		templateRegex = regexp.MustCompile(`\${([^}]+)}`)
	})
	return templateRegex
}

// GetAllTemplateVars 将结构体转换为变量映射
func GetAllTemplateVars(val *TemplateEngineVal) map[TemplateEngineVarType]interface{} {
	if val == nil {
		return make(map[TemplateEngineVarType]interface{})
	}

	// 预计扩展字段大小，避免动态扩容
	extFieldsSize := 0
	if val.ExtFields != nil {
		extFieldsSize = len(val.ExtFields)
	}

	// 创建足够大的初始容量，减少扩容操作
	variables := make(map[TemplateEngineVarType]interface{}, 5+extFieldsSize)

	// 添加基本字段
	variables[UserID] = val.UserID
	variables[LevelName] = val.LevelName
	variables[FarmName] = val.FarmName
	variables[Username] = val.Username
	variables[Nickname] = val.Nickname

	// 添加扩展字段
	if val.ExtFields != nil {
		for key, value := range val.ExtFields {
			variables[TemplateEngineVarType(key)] = value
		}
	}

	return variables
}

// TemplateEngine 根据不同类型替换文本中的变量
func TemplateEngine(template string, val *TemplateEngineVal) string {
	// 如果模板为空或值为nil，直接返回
	if template == "" || val == nil {
		return template
	}

	// 快速检查是否包含变量标记，避免不必要的处理
	if !strings.Contains(template, "${") {
		return template
	}

	// 获取所有变量值
	variables := GetAllTemplateVars(val)
	if len(variables) == 0 {
		return template
	}

	// 使用builder提高字符串连接效率
	var result strings.Builder
	lastIndex := 0

	// 使用Find方法找到所有匹配项的索引
	matches := getTemplateRegex().FindAllStringSubmatchIndex(template, -1)
	if len(matches) == 0 {
		return template
	}

	// 预分配足够的空间
	result.Grow(len(template) + 100)

	for _, match := range matches {
		// 添加匹配前的内容
		result.WriteString(template[lastIndex:match[0]])

		// 获取变量名
		varName := template[match[2]:match[3]]

		// 查找变量值
		value, exists := variables[TemplateEngineVarType(varName)]
		if !exists {
			// 变量不存在，保留原始表达式
			result.WriteString(template[match[0]:match[1]])
		} else {
			// 格式化并添加值
			formattedValue := formatValue(value, template[match[0]:match[1]])
			result.WriteString(formattedValue)
		}

		// 更新最后处理的位置
		lastIndex = match[1]
	}

	// 添加剩余内容
	if lastIndex < len(template) {
		result.WriteString(template[lastIndex:])
	}

	return result.String()
}

// formatValue 根据值类型格式化为字符串
func formatValue(value interface{}, defaultValue string) string {
	if value == nil {
		return ""
	}

	switch v := value.(type) {
	case string:
		return v
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return fmt.Sprintf("%d", v)
	case float32, float64:
		return fmt.Sprintf("%.2f", v)
	case bool:
		return fmt.Sprintf("%t", v)
	case []string:
		return strings.Join(v, ", ")
	case []interface{}:
		// 预分配内存优化
		items := make([]string, 0, len(v))
		for _, item := range v {
			items = append(items, fmt.Sprintf("%v", item))
		}
		return strings.Join(items, ", ")
	case map[string]interface{}:
		jsonStr, err := jsonx.Marshal(v)
		if err != nil {
			logx.Errorf("转换JSON失败: %v", err)
			return defaultValue
		}
		return string(jsonStr)
	case time.Time:
		return v.Format(time.DateTime)
	default:
		return fmt.Sprintf("%v", v)
	}
}

// NewTemplateEngineVal 创建模板变量值对象的便捷方法
func NewTemplateEngineVal() *TemplateEngineVal {
	return &TemplateEngineVal{
		ExtFields: make(map[string]interface{}),
	}
}

// WithUserID 设置用户ID
func (t *TemplateEngineVal) WithUserID(userID int64) *TemplateEngineVal {
	t.UserID = userID
	return t
}

// WithUsername 设置用户名
func (t *TemplateEngineVal) WithUsername(username string) *TemplateEngineVal {
	t.Username = username
	return t
}

// WithNickname 设置昵称
func (t *TemplateEngineVal) WithNickname(nickname string) *TemplateEngineVal {
	t.Nickname = nickname
	return t
}

// WithLevelName 设置等级名称
func (t *TemplateEngineVal) WithLevelName(levelName string) *TemplateEngineVal {
	t.LevelName = levelName
	return t
}

// WithFarmName 设置农场名称
func (t *TemplateEngineVal) WithFarmName(farmName string) *TemplateEngineVal {
	t.FarmName = farmName
	return t
}

// WithExtField 添加扩展字段
func (t *TemplateEngineVal) WithExtField(key string, value interface{}) *TemplateEngineVal {
	if t.ExtFields == nil {
		t.ExtFields = make(map[string]interface{})
	}
	t.ExtFields[key] = value
	return t
}

// WithExtFields 批量添加扩展字段
func (t *TemplateEngineVal) WithExtFields(fields map[string]interface{}) *TemplateEngineVal {
	if fields == nil {
		return t
	}

	if t.ExtFields == nil {
		t.ExtFields = make(map[string]interface{}, len(fields))
	}

	for k, v := range fields {
		t.ExtFields[k] = v
	}
	return t
}

// Clone 克隆一个新的模板值对象
func (t *TemplateEngineVal) Clone() *TemplateEngineVal {
	if t == nil {
		return NewTemplateEngineVal()
	}

	clone := &TemplateEngineVal{
		UserID:    t.UserID,
		LevelName: t.LevelName,
		FarmName:  t.FarmName,
		Username:  t.Username,
		Nickname:  t.Nickname,
	}

	// 复制扩展字段
	if t.ExtFields != nil && len(t.ExtFields) > 0 {
		clone.ExtFields = make(map[string]interface{}, len(t.ExtFields))
		for k, v := range t.ExtFields {
			clone.ExtFields[k] = v
		}
	}

	return clone
}

//func main() {
//	template := "用户${username}(ID:${user_id})在${farm_name}种植了${plant_count}种植物，当前等级为${level_name}。收获时间：${harvest_time}"
//
//	// 创建扩展字段映射
//	extFields := map[string]interface{}{
//		"plant_count":  8,
//		"harvest_time": time.Now().Add(24 * time.Hour),
//	}
//
//	val := vars.NewTemplateEngineVal().
//		WithUserID(10086).
//		WithUsername("王五").
//		WithFarmName("快乐果园").
//		WithLevelName("黄金会员").
//		WithExtFields(extFields)
//
//	result := vars.TemplateEngine(template, val)
//
//	fmt.Println(result)
//	// 输出类似: 用户王五(ID:10086)在快乐果园种植了8种植物，当前等级为黄金会员。收获时间：2025-04-27 18:30:00
//}
//func ProcessBatchNotifications(ctx context.Context, userID int64, templates []string) {
//	// 创建基础值对象
//	baseVal := vars.NewTemplateEngineVal().
//		WithUserID(userID).
//		WithUsername("张三").
//		WithFarmName("快乐农场")
//
//	// 处理每个模板
//	for i, template := range templates {
//		// 克隆基础值对象
//		val := baseVal.Clone()
//
//		// 添加特定于此模板的值
//		val.WithExtField("template_index", i+1).
//			WithExtField("current_time", time.Now())
//
//		// 应用模板
//		result := vars.TemplateEngine(template, val)
//
//		// 进一步处理结果...
//		fmt.Println(result)
//	}
//}
