package owner

import (
	"github.com/bububa/oppo-omni/core"
	"github.com/bububa/oppo-omni/model"
	"github.com/bububa/oppo-omni/model/communal/owner"
)

// 客户余额查询
func Balance(clt *core.SDKClient, ownerID uint64) ([]owner.BalanceAccount, error) {
	var req model.BaseRequest
	req.SetOwnerID(ownerID)
	req.SetResourceName("communal")
	req.SetResourceAction("owner/balance")
	var ret owner.BalanceResponse
	err := clt.Post(req, &ret)
	if err != nil {
		return nil, err
	}
	return ret.Data, nil
}
