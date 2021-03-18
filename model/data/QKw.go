package data

import (
	"github.com/bububa/oppo-omni/enum"
	"github.com/bububa/oppo-omni/model"
)

type QKwRequest struct {
	model.BaseRequest
	BeginTime int64              `json:"beginTime,omitempty"` // 开始时间
	EndTime   int64              `json:"endTime,omitempty"`   // 结束时间
	TimeLevel enum.DataTimeLevel `json:"timeLevel,omitempty"` // 时间粒度
	AdIDs     []uint64           `json:"adIds,omitempty"`
}
