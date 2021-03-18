package data

import (
	"github.com/bububa/oppo-omni/core"
	"github.com/bububa/oppo-omni/enum"
	"github.com/bububa/oppo-omni/model/data"
)

// 首页-TOP 榜
func QTodayTop(clt *core.SDKClient, ownerID uint64, deminsion *enum.DataDemision) (*data.QTodayTopResult, error) {
	var req data.QTodayTopRequest
	req.SetOwnerID(ownerID)
	req.SetResourceName("data")
	req.SetResourceAction("Q/today/top")
	req.Demision = deminsion
	var ret data.QTodayTopResponse
	err := clt.Post(req, &ret)
	if err != nil {
		return nil, err
	}
	return ret.Data, nil
}
