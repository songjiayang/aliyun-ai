package face

import "github.com/songjiayang/aliyun-ai/client"

const (
	ImageTypeUrl     = "0"
	ImageTypeContent = "1"
)

type Client struct {
	*client.Client
}

func NewClient(cfg *client.Config) *Client {
	return &Client{
		Client: client.New(cfg),
	}
}
