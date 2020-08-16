package client

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/qiniu/x/bytes/seekable"
)

type Client struct {
	cfg *Config

	*http.Client
}

func New(cfg *Config) *Client {
	client := &Client{
		cfg: cfg,
	}

	client.Client = &http.Client{
		Transport: client,
		Timeout:   cfg.Timeout,
	}

	return client
}

func (client *Client) Send(method, path, contentType string, body io.Reader) (data []byte, err error) {
	request, err := http.NewRequest(method, client.cfg.Host+path, body)
	if err != nil {
		return
	}

	request.Header.Set("Content-Type", contentType)
	request.Header.Set("Accept", contentType)
	request.Header.Set("Date", time.Now().Format("Mon, 02 Jan 2006 15:04:05 GMT"))

	response, err := client.Do(request)
	if err != nil {
		return
	}
	defer response.Body.Close()

	data, err = ioutil.ReadAll(response.Body)
	if response.StatusCode/100 != 2 {
		err = errors.New(string(data))
		data = nil
	}

	return
}

func (client *Client) RoundTrip(req *http.Request) (resp *http.Response, err error) {
	err = client.AuthRequest(req)
	if err != nil {
		return
	}

	return http.DefaultTransport.RoundTrip(req)
}

func (client *Client) AuthRequest(req *http.Request) (err error) {
	bodyMd5 := md5.New()
	ctype := req.Header.Get("Content-Type")

	switch {
	case req.ContentLength != 0 && req.Body != nil &&
		ctype != "" && ctype != "application/octet-stream":
		seeker, tmperr := seekable.New(req)
		if tmperr != nil {
			err = tmperr
			return
		}

		bodyMd5.Write(seeker.Bytes())
	}

	hash := hmac.New(sha1.New, client.cfg.AccessKeySecret)

	io.WriteString(hash, req.Method+"\n")
	io.WriteString(hash, req.Header.Get("Accept")+"\n")
	io.WriteString(hash, base64.StdEncoding.EncodeToString(bodyMd5.Sum(nil))+"\n")
	io.WriteString(hash, ctype+"\n")
	io.WriteString(hash, req.Header.Get("Date")+"\n")
	io.WriteString(hash, req.URL.Path)

	sign := base64.StdEncoding.EncodeToString(hash.Sum(nil))

	req.Header.Set("Authorization", "Dataplus "+client.cfg.AccessKey+":"+sign)

	return nil
}
