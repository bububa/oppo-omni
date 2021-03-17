package clue

import (
	"github.com/bububa/oppo-omni/core"
	"github.com/bububa/oppo-omni/model"
	"github.com/bububa/oppo-omni/model/clue"
)

// H5 数据回传转化
func SendData(clt *core.SDKClient, req *clue.SendDataRequest) error {
	var ret model.BaseResponse
	req.SetBaseUrl("https://sapi.ads.oppomobile.com")
	req.SetVersion("v1")
	req.SetResourceName("clue")
	req.SetResourceAction("sendData")
	return clt.Post(req, &ret)
}
