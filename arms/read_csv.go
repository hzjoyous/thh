package arms

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func ReadCsv(filepath string) (c chan []string) {
	//打开文件(只读模式)，创建io.read接口实例
	c = make(chan []string)
	go func() {
		defer close(c)
		opencast, err := os.Open(filepath)
		if err != nil {
			log.Println("csv文件打开失败！")
		}
		defer func(opencast *os.File) {
			err := opencast.Close()
			if err != nil {
				fmt.Println(err)
			}
		}(opencast)
		ReadCsv := csv.NewReader(opencast)
		for {
			read, err := ReadCsv.Read()
			if err != nil {
				break
			}
			c <- read
		}
	}()

	//获取一行内容，一般为第一行内容
	//返回切片类型：[chen  hai wei]
	return
}
