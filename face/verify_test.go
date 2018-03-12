package face

import (
	"encoding/base64"
	"io/ioutil"
	"os"
	"testing"
	"time"

	"github.com/golib/assert"
	"github.com/songjiayang/aliyun-ai/client"
)

func TestVerify(t *testing.T) {
	faceClient := NewClient(client.NewConfig(
		"https://dtplus-cn-shanghai.data.aliyuncs.com",
		os.Getenv("ALIYUN_ACCESS_KEY"),
		os.Getenv("ALIYUN_ACCESS_SECRET"),
		30*time.Second,
	))

	image1, _ := ioutil.ReadFile("./face1.png")
	image2, _ := ioutil.ReadFile("./face1.png")

	img1 := base64.StdEncoding.EncodeToString(image1)
	img2 := base64.StdEncoding.EncodeToString(image2)

	ret, err := faceClient.Verify(ImageTypeContent, img1, img2)

	assertion := assert.New(t)
	assertion.Nil(err)
	assertion.True(ret.IsOK())

	image3, _ := ioutil.ReadFile("./face2.png")
	img3 := base64.StdEncoding.EncodeToString(image3)

	ret, err = faceClient.Verify(ImageTypeContent, img1, img3)

	assertion.Nil(err)
	assertion.False(ret.IsOK())
}
