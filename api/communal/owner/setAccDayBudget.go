package owner

import (
	"github.com/bububa/oppo-omni/core"
	"github.com/bububa/oppo-omni/model"
	"github.com/bububa/oppo-omni/model/communal/owner"
)

// 客户日预算设置
func SetAccDayBudget(clt *core.SDKClient, ownerID uint64, budget int64) error {
	var req owner.SetAccDayBudgetRequest
	req.AccDayBudget = budget
	req.SetOwnerID(ownerID)
	req.SetResourceName("communal")
	req.SetResourceAction("owner/setAccDayBudget")
	var ret model.BaseResponse
	return clt.Post(req, &ret)
}
