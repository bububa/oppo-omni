package core

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/bububa/oppo-omni/core/internal/debug"
	"github.com/bububa/oppo-omni/model"
	"github.com/bububa/oppo-omni/util"
)

type SDKClient struct {
	AppID  string
	AppKey string
	debug  bool
}

func NewSDKClient(appID string, appKey string) *SDKClient {
	return &SDKClient{
		AppID:  appID,
		AppKey: appKey,
	}
}

func (c *SDKClient) SetDebug(debug bool) {
	c.debug = debug
}

func (c *SDKClient) token(ownerID uint64, ts int64) string {
	rawToken := fmt.Sprintf("%d,%s,%d,%s", ownerID, c.AppID, ts, c.sign(ts))
	return base64.StdEncoding.EncodeToString([]byte(rawToken))
}

func (c *SDKClient) sign(ts int64) string {
	return util.Sha1(fmt.Sprintf("%s%s%d", c.AppID, c.AppKey, ts))
}

func (c *SDKClient) Post(req model.Request, resp model.Response) error {
	reqBytes, _ := json.Marshal(req)
	reqURL := fmt.Sprintf("%s/%s/%s/%s", req.GetBaseUrl(), req.GetVersion(), req.ResourceName(), req.ResourceAction())
	ts := time.Now().Unix()
	token := c.token(req.GetOwnerID(), ts)
	debug.PrintPostJSONRequest(reqURL, reqBytes, c.debug)
	httpReq, err := http.NewRequest("POST", reqURL, bytes.NewReader(reqBytes))
	if err != nil {
		return err
	}
	httpReq.Header.Add("Content-Type", req.GetContentType())
	httpReq.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	httpResp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return err
	}
	defer httpResp.Body.Close()
	err = debug.DecodeJSONHttpResponse(httpResp.Body, resp, c.debug)
	if err != nil {
		debug.PrintError(err, c.debug)
		return err
	}
	if resp.IsError() {
		return resp
	}
	return nil
}
