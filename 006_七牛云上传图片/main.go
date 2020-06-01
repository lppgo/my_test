package qiniu

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"

	"git.yeeuu.com/youjia/mountain/services/userServices"
	"git.yeeuu.com/youjia/mountain/utils/glog"
	"git.yeeuu.com/youjia/mountain/utils/qiniu"
	"github.com/labstack/echo"
	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
)

var (
	accessKey = "ydDXs_***********-ZGN"
	secretKey = "KIsR4Ugf**********Astk1sA_RBSfD"
	bucket    = "youjia-web" //存储空间
	img_type  = "png"
	host      = "https://web-cdn.***.com" //外链域名
)

// // 获取下载七牛图片连接
// func GetDownLoadUrl(userId int) string {
// 	key := fmt.Sprintf("user/img/%s", "") + "." + img_type
// 	mac := qbox.NewMac(accessKey, secretKey)
// 	deadline := time.Now().Add(time.Second * 3600).Unix() //1小时有效期
// 	privateAccessURL := storage.MakePrivateURL(mac, host, key, deadline)
// 	return privateAccessURL
// }

//路由处理函数
func UploadUserImages(c echo.Context) error {
	file, err := c.FormFile("protraitFile")
	uid, ok := c.Get("userId").(int)
	if !ok {
		glog.Logger().Println("获取currentUid错误")
		return c.JSON(http.StatusOK, Fail(UploadFileError))
	}
	if err != nil {
		glog.Logger().Println("FormFile错误", err.Error())
		return c.JSON(http.StatusOK, Fail(UploadFileError))
	}
	src, err := file.Open()
	if err != nil {
		glog.Logger().Println("OpenFile错误", err.Error())
		return c.JSON(http.StatusOK, Fail(OpenFileError))
	}
	defer src.Close()
	dst, err := os.Create(file.Filename)
	if err != nil {
		glog.Logger().Println("Os.Create错误", err.Error())
		return c.JSON(http.StatusOK, Fail(UploadFileError))
	}
	defer dst.Close()
	if _, err := io.Copy(dst, src); err != nil {
		glog.Logger().Println("io.Copy错误", err.Error())
		return c.JSON(http.StatusOK, Fail(UploadFileError))
	}
	defer os.Remove(file.Filename)
	url, err := qiniu.UploadImage(file.Filename)
	if err != nil {
		glog.Logger().Println("上传文件到七牛云错误", err.Error())
		return c.JSON(http.StatusOK, Fail(UploadFileError))
	}
	glog.Logger().Println("七牛云上传图片URL", url)
	//更新用户
	err = userServices.UpdateUserPortrait(uid, url)
	if err != nil {
		glog.Logger().Println(err.Error())
		// return c.JSON(http.StatusOK, Fail(UploadFileError))
	}
	return c.JSON(http.StatusOK, Success(url))
}

// 七牛上传头像工具
func UploadImage(localFile string) (url string, err error) {
	// uid := utils.GenUUIDById()
	timstamp := time.Now().Unix()
	upTimestamp := strconv.FormatInt(timstamp, 10)
	key := fmt.Sprintf("user/img/%s_%s", upTimestamp, localFile)
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)
	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Zone = &storage.ZoneHuadong
	// 是否使用https域名
	cfg.UseHTTPS = false
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = true
	// 构建表单上传的对象
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	// 可选配置
	putExtra := storage.PutExtra{}
	err = formUploader.PutFile(context.Background(), &ret, upToken, key, localFile, &putExtra)
	if err != nil {
		glog.Logger().Println("上传图片到七牛云错误:", err)
		return url, err
	}
	url = host + "/" + ret.Key
	return
}

//获取7牛token
func GetUpToken() string {
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)
	// fmt.Println("---------------")
	// fmt.Println(upToken)
	// fmt.Println("---------------")
	return upToken
}
