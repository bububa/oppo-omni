package agency

import (
	"github.com/bububa/oppo-omni/core"
	"github.com/bububa/oppo-omni/model"
	"github.com/bububa/oppo-omni/model/communal/agency"
)

// 代理商余额查询
func Balance(clt *core.SDKClient) ([]agency.BalanceAccount, error) {
	var req model.BaseRequest
	req.SetResourceName("communal")
	req.SetResourceAction("agency/balance")
	var ret agency.BalanceResponse
	err := clt.Post(req, &ret)
	if err != nil {
		return nil, err
	}
	return ret.Data, nil
}
