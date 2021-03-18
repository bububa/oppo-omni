package data

import (
	"github.com/bububa/oppo-omni/enum"
	"github.com/bububa/oppo-omni/model"
)

type QTodayTopRequest struct {
	model.BaseRequest
	Demision *enum.DataDemision `json:"demision,omitempty"`
}

type QTodayTopResponse struct {
	model.BaseResponse
	Data *QTodayTopResult `json:"data,omitempty"`
}

type QTodayTopResult struct {
	AdID      uint64 `json:"adId,omitempty"`      // 广告ID
	AdName    string `json:"adName,omitempty"`    // 广告名称
	OrderData int64  `json:"orderData,omitempty"` // 排序数据值(如果是下载量排序，就是下载次数，如果是消耗，就是消耗金额)
}
