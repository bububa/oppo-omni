package data

import (
	"github.com/bububa/oppo-omni/core"
	"github.com/bububa/oppo-omni/model/data"
)

// 广告组-列表
func QAdgroupList(clt *core.SDKClient, req *data.QAdgroupListRequest) (*data.QListResult, error) {
	req.SetResourceName("data")
	req.SetResourceAction("Q/adgroup/list")
	var ret data.QListResponse
	err := clt.Post(req, &ret)
	if err != nil {
		return nil, err
	}
	return ret.Data, nil
}
