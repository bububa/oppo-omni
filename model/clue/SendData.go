package clue

import (
	"github.com/bububa/oppo-omni/enum"
	"github.com/bububa/oppo-omni/model"
)

type SendDataRequest struct {
	model.BaseRequest
	PageID        uint64                 `json:"pageId,omitempty"`        // 落地页 Id:投放广告到投放端时 生成,投放广告到播放端时添加
	Ip            string                 `json:"ip,omitempty"`            // 用户 IP:广告主收集
	Tid           string                 `json:"tid,omitempty"`           // traceId:播放时追加在 URL 上
	Lbid          string                 `json:"lbid,omitempty"`          // 流量号:播放时追加在 URL 上
	Items         []DataItem             `json:"items,omitempty"`         // 提交信息:页面内容
	TransformType enum.ClueTransformType `json:"transformType,omitempty"` // 转换类型:广告主提供值
	PageType      enum.CluePageType      `json:"pageType,omitempty"`      // 落地页类型:默认
}
