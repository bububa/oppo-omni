package data

import (
	"github.com/bububa/oppo-omni/core"
	"github.com/bububa/oppo-omni/model/data"
)

// 广告计划-列表
func QPlanList(clt *core.SDKClient, req *data.QPlanListRequest) (*data.QListResult, error) {
	req.SetResourceName("data")
	req.SetResourceAction("Q/plan/list")
	var ret data.QListResponse
	err := clt.Post(req, &ret)
	if err != nil {
		return nil, err
	}
	return ret.Data, nil
}
