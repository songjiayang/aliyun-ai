package face

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/songjiayang/aliyun-ai/client"
)

// detailt to see https://help.aliyun.com/knowledge_detail/53535.html
const (
	verifyAPI = "/face/verify"
)

type VerfiyResult struct {
	Errno      int       `json:"errno"`
	ErrMsg     string    `json:"err_msg"`
	RequestId  string    `json:"request_id"`
	Confidence float64   `json:"confidence"`
	Thresholds []float64 `json:"thresholds"`
}

func (out *VerfiyResult) IsOK(threshold ...float64) bool {
	checkThreshold := float64(65)

	if len(threshold) > 0 {
		checkThreshold = threshold[0]
	}

	return out.Confidence > checkThreshold
}

func (c *Client) Verify(t, img1, img2 string) (ret *VerfiyResult, err error) {
	params := map[string]string{
		"type": t,
	}

	if t == ImageTypeUrl {
		params["image_url_1"] = img1
		params["image_url_2"] = img2
	} else {
		params["content_1"] = img1
		params["content_2"] = img2
	}

	buf, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	body, err := c.Send(http.MethodPost, verifyAPI, client.ContentTypeJson, bytes.NewReader(buf))
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &ret)
	if err != nil {
		return
	}

	if ret.Errno != 0 {
		return nil, errors.New(ret.ErrMsg)
	}

	return
}
