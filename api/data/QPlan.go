package data

import (
	"github.com/bububa/oppo-omni/core"
	"github.com/bububa/oppo-omni/model/data"
)

// 广告计划-图表
func QPlan(clt *core.SDKClient, req *data.QPlanRequest) (*data.QResult, error) {
	req.SetResourceName("data")
	req.SetResourceAction("Q/plan")
	var ret data.QResponse
	err := clt.Post(req, &ret)
	if err != nil {
		return nil, err
	}
	return ret.Data, nil
}
