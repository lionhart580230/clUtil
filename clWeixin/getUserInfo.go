package clWeixin

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/lionhart580230/clUtil/clHttpClient"
)

type WXRespUserInfo struct {
	ErrorCode  uint32   `json:"error_code"`
	ErrMsg     string   `json:"err_msg"`
	OpenId     string   `json:"openid"`
	Nickname   string   `json:"nickname"`
	Sex        uint32   `json:"sex"`
	Province   string   `json:"province"`
	City       string   `json:"city"`
	Country    string   `json:"country"`
	HeadImgUrl string   `json:"headimgurl"`
	Privilege  []string `json:"privilege"`
	UnionId    string   `json:"unionid"`
}

// 获取用户信息
func GetUserInfo(_appId string, _appSecret string, _code string) (error, *WXRespUserInfo) {
	err, accessTokenInfo := UserCodeToAccessToken(_appId, _appSecret, _code)
	if err != nil {
		return err, nil
	}

	client := clHttpClient.NewClient("https://api.weixin.qq.com/sns/userinfo")
	client.AddParam("access_token", accessTokenInfo.AccessToken)
	client.AddParam("openid", accessTokenInfo.OpenId)
	client.AddParam("lang", "zh_CN")

	res, err := client.Do()
	if err != nil {
		return err, nil
	}
	var respObj WXRespUserInfo
	if err := json.Unmarshal([]byte(res.Body), &respObj); err != nil {
		return err, nil
	}
	if respObj.ErrorCode != 0 {
		return errors.New(fmt.Sprintf("微信返回错误: %v - %v", respObj.ErrorCode, respObj.ErrMsg)), nil
	}
	return nil, &respObj
}
