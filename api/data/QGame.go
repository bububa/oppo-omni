package data

import (
	"github.com/bububa/oppo-omni/core"
	"github.com/bububa/oppo-omni/model/data"
)

// 游戏-图表
func QGame(clt *core.SDKClient, req *data.QGameRequest) (*data.QGameResult, error) {
	req.SetResourceName("data")
	req.SetResourceAction("Q/game")
	var ret data.QGameResponse
	err := clt.Post(req, &ret)
	if err != nil {
		return nil, err
	}
	return ret.Data, nil
}
