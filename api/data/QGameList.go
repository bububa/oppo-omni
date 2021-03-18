package data

import (
	"github.com/bububa/oppo-omni/core"
	"github.com/bububa/oppo-omni/model/data"
)

// 游戏-列表
func QGameList(clt *core.SDKClient, req *data.QGameListRequest) (*data.QListResult, error) {
	req.SetResourceName("data")
	req.SetResourceAction("Q/game/list")
	var ret data.QListResponse
	err := clt.Post(req, &ret)
	if err != nil {
		return nil, err
	}
	return ret.Data, nil
}
