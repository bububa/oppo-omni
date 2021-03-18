package data

import (
	"github.com/bububa/oppo-omni/enum"
	"github.com/bububa/oppo-omni/model"
)

type QGameRequest struct {
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

type QGameResponse struct {
	model.BaseResponse
	Data *QGameResult `json:"data,omitemtpy"`
}

type QGameResult struct {
	AdNewSsoid    int64       `json:"adNewSsoid,omitempty"`    // 注册人数
	AdPrice       int64       `json:"adPrice,omitempty"`       // 广告实扣金额
	AdNewPaySsoid int64       `json:"adNewPaySsoid,omitempty"` // 新增付费人数
	Download      int64       `json:"download,omitempty"`      // 下载次数
	DataList      []QListItem `json:"dataList,omitempty"`      // 按时间粒度的列表
}
