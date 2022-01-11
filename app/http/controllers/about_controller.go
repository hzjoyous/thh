package controllers

import (
	"bytes"
	"fmt"
	"github.com/fogleman/gg"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"thh/helpers/db"
	Logger "thh/helpers/logger"
	"time"
)

func About(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello~ Now you see a json from gin",
	})
}

func ShowPic(c *gin.Context) {

	const S = 1024
	dc := gg.NewContext(S, S)
	dc.SetRGBA(0, 0, 0, 0.1)
	for i := 0; i < 360; i += 15 {
		dc.Push()
		dc.RotateAbout(gg.Radians(float64(i)), S/2, S/2)
		dc.DrawEllipse(S/2, S/2, S*7/16, S/8)
		dc.Fill()
		dc.Pop()
	}

	w := c.Writer

	buffer := bytes.NewBuffer(nil)
	err := dc.EncodePNG(buffer)
	if err != nil {
		return
	}

	b := buffer.Bytes()

	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")
	w.Header().Set("Content-Type", "image/png")
	http.ServeContent(w, c.Request, "name", time.Time{}, bytes.NewReader(b))
}

func ShowPic2(c *gin.Context) {
	const S = 1024
	dc := gg.NewContext(S, S)
	dc.SetRGBA(0, 0, 0, 0.1)
	for i := 0; i < 360; i += 15 {
		dc.Push()
		dc.RotateAbout(gg.Radians(float64(i)), S/2, S/2)
		dc.DrawEllipse(S/2, S/2, S*7/16, S/8)
		dc.Fill()
		dc.Pop()
	}

	buffer := bytes.NewBuffer(nil)
	err := dc.EncodePNG(buffer)
	if err != nil {
		return
	}

	b := buffer.Bytes()
	c.Header("Cache-Control", "no-cache, no-store, must-revalidate")
	c.Header("Pragma", "no-cache")
	c.Header("Expires", "0")
	c.Data(http.StatusOK, "image/png", b)

}

func SetData(c *gin.Context) {
	key := c.DefaultQuery("key", "key")
	data := c.DefaultQuery("dataRep", "dataRep")

	DB.KVDB().Set(key, data)

	c.JSON(http.StatusOK, gin.H{
		"key":  key,
		"dataRep": data,
	})
}

func GetData(c *gin.Context) {
	key := c.DefaultQuery("key", "key")
	entry := DB.KVDB().Get(key)

	c.JSON(http.StatusOK, gin.H{
		"message": string(entry),
	})
}

func Upload(ctx *gin.Context) {
	name := ctx.PostForm("name")
	file, err := ctx.FormFile("file")
	if err != nil {
		fmt.Println("获取数据失败")
		ctx.JSON(http.StatusOK, gin.H{
			"code":    1,
			"message": "获取数据失败" + err.Error(),
		})
	} else {
		fmt.Println("接收的数据", name, file.Filename)
		//获取文件名称
		fmt.Println(file.Filename)
		//文件大小
		fmt.Println(file.Size)
		//获取文件的后缀名
		extstring := path.Ext(file.Filename)
		fmt.Println(extstring)
		//根据当前时间鹾生成一个新的文件名
		fileNameInt := time.Now().Unix()
		fileNameStr := strconv.FormatInt(fileNameInt, 10)
		//新的文件名
		fileName := fileNameStr + extstring
		//保存上传文件
		filePath := filepath.Join(Mkdir("upload"), "/", fileName)
		ctx.SaveUploadedFile(file, filePath)
		ctx.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "success",
		})
	}
}
func UploadAll(ctx *gin.Context) {
	if form, err := ctx.MultipartForm(); err == nil {
		//1.获取文件
		files := form.File["file"]
		//2.循环全部的文件
		for _, file := range files {
			// 3.根据时间鹾生成文件名
			fileNameInt := time.Now().Unix()
			fileNameStr := strconv.FormatInt(fileNameInt, 10)
			//4.新的文件名(如果是同时上传多张图片的时候就会同名，因此这里使用时间鹾加文件名方式)
			fileName := fileNameStr + file.Filename
			//5.保存上传文件
			filePath := filepath.Join(Mkdir("upload"), "/", fileName)
			ctx.SaveUploadedFile(file, filePath)
		}
		ctx.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "上传成功",
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "获取数据失败",
		})
	}
}
func Mkdir(basePath string) string {
	//	1.获取当前时间,并且格式化时间
	folderName := time.Now().Format("2006/01/02")
	folderPath := filepath.Join(basePath, folderName)
	//使用mkdirall会创建多层级目录
	os.MkdirAll(folderPath, os.ModePerm)
	return folderPath
}

func LoggerTest(c *gin.Context) {
	Logger.Info("日志日志日志日志日志日志日志日志日志日志")
	Logger.Info("日志日志日志日志日志日志日志日志日志日志")
	Logger.Info("日志日志日志日志日志日志日志日志日志日志")
	Logger.Info("日志日志日志日志日志日志日志日志日志日志")
	Logger.Info("日志日志日志日志日志日志日志日志日志日志")
	c.JSON(200, "ok")
}
