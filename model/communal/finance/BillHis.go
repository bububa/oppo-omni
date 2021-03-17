package finance

import (
	"github.com/bububa/oppo-omni/enum"
	"github.com/bububa/oppo-omni/model"
)

type BillHisRequest struct {
	model.BaseRequest
	StartTime int64                 `json:"startTime,omitempty"` // 申请充值时间,查询开始时间
	EndTime   int64                 `json:"endTime,omitempty"`   // 申请充值时间 , 查询结束时间
	LoginType int64                 `json:"loginType,omitempty"` // 登录类型
	BillType  *enum.FinanceBillType `json:"billType,omitempty"`  // 流水类型; 0:存入, 1:支出
	BillSub   *enum.FinanceBillSub  `json:"billSub,omitempty"`   // 账务类型; 100:充值, 200:锁定, 201:解锁, 300:划账 (转账)
	SubAcc    *enum.FinanceSubAcc   `json:"subAcc,omitempty"`    // 账户类型; 0:现金账户; 1:返货账户; 9999:赠送账户
	Page      int                   `json:"page,omitempty"`      // 当前页
	PageSize  int                   `json:"pageCount,omitempty"` // 一页个数
}

type BillHisResponse struct {
	model.BaseResponse
	Data *BillHisResult `json:"data,omitempty"`
}

type BillHisResult struct {
	ItemCount  int           `json:"itemCount,omitempty"`  // 当前数量
	TotalCount int           `json:"totalCount,omitempty"` // 总数
	Items      []BillHisItem `json:"items,omitempty"`      // 数据列表
}

type BillHisItem struct {
	BillTime        int64                `json:"billTime,omitempty"` // 充值流水ID
	SubAcc          enum.FinanceSubAcc   `json:"subAcc,omitempty"`
	BillType        enum.FinanceBillType `json:"billType,omitempty"`
	BillSub         enum.FinanceBillSub  `json:"billSub,omitempty"`
	BillMoney       int64                `json:"billMoney,omitempty"`       // 充值金额 ,单位分
	BillDesc        string               `json:"billDesc,omitempty"`        // 流水描述
	AccID           uint64               `json:"accId,omitempty"`           // 当前用户账务ID
	RelateAccID     uint64               `json:"relateAccId,omitempty"`     // 客户账务ID (从账号账务ID)
	RelateOwnerID   uint64               `json:"relateOwnerId,omitempty"`   // 客户ID (从账号ID)
	RelateOwnerName string               `json:"relateOwnerName,omitempty"` // 客户名称 (从账号名称)
}
