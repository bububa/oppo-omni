package data

import (
	"github.com/bububa/oppo-omni/core"
	"github.com/bububa/oppo-omni/model/data"
)

// 小游戏-列表
func QQuickAppGameList(clt *core.SDKClient, req *data.QQuickAppGameListRequest) (*data.QQuickAppGameListResult, error) {
	req.SetResourceName("data")
	req.SetResourceAction("Q/quickApp/game/list")
	var ret data.QQuickAppGameListResponse
	err := clt.Post(req, &ret)
	if err != nil {
		return nil, err
	}
	return ret.Data, nil
}
