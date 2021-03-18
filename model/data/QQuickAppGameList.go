package data

import (
	"github.com/bububa/oppo-omni/enum"
	"github.com/bububa/oppo-omni/model"
)

type QQuickAppGameListRequest struct {
	model.BaseRequest
	ChnID         *enum.DataChnID         `json:"chnId,omitempty"`         //  推广渠道
	ExtensionType *enum.DataExtensionType `json:"extensionType,omitempty"` // 推广目的
	ShowType      *enum.DataShowType      `json:"showType,omitempty"`      // 推广样式
	AdSpec        int                     `json:"adSpec,omitempty"`        // 推广规格
	BeginTime     int64                   `json:"beginTime,omitempty"`     // 开始时间
	EndTime       int64                   `json:"endTime,omitempty"`       // 结束时间
	TimeLevel     enum.DataTimeLevel      `json:"timeLevel,omitempty"`     // 时间粒度
	AdIDs         []uint64                `json:"adIds,omitempty"`
	Page          int                     `json:"page,omitempty"`
	PageSize      int                     `json:"pageCount,omitempty"`
}
