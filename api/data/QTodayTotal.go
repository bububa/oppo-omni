package data

import (
	"github.com/bububa/oppo-omni/core"
	"github.com/bububa/oppo-omni/model/data"
)

// 首页-整体数据
func QTodayTotal(clt *core.SDKClient, ownerID uint64, adID uint64) (*data.QTodayTotalResult, error) {
	var req data.QTodayTotalRequest
	req.SetOwnerID(ownerID)
	req.AdID = adID
	req.SetResourceName("data")
	req.SetResourceAction("Q/today/total")
	var ret data.QTodayTotalResponse
	err := clt.Post(req, &ret)
	if err != nil {
		return nil, err
	}
	return ret.Data, nil
}
