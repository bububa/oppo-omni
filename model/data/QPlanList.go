package data

import (
	"github.com/bububa/oppo-omni/enum"
	"github.com/bububa/oppo-omni/model"
)

type QPlanListRequest struct {
	model.BaseRequest
	ChnID         *enum.DataChnID         `json:"chnId,omitempty"`         //  推广渠道
	ExtensionType *enum.DataExtensionType `json:"extensionType,omitempty"` // 推广目的
	BeginTime     int64                   `json:"beginTime,omitempty"`     // 开始时间
	EndTime       int64                   `json:"endTime,omitempty"`       // 结束时间
	TimeLevel     enum.DataTimeLevel      `json:"timeLevel,omitempty"`     // 时间粒度
	PlanIDs       []uint64                `json:"planIds,omitempty"`
	Page          int                     `json:"page,omitempty"`
	PageSize      int                     `json:"pageCount,omitempty"`
}
