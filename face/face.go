package face

import "github.com/songjiayang/aliyun-ai/client"

type Client struct {
	*client.Client
}

func NewClient(cfg *client.Config) *Client {
	return &Client{
		Client: client.New(cfg),
	}
}
