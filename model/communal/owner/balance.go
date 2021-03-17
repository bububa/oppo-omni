package owner

import "github.com/bububa/oppo-omni/model"

type BalanceResponse struct {
	model.BaseResponse
	Data []BalanceAccount `json:"data,omitempty"`
}

type BalanceAccount struct {
	AccID             uint64 `json:"accId,omitempty"`             // 账务ID
	AccDayBudget      int64  `json:"accDayBudget,omitempty"`      // 账户日预算(分)
	AccDayBudgetLimit int64  `json:"accDayBudgetLimit,omitempty"` // 账户日预算最低限额(分)
	CashBalance       int64  `json:"cashBal,omitempty"`           // 现金余额(分)
	CashLockBalance   int64  `json:"cashLockBal,omitempty"`       // 现金账户锁定金额(分)
	CashCost          int64  `json:"cashCost,omitempty"`          // 现金账户消耗(分)
	RebateBalance     int64  `json:"rebateBal,omitempty"`         // 返回账户余额(分)
	RebateCost        int64  `json:"rebateCost,omitempty"`        // 返回账户消耗(分)
	VirtualBalance    int64  `json:"virBal,omitempty"`            // 虚拟账户金额(分)
	VirtualCost       int64  `json:"virCost,omitempty"`           // 虚拟账户消耗(分)
}
