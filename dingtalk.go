package dingtalk

import (
	"fmt"
)

// 钉钉API地址
const (
	DINGTALK_API_URL = "https://oapi.dingtalk.com"
)

// 钉钉API公共响应
type CommonResponse struct {
	ErrCode int
	ErrMsg  string
}

func (resp *CommonResponse) Error() error {
	if resp.ErrCode != 0 || (resp.ErrMsg != "ok" && resp.ErrMsg != "") {
		return fmt.Errorf("[%d]%s", resp.ErrCode, resp.ErrMsg)
	}

	return nil
}
