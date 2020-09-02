package wxloginDecrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"git.yeeuu.com/youjia/mountain/utils/respErrors"

	"github.com/labstack/echo/v4"
)

// 微信登陆 登陆解密数据

type params struct {
	SessionKey string `json:"sessionKey"`
}

type code2Session struct {
	OpenId     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionId    string `json:"unionid"`
	ErrCode    int    `json:"errcode"`
	ErrMsg     string `json:"errmsg"`
}

var (
	grant_type    = "authorization_code"
	youjia_appId  = "wx***065c24"
	youjia_secret = "e2b4***f8eac16"
)

func (c code2Session) String() string {
	by, _ := json.Marshal(c)
	return string(by)
}

//微信登录前端请求数据
type WxreqData struct {
	EncryptedData string `json:"encryptedData"`
	Iv            string `json:"iv"`
	Code          string `json:"code"`
}

type WxLoginData struct {
	PhoneNumber     string `json:"phoneNumber"`
	PurePhoneNumber string `json:"purePhoneNumber"`
	CountryCode     string `json:"countryCode"`
	// Watermark       struct {
	// 	AppId     string `json:"appid"`
	// 	Timestamp string `json:"timestamp"`
	// } `json:"watermark"`
}

func AesCBCDecrypt() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			var wxreq WxreqData
			if err := c.Bind(&wxreq); err != nil {
				e := respErrors.New(10005, "用户尚未登陆1")
				return c.JSON(http.StatusOK, e)
			}
			s := params{}
			by, openid, err := s.decrypt(wxreq)
			if err != nil {
				e := respErrors.New(10005, "用户尚未登陆2")
				return c.JSON(http.StatusOK, e)
			}
			var wxdata WxLoginData
			if err := json.Unmarshal(by, &wxdata); err != nil {
				e := respErrors.New(10005, "用户尚未登陆3")
				return c.JSON(http.StatusOK, e)
			}
			fmt.Println("-----------wxdata------------")
			log.Printf("%+v\n", wxdata)
			fmt.Println("---------wxdata--------------")
			c.Set("wxRawData", wxdata)
			c.Set("openid", openid)
			return next(c)
		}
	}
}

//解码微信登录数据
func (p params) decrypt(wx WxreqData) ([]byte, string, error) {
	//利用code登录微信服务器，获取SessionKey和OpenId
	code2 := p.getRequestUrl(wx.Code)
	if code2.SessionKey == "" {
		fmt.Println("获取sissionKey失败")
		return nil, "", errors.New("获取sissionKey失败")
	}
	if code2.OpenId == "" {
		fmt.Println("获取OpenId失败")
		return nil, "", errors.New("获取OpenId失败")
	}
	fmt.Println("=====Session_key=====:", code2.SessionKey)
	fmt.Println("=====Open_id=========:", code2.OpenId)
	//aes解码微信的公开数据
	sessionKeyBytes, err := base64.StdEncoding.DecodeString(code2.SessionKey)
	if err != nil {
		return nil, "", err
	}
	decodeBytes, err := base64.StdEncoding.DecodeString(wx.EncryptedData)
	if err != nil {
		return nil, "", err
	}
	ivBytes, err := base64.StdEncoding.DecodeString(wx.Iv)
	if err != nil {
		return nil, "", err
	}
	//
	block, err := aes.NewCipher(sessionKeyBytes)
	if err != nil {
		return nil, "", err
	}
	//blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, ivBytes)
	origData := make([]byte, len(decodeBytes))
	blockMode.CryptBlocks(origData, decodeBytes)
	//获取的数据尾端有'/x0e'占位符,去除它
	for i, ch := range origData {
		if ch == '\x0e' {
			origData[i] = ' '
		}
	}
	fmt.Println("-----===5===---")
	fmt.Println(string(origData))
	fmt.Println("-----===5===---")
	return origData, code2.OpenId, nil
}

func (p params) getRequestUrl(code string) (code2 code2Session) {
	fmt.Println("code", code)
	u := "https://api.weixin.qq.com/sns/jscode2session?" + "appid=" + youjia_appId + "&secret=" + youjia_secret + "&js_code=" + code + "&grant_type=" + grant_type
	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return code2
	}
	req.Header.Set("accept", "*/*")
	req.Header.Set("user-agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1")
	cl := &http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := cl.Do(req)
	if err != nil {
		log.Println(err)
		return code2
	}
	log.Println(resp.Header)
	if err := json.NewDecoder(resp.Body).Decode(&code2); err != nil {
		return code2
	}
	fmt.Printf("code2:", code2) //打印返回结果
	return code2
}
