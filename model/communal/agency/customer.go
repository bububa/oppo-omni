package agency

import (
	"github.com/bububa/oppo-omni/enum"
	"github.com/bububa/oppo-omni/model"
)

type CustomerListRequest struct {
	model.BaseRequest
	ID          uint64                    `json:"id,omitempty"`        // 定向ID
	OwnerName   string                    `json:"ownerName,omitempty"` // 广告主ID
	AuditStatus *enum.CustomerAuditStatus `json:"audit_status"`        // 审核状态
	Page        int                       `json:"page,omitempty"`      // 当前页
	PageSize    int                       `json:"pageCount,omitempty"` // 每页数量
}

type CustomerListResponse struct {
	model.BaseResponse
	Data *CustomerListResult `json:"data,omitempty"`
}

type CustomerListResult struct {
	ItemCount  int        `json:"itemCount,omitempty"`
	TotalCount int        `json:"totalCount,omitempty"`
	Items      []Customer `json:"items,omitempty"`
}

type Customer struct {
	OwnerID         uint64                   `json:"ownerId,omitempty"`         // 广告主ID
	AccID           uint64                   `json:"accId,omitempty"`           // 账务ID
	OwnerName       string                   `json:"ownerName,omitempty"`       // 广告主名称
	AuditStatus     enum.CustomerAuditStatus `json:"auditStatus,omitempty"`     // 审核状态
	InsertTime      int64                    `json:"insertTime,omitempty"`      // 创建时间(秒)
	UpdateTime      int64                    `json:"updateTime,omitempty"`      // 最后更新时间(秒)
	CashBalance     int64                    `json:"cashBal,omitempty"`         // 现金余额(分)
	CashLockBalance int64                    `json:"cashLockBal,omitempty"`     // 现金账户锁定金额(分)
	CashCost        int64                    `json:"cashCost,omitempty"`        // 现金账户消耗(分)
	RebateBalance   int64                    `json:"rebateBal,omitempty"`       // 返回账户余额(分)
	RebateDayBudget int64                    `json:"rebateDayBudget,omitempty"` // 返回账户日预算(分)
	RebateCost      int64                    `json:"rebateCost,omitempty"`      // 返回账户消耗(分)
	VirtualBalance  int64                    `json:"virBal,omitempty"`          // 虚拟账户金额(分)
	VirtalCost      int64                    `json:"virCost,omitempty"`         // 虚拟账户消耗(分)
}
