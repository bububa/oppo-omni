package data

import (
	"github.com/bububa/oppo-omni/enum"
	"github.com/bububa/oppo-omni/model"
)

type QQuickAppGameRequest struct {
	model.BaseRequest
	ChnID         *enum.DataChnID         `json:"chnId,omitempty"`         //  推广渠道
	ExtensionType *enum.DataExtensionType `json:"extensionType,omitempty"` // 推广目的
	ShowType      *enum.DataShowType      `json:"showType,omitempty"`      // 推广样式
	AdSpec        int                     `json:"adSpec,omitempty"`        // 推广规格
	BeginTime     int64                   `json:"beginTime,omitempty"`     // 开始时间
	EndTime       int64                   `json:"endTime,omitempty"`       // 结束时间
	TimeLevel     enum.DataTimeLevel      `json:"timeLevel,omitempty"`     // 时间粒度
	AdIDs         []uint64                `json:"adIds,omitempty"`
}

type QQuickAppGameResponse struct {
	model.BaseResponse
	Data *QQuickAppGameResult `json:"data,omitemtpy"`
}

type QQuickAppGameResult struct {
	AccCostTotal    int64                   `json:"accCostTotal,omitempty"`    // 实扣金额
	EcpmTotal       int64                   `json:"ecpmTotal,omitempty"`       // ECPM
	ActiveNumsTotal int64                   `json:"activeNumsTotal,omitempty"` // 新增用户数
	ActiveCostTotal int64                   `json:"activeCostTotal,omitempty"` // 新增用户成本
	DataList        []QQuickAppGameListItem `json:"dataList,omitempty"`        // 按时间粒度的列表
}
