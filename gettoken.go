package dingtalk

import (
	"encoding/json"
	"net/http"
	"net/url"
)

// API地址
const (
	GETTOKEN_URL = DINGTALK_API_URL + "/gettoken"
)

var accessToken string

// 刷新AccessToken
func RefreshAccessToken(appkey, appsecret string) error {
	param := url.Values{
		"appkey":    {appkey},
		"appsecret": {appsecret},
	}

	resp, err := http.Get(GETTOKEN_URL + "?" + param.Encode())
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var info struct {
		CommonResponse
		AccessToken string `json:"access_token"`
	}

	err = json.NewDecoder(resp.Body).Decode(&info)
	if err != nil {
		return err
	}

	if err := info.CommonResponse.Error(); err != nil {
		return err
	}

	accessToken = info.AccessToken
	return nil
}
