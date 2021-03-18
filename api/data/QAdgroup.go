package data

import (
	"github.com/bububa/oppo-omni/core"
	"github.com/bububa/oppo-omni/model/data"
)

// 广告组-图表
func QAdgroup(clt *core.SDKClient, req *data.QAdgroupRequest) (*data.QResult, error) {
	req.SetResourceName("data")
	req.SetResourceAction("Q/adgroup")
	var ret data.QResponse
	err := clt.Post(req, &ret)
	if err != nil {
		return nil, err
	}
	return ret.Data, nil
}
