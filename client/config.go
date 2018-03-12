package client

import "time"

type Config struct {
	Host string

	AccessKey       string
	AccessKeySecret []byte

	Timeout time.Duration
}

func NewConfig(host, id, secret string, timeout time.Duration) *Config {
	return &Config{
		Host:            host,
		AccessKey:       id,
		AccessKeySecret: []byte(secret),
		Timeout:         timeout,
	}
}
