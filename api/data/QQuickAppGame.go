package data

import (
	"github.com/bububa/oppo-omni/core"
	"github.com/bububa/oppo-omni/model/data"
)

// 小游戏-图表
func QQuickAppGame(clt *core.SDKClient, req *data.QQuickAppGameRequest) (*data.QQuickAppGameResult, error) {
	req.SetResourceName("data")
	req.SetResourceAction("Q/quickApp/game")
	var ret data.QQuickAppGameResponse
	err := clt.Post(req, &ret)
	if err != nil {
		return nil, err
	}
	return ret.Data, nil
}
