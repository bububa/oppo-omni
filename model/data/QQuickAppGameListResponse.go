package data

import (
	"github.com/bububa/oppo-omni/model"
)

type QQuickAppGameListResponse struct {
	model.BaseResponse
	Data *QQuickAppGameListResult `json:"data,omitemtpy"`
}

type QQuickAppGameListResult struct {
	ItemCount  int64                   `json:"itemCount,omitempty"`
	TotalCount int64                   `json:"totalCount,omitempty"`
	Items      []QQuickAppGameListItem `json:"items,omitempty"`
}

type QQuickAppGameListItem struct {
	StatTime        int64   `json:"statTime,omitempty"`        // 流水时间,unixtime
	PlanID          uint64  `json:"planId,omitempty"`          // 计划ID
	PlanName        string  `json:"planName,omitempty"`        // 计划名称
	GroupID         uint64  `json:"groupId,omitempty"`         // 广告组ID
	GroupName       string  `json:"groupName,omitempty"`       // 广告组名
	AdID            uint64  `json:"adId,omitempty"`            // 广告ID
	AdName          string  `json:"adName,omitempty"`          // 广告名称
	Keywords        string  `json:"keywords,omitempty"`        // 关键词
	AppID           int64   `json:"appId,omitempty"`           // 应用ID
	AppName         string  `json:"appName,omitempty"`         // 应用名称
	AccCost         int64   `json:"accCost,omitempty"`         // 消耗(单位:分)
	Ecpm            string  `json:"ecpm,omitempty"`            // ECPM
	AdPrice         int64   `json:"adPrice,omitempty"`         // 广告实扣金额
	ExposeNums      int64   `json:"exposeNums,omitempty"`      // 曝光次数
	ClickNums       int64   `json:"clickNums,omitempty"`       // 点击次数
	ClickRate       string  `json:"clickRate,omitempty"`       // 点击率
	ClickPrice      string  `json:"clickPrice,omitempty"`      // 点击均价
	ActiveNums      int64   `json:"activeNums,omitempty"`      // 激活人数
	ActiveRate      string  `json:"activeRate,omitempty"`      // 激活率
	ActivePrice     string  `json:"activePrice,omitempty"`     // 激活成本
	NewUserBuyLtv1  float64 `json:"newUserBuyLtv1,omitempty"`  // 广告变现LTV1
	NewUserBuyLtv3  float64 `json:"newUserBuyLtv3,omitempty"`  // 广告变现LTV3
	NewUserBuyLtv7  float64 `json:"newUserBuyLtv7,omitempty"`  // 广告变现LTV7
	NewUserBuyLtv30 float64 `json:"newUserBuyLtv30,omitempty"` // 广告变现LTV30
	NewUserBuyRoi1  string  `json:"newUserBuyRoi1,omitempty"`  // 广告变现ROI1
	NewUserBuyRoi3  string  `json:"newUserBuyRoi3,omitempty"`  // 广告变现ROI3
	NewUserBuyRoi7  string  `json:"newUserBuyRoi7,omitempty"`  // 广告变现ROI7
	NewUserBuyRoi30 string  `json:"newUserBuyRoi30,omitempty"` // 广告变现ROI30
	BuyUserIncome0  float64 `json:"buyUserIncome0,omitempty"`  // 首日广告收入
	NewPayUserNums  int64   `json:"newPayUserNums,omitempty"`  // 新增付费人数
	PayTransferRate string  `json:"payTransferRate,omitempty"` // 付费转化率
	PayCost         float64 `json:"payCost,omitempty"`         // 付费成本
}
