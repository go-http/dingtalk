package dingtalk

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// API地址
const (
	USER_SIMPLELIST_URL = DINGTALK_API_URL + "/user/simplelist"
)

// 用户
type User struct {
	Userid string
	Name   string
}

// 获取部门列表API的返回数据
type UserSimplelistResponse struct {
	CommonResponse
	HasMore  bool
	UserList []User
}

// 获取指定部门的子部门列表
func GetUsers(deptId int) ([]User, error) {
	param := url.Values{
		"access_token":  {accessToken},
		"department_id": {fmt.Sprintf("%d", deptId)},
	}

	resp, err := http.Get(USER_SIMPLELIST_URL + "?" + param.Encode())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var info UserSimplelistResponse
	err = json.NewDecoder(resp.Body).Decode(&info)
	if err != nil {
		return nil, err
	}

	if err := info.CommonResponse.Error(); err != nil {
		return nil, err
	}

	return info.UserList, nil
}
