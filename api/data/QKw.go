package data

import (
	"github.com/bububa/oppo-omni/core"
	"github.com/bububa/oppo-omni/model/data"
)

// 关键词-图表
func QKw(clt *core.SDKClient, req *data.QKwRequest) (*data.QResult, error) {
	req.SetResourceName("data")
	req.SetResourceAction("Q/kw")
	var ret data.QResponse
	err := clt.Post(req, &ret)
	if err != nil {
		return nil, err
	}
	return ret.Data, nil
}
