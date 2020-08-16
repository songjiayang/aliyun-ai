package face

import (
	"io/ioutil"
	"os"
	"testing"
	"time"

	"github.com/songjiayang/aliyun-ai/client"
	"github.com/stretchr/testify/assert"
)

func TestVerify(t *testing.T) {
	faceClient := NewClient(client.NewConfig(
		"https://dtplus-cn-shanghai.data.aliyuncs.com",
		os.Getenv("ALIYUN_ACCESS_KEY"),
		os.Getenv("ALIYUN_ACCESS_SECRET"),
		30*time.Second,
	))

	img1, _ := ioutil.ReadFile("./face1.png")
	img2, _ := ioutil.ReadFile("./face1.png")

	ret, err := faceClient.VerifyWithContent(img1, img2)

	assertion := assert.New(t)
	assertion.Nil(err)
	assertion.True(ret.IsOK())

	img3, _ := ioutil.ReadFile("./face2.png")

	ret, err = faceClient.VerifyWithContent(img1, img3)

	assertion.Nil(err)
	assertion.False(ret.IsOK())
}
