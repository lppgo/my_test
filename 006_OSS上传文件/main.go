package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {

}

func UploadFile(c *gin.Context) (err error) {
	//    通过form-data上传文件，文件名：file
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}

	fileHandle, err := file.Open() //打开上传文件
	if err != nil {
		return err
	}
	defer fileHandle.Close()
	fileByte, err := ioutil.ReadAll(fileHandle) //获取上传文件字节流
	if err != nil {
		return err
	}

	url, err := upload_service.Upload(file.Filename, fileByte)

	c.JSON(http.StatusOK, gin.H{
		"error":    "",
		"errno":    "0",
		"dataType": "OBJECT",
		"data": gin.H{
			"url": url,
		},
	})
	return nil
}

func Upload(fileName string, fileByte []byte) (url string, err error) {
	//    oss 的相关配置信息
	bucketName := config.GConfig.GetString("oss.bucket")
	endpoint := config.GConfig.GetString("oss.Endpoint")
	accessKeyId := config.GConfig.GetString("oss.AccessKeyId")
	accessKeySecret := config.GConfig.GetString("oss.AccessKeySecret")
	domain := config.GConfig.GetString("oss.domain")

	//创建OSSClient实例
	client, err := oss.New(endpoint, accessKeyId, accessKeySecret)
	if err != nil {
		return url, err
	}

	// 获取存储空间
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return url, err
	}

	//上传阿里云路径
	folderName := time.Now().Format("2006-01-02")
	yunFileTmpPath := filepath.Join("uploads", folderName) + "/" + fileName //uploads/2020-06-17/GoLang (1).docx

	// 上传Byte数组
	err = bucket.PutObject(yunFileTmpPath, bytes.NewReader([]byte(fileByte)))
	if err != nil {
		return url, err
	}

	return domain + "/" + yunFileTmpPath, nil
}
