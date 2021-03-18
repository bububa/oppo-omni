package data

import (
	"github.com/bububa/oppo-omni/core"
	"github.com/bububa/oppo-omni/model/data"
)

// 关键词-列表
func QKwList(clt *core.SDKClient, req *data.QKwListRequest) (*data.QListResult, error) {
	req.SetResourceName("data")
	req.SetResourceAction("Q/kw/list")
	var ret data.QListResponse
	err := clt.Post(req, &ret)
	if err != nil {
		return nil, err
	}
	return ret.Data, nil
}
