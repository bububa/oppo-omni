package data

import "github.com/bububa/oppo-omni/model"

type QTodayTotalRequest struct {
	model.BaseRequest
	AdID uint64 `json:"adId,omitempty"` // 广告ID
}
type QTodayTotalResponse struct {
	model.BaseResponse
	Data *QTodayTotalResult `json:"data,omitempty"`
}

type QTodayTotalResult struct {
	StatTime int64               `json:"statTime,omitempty"` // 流水时间,unixtime
	Expose   int64               `json:"expose,omitempty"`   // 曝光次数
	Click    int64               `json:"click,omitempty"`    // 点击次数
	Download int64               `json:"download,omitempty"` // 下载次数
	Cost     int64               `json:"cost,omitempty"`     // 消耗(单位:分)
	DataList []QTodayTotalResult `json:"dataList,omitempty"` // 按时间粒度的列表
}
