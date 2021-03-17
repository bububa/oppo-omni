package clue

import "github.com/bububa/oppo-omni/enum"

// 来源于用户填写表单
type DataItem struct {
	Column  int                   `json:"column,omitempty"`  // 列名
	Type    enum.ClueDataItemType `json:"type,omitempty"`    // 类型
	IfNeed  *bool                 `json:"ifNeed,omitempty"`  // 是否必须
	Desc    string                `json:"desc,omitempty"`    // 描述
	Value   int64                 `json:"value,omitempty"`   // 值
	Options []string              `json:"options,omitempty"` // 选项
}
