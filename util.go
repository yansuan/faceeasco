package faceeasco

import (
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"github.com/google/uuid"
)

const (
	Base64ImageHeader = "data:image/jpeg;base64,"
)

func Base64ToImage(base []byte) (result []byte, err error) {
	result, err = base64.StdEncoding.DecodeString(string(base))
	if err != nil {
		return
	}

	return
}

func ImageToBase64(img []byte) (result []byte, err error) {
	s := base64.StdEncoding.EncodeToString(img)
	s = Base64ImageHeader + s
	result = []byte(s)

	return
}

func FileRead(file string) (bytes []byte, err error) {
	bytes, err = ioutil.ReadFile(file)
	return
}
func FileDownload(url string, targetPath string) (err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	if !(resp.StatusCode >= 200 && resp.StatusCode < 300) {
		err = errors.New(fmt.Sprintf("%s", resp.Status))
		return
	}

	// 创建一个文件用于保存
	out, err := os.Create(targetPath)
	if err != nil {
		return
	}
	defer out.Close()

	// 然后将响应流和文件流对接起来
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return
	}
	return
}

func NewUUID() string {
	return uuid.NewString()
}
