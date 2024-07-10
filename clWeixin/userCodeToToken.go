package clWeixin

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/lionhart580230/clUtil/clHttpClient"
)

type WxResUserCodeToToken struct {
	ErrCode        uint32 `json:"errcode"`
	ErrMsg         string `json:"errmsg"`
	AccessToken    string `json:"access_token"`
	ExpiresIn      uint32 `json:"expires_in"`
	RefreshToken   string `json:"refresh_token"`
	OpenId         string `json:"openid"`
	Scope          string `json:"scope"`
	IsSnapshotUser uint32 `json:"is_snapshotuser"`
	UnionId        string `json:"unionid"`
}

func UserCodeToAccessToken(_appId string, _secret string, _code string) (error, *WxResUserCodeToToken) {
	client := clHttpClient.NewClient("https://api.weixin.qq.com/sns/oauth2/access_token")
	client.AddParam("grant_type", "authorization_code")
	client.AddParam("appid", _appId)
	client.AddParam("secret", _secret)
	client.AddParam("code", _code)
	resp, err := client.Do()
	if err != nil {
		return err, nil
	}
	var respObj WxResUserCodeToToken
	if err := json.Unmarshal([]byte(resp.Body), &respObj); err != nil {
		return err, nil
	}
	if respObj.ErrCode != 0 {
		return errors.New(fmt.Sprintf("微信返回错误: %v - %v", respObj.ErrCode, respObj.ErrMsg)), nil
	}
	return nil, &respObj
}
