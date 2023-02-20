package api

import (
	"crypto/sha1"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/go-resty/resty/v2"
)

type GooseClient struct {
	USER  string `json:"user"`     // 飞鹅云后台注册用户名
	UKEY  string `json:"user_key"` // 实名认证ukey
	SN    string `json:"sn"`       // 打印机SN码
	Debug string `json:"debug"`    // 是否debug
}

// CommonRequestParams 构建公共请求参数
func (c *GooseClient) commonRequestParams() (map[string]string, error) {
	if len(c.USER) == 0 {
		return nil, errors.New("user is required")
	}
	if len(c.UKEY) == 0 {
		return nil, errors.New("ukey is required")
	}
	if len(c.SN) == 0 {
		return nil, errors.New("sn is required")
	}
	if len(c.Debug) == 0 {
		c.Debug = "0"
	}
	sign, sTime := c.sign()
	value := make(map[string]string, 0)
	value["user"] = c.USER
	value["stime"] = sTime
	value["sig"] = sign
	value["debug"] = c.Debug // 测试为 "1"
	return value, nil
}

// CommonHttpRequest 公共请求
func (c *GooseClient) commonHttpRequest(value map[string]string) (string, error) {
	clint := resty.New()
	resp, err := clint.R().
		SetFormData(value).
		ForceContentType("application/x-www-form-urlencoded").
		Post(URL)
	if err != nil {
		return "", err
	}

	if resp.StatusCode() != http.StatusOK {
		return "", errors.New(fmt.Sprintf("request fly false %d", resp.StatusCode()))
	}

	var printResp CommonResponse
	err = json.Unmarshal(resp.Body(), &printResp)
	if err != nil {
		return "", errors.New(fmt.Sprintf("unmarshal resp fasle,err=%v", err))
	}
	if printResp.Ret != 0 {
		return "", errors.New(fmt.Sprintf("err:%s", printResp.Data))
	}
	return printResp.Data, nil
}

// Sign 创建签名/返回签名/签名时间
func (c *GooseClient) sign() (sign string, sTime string) {
	t := time.Now().Unix()
	sTime = strconv.FormatInt(t, 10)
	key := c.USER + c.UKEY + sTime
	h := sha1.New()
	h.Write([]byte(key))
	sign = fmt.Sprintf("%x", h.Sum(nil))
	return sign, sTime
}

// PrintMSG 打印小票订单
func (c *GooseClient) PrintMSG(content string) (string, error) {
	value, err := c.commonRequestParams()
	if err != nil {
		return "", err
	}
	value["apiname"] = PrintMSG
	value["sn"] = c.SN
	value["content"] = content
	return c.commonHttpRequest(value)
}

// DelSqs 清空待打印信息队列
func (c *GooseClient) DelSqs() (string, error) {
	value, err := c.commonRequestParams()
	if err != nil {
		return "", err
	}
	value["sn"] = c.SN
	value["apiname"] = DelSqs

	return c.commonHttpRequest(value)
}

// OrderStatus 订单状态
func (c *GooseClient) OrderStatus(orderId string) (string, error) {
	value, err := c.commonRequestParams()
	if err != nil {
		return "", err
	}
	value["appname"] = OrderStatus
	value["orderid"] = orderId

	return c.commonHttpRequest(value)
}

// PrinterStatus 打印机状态
func (c *GooseClient) PrinterStatus() (string, error) {
	value, err := c.commonRequestParams()
	if err != nil {

	}
	value["sn"] = c.SN
	value["appname"] = PrinterStatus

	return c.commonHttpRequest(value)
}
