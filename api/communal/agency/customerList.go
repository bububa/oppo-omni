package agency

import (
	"github.com/bububa/oppo-omni/core"
	"github.com/bububa/oppo-omni/model/communal/agency"
)

// 代理商客户查询
func CustomerList(clt *core.SDKClient, req *agency.CustomerListRequest) (*agency.CustomerListResult, error) {
	req.SetResourceName("communal")
	req.SetResourceAction("agency/customer/list")
	var ret agency.CustomerListResponse
	err := clt.Post(req, &ret)
	if err != nil {
		return nil, err
	}
	return ret.Data, nil
}
