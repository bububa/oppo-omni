package data

import (
	"github.com/bububa/oppo-omni/core"
	"github.com/bububa/oppo-omni/model/data"
)

// 广告-图表
func QAd(clt *core.SDKClient, req *data.QAdRequest) (*data.QResult, error) {
	req.SetResourceName("data")
	req.SetResourceAction("Q/ad")
	var ret data.QResponse
	err := clt.Post(req, &ret)
	if err != nil {
		return nil, err
	}
	return ret.Data, nil
}
