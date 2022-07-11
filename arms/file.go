package arms

import (
	"github.com/spf13/cast"
	"io/ioutil"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"thh/arms/ehandle"
)

var basePath string

func SetBasePath(path string) {
	basePath = path
}

func StorageGet(filename string) string {
	data, _ := FileGetContents(basePath + filename)
	return cast.ToString(data)
}

func StoragePut(filename string, data any, append bool) error {
	return FilePutContents(basePath+filename, data, append)
}

// Put 将数据存入文件
func Put(data []byte, to string) (err error) {
	err = ioutil.WriteFile(to, data, 0644)
	return
}

func FileGetContents(filename string) ([]byte, error) {
	return ioutil.ReadFile(filename)
}

// FilePutContents file_put_contents
func FilePutContents(filename string, data any, isAppend ...bool) error {
	if dir := filepath.Dir(filename); dir != "" {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}
	}
	bData := []byte(cast.ToString(data))
	needAppend := false
	if len(isAppend) > 0 && isAppend[0] == true {
		needAppend = true
	}
	if needAppend {
		fl, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
		if err != nil {
			return err
		}
		defer func(fl *os.File) {
			ehandle.PrIF(fl.Close())
		}(fl)
		_, err = fl.Write(bData)
		return err
	} else {
		return ioutil.WriteFile(filename, bData, 0644)
	}
}

func IsExist(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

func UrlDecode(s string) string {
	r, err := url.QueryUnescape(s)
	if err != nil {
		return ""
	}
	return r
}

func UrlEncode(s string) string {
	return url.QueryEscape(s)
}

func FileNameWithoutExtension(fileName string) string {
	return strings.TrimSuffix(fileName, filepath.Ext(fileName))
}
