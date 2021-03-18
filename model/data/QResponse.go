package data

import (
	"github.com/bububa/oppo-omni/model"
)

type QResponse struct {
	model.BaseResponse
	Data *QResult `json:"data,omitemtpy"`
}

type QResult struct {
	Expose   int64       `json:"expose,omitempty"`   // 曝光次数
	Click    int64       `json:"click,omitempty"`    // 点击次数
	Download int64       `json:"download,omitempty"` // 下载次数
	Cost     int64       `json:"cost,omitempty"`     // 消耗(单位:分)
	DataList []QListItem `json:"dataList,omitempty"` // 按时间粒度的列表
}
