package data

import (
	"github.com/bububa/oppo-omni/core"
	"github.com/bububa/oppo-omni/model/data"
)

// 广告-列表
func QAdList(clt *core.SDKClient, req *data.QAdListRequest) (*data.QListResult, error) {
	req.SetResourceName("data")
	req.SetResourceAction("Q/ad/list")
	var ret data.QListResponse
	err := clt.Post(req, &ret)
	if err != nil {
		return nil, err
	}
	return ret.Data, nil
}
