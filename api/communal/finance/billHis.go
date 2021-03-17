package owner

import (
	"github.com/bububa/oppo-omni/core"
	"github.com/bububa/oppo-omni/model/communal/finance"
)

// 账务流水查询
func BillHis(clt *core.SDKClient, req *finance.BillHisRequest) (*finance.BillHisResult, error) {
	req.SetResourceName("communal")
	req.SetResourceAction("finance/billHis")
	var ret finance.BillHisResponse
	err := clt.Post(req, &ret)
	if err != nil {
		return nil, err
	}
	return ret.Data, nil
}
