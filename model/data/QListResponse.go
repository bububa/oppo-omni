package data

import (
	"github.com/bububa/oppo-omni/model"
)

type QListResponse struct {
	model.BaseResponse
	Data *QListResult `json:"data,omitemtpy"`
}

type QListResult struct {
	ItemCount  int64       `json:"itemCount,omitempty"`
	TotalCount int64       `json:"totalCount,omitempty"`
	Items      []QListItem `json:"items,omitempty"`
}

type QListItem struct {
	StatTime              int64   `json:"statTime,omitempty"`              // 流水时间,unixtime
	PlanID                uint64  `json:"planId,omitempty"`                // 计划ID
	PlanName              string  `json:"planName,omitempty"`              // 计划名称
	AdGroupID             uint64  `json:"adGroupId,omitempty"`             // 广告组ID
	AdGroupName           string  `json:"adGroupName,omitempty"`           // 广告组名
	AdID                  uint64  `json:"adId,omitempty"`                  // 广告ID
	AdName                string  `json:"adName,omitempty"`                // 广告名称
	Keywords              string  `json:"keywords,omitempty"`              // 关键词
	AppID                 int64   `json:"appId,omitempty"`                 // 应用ID
	AppName               string  `json:"appName,omitempty"`               // 应用名称
	AdPrice               int64   `json:"adPrice,omitempty"`               // 广告实扣金额
	Expose                int64   `json:"expose,omitempty"`                // 曝光次数
	Click                 int64   `json:"click,omitempty"`                 // 点击次数
	Cost                  int64   `json:"cost,omitempty"`                  // 消耗(单位:分)
	ClickRate             string  `json:"clickRate,omitempty"`             // 点击率
	ClickPrice            string  `json:"clickPrice,omitempty"`            // 点击均价
	Download              int64   `json:"download,omitempty"`              // 下载次数
	DownloadRate          string  `json:"downloadRate,omitempty"`          // 下载率
	DownloadPrice         string  `json:"downloadPrice,omitempty"`         // 下载均价
	Ecpm                  string  `json:"ecpm,omitempty"`                  // ECPM
	BidRank               string  `json:"bidRank,omitempty"`               // 竞价排名
	FormOrderNums         int64   `json:"formOrderNums,omitempty"`         // 表单提交量
	FormPrice             string  `json:"formPrice,omitempty"`             // 表单提交成本
	ConvertActive         int64   `json:"convertActive,omitempty"`         // 回传激活量
	ConvertActivePrice    string  `json:"convertActivePrice,omitempty"`    // 回传激活成本
	ConvertRegister       int64   `json:"convertRegister,omitempty"`       // 注册量
	ConvertRegisterPrice  string  `json:"convertRegisterPrice,omitempty"`  // 注册成本
	ConvertRetention      int64   `json:"convertRetention,omitempty"`      // 回传留存
	ConvertRetentionPrice string  `json:"convertRetentionPrice,omitempty"` // 回传留存成本
	ConvertAppCredit      int64   `json:"convertAppCredit,omitempty"`      // 应用内授信
	ConvertAppCreditPrice string  `json:"convertAppCreditPrice,omitempty"` // 应用内授信成本
	ConvertAppOrder       int64   `json:"convertAppOrder,omitempty"`       // 应用内下单(电商)
	ConvertAppOrderPrice  string  `json:"convertAppOrderPrice,omitempty"`  // 应用内下单(电商成本)
	ConvertAppPay         int64   `json:"convertAppPay,omitempty"`         // 应用内付费
	ConvertAppPayPrice    string  `json:"convertAppPayPrice,omitempty"`    // 应用内付费成本
	ConvertAppCustom      int64   `json:"convertAppCustom,omitempty"`      // 应用内自定义转化
	ConvertAppCustomPrice string  `json:"convertAppCustomPrice,omitempty"` // 应用内自定义转化成本
	ActiveNums            int64   `json:"activeNums,omitempty"`            // 激活人数
	ActiveRate            string  `json:"activeRate,omitempty"`            // 激活率
	ActivePrice           string  `json:"activePrice,omitempty"`           // 激活成本
	InterBuyLtv1          float64 `json:"interBuyLtv1,omitempty"`          // 广告变现LTV1
	InterBuyLtv3          float64 `json:"interBuyLtv3,omitempty"`          // 广告变现LTV3
	InterBuyLtv7          float64 `json:"interBuyLtv7,omitempty"`          // 广告变现LTV7
	InterBuyLtv30         float64 `json:"interBuyLtv30,omitempty"`         // 广告变现LTV30
	InterBuyRoi1          string  `json:"interBuyRoi1,omitempty"`          // 广告变现ROI1
	InterBuyRoi3          string  `json:"interBuyRoi3,omitempty"`          // 广告变现ROI3
	InterBuyRoi7          string  `json:"interBuyRoi7,omitempty"`          // 广告变现ROI7
	InterBuyRoi30         string  `json:"interBuyRoi30,omitempty"`         // 广告变现ROI30
	BuyUserIncome0        float64 `json:"buyUserIncome0,omitempty"`        // 首日广告收入
	AdNewSsoid            int64   `json:"adNewSsoid,omitempty"`            // 注册人数
	RegisterRate          string  `json:"registerRate,omitempty"`          // 注册率
	RegisterPrice         string  `json:"registerPrice,omitempty"`         // 注册成本
	AdNewPaySsoid         int64   `json:"adNewPaySsoid,omitempty"`         // 新增付费人数
	NewPayRate            string  `json:"newPayRate,omitempty"`            // 付费转化率
	NewPayPrice           string  `json:"newPayPrice,omitempty"`           // 付费成本
	NewUserLtv1           float64 `json:"newUserLtv1,omitempty"`           // 内购LTV1
	NewUserLtv3           float64 `json:"newUserLtv3,omitempty"`           // 内购LTV3
	NewUserLtv7           float64 `json:"newUserLtv7,omitempty"`           // 内购LTV7
	NewUserLtv30          float64 `json:"newUserLtv30,omitempty"`          // 内购LTV30
	NewUserRoi1           string  `json:"newUserRoi1,omitempty"`           // 内购ROI1
	NewUserRoi3           string  `json:"newUserRoi3,omitempty"`           // 内购ROI3
	NewUserRoi7           string  `json:"newUserRoi7,omitempty"`           // 内购ROI7
	NewUserRoi30          string  `json:"newUserRoi30,omitempty"`          // 内购ROI30
	AdNewPay              int64   `json:"adNewPay,omitempty"`              // 首日付费流水
	TotalPay              int64   `json:"totalPay,omitempty"`              // 总付费流水
}
