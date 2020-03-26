package dingtalk

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// API地址
const (
	USER_SIMPLELIST_URL = DINGTALK_API_URL + "/user/simplelist"
	USER_GET_URL        = DINGTALK_API_URL + "/user/get"
)

type UserRole struct {
	Id        int
	Name      string
	GroupName string
}

// 用户
type User struct {
	Unionid         string
	Remark          string
	Userid          string
	IsLeaderInDepts string
	IsBoss          bool
	HiredDate       int
	IsSenior        bool
	Tel             string
	Department      []int
	WorkPlace       string
	Email           string
	OrderInDepts    string
	Mobile          string
	Errmsg          string
	Active          bool
	Avatar          string
	IsAdmin         bool
	IsHide          bool
	Jobnumber       string
	Name            string
	Extattr         interface{}
	StateCode       string
	Position        string
	Roles           []UserRole
}

func (u *User) DeptLeaderInfo() map[int]bool {
	fragments := strings.Split(u.IsLeaderInDepts[1:len(u.IsLeaderInDepts)-1], ",")

	info := make(map[int]bool)
	for _, fragment := range fragments {
		kv := strings.Split(fragment, ":")
		id, _ := strconv.Atoi(kv[0])
		info[id] = kv[1] == "true"
	}
	json.Unmarshal([]byte(u.IsLeaderInDepts), &info)
	return info
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

// 获取用户详情的API的返回数据
type UserGetResponse struct {
	CommonResponse
	User
}

// 获取指定用户详情
func GetUser(userId string) (User, error) {
	param := url.Values{
		"access_token": {accessToken},
		"userid":       {userId},
	}

	resp, err := http.Get(USER_GET_URL + "?" + param.Encode())
	if err != nil {
		return User{}, err
	}
	defer resp.Body.Close()

	var info UserGetResponse
	err = json.NewDecoder(resp.Body).Decode(&info)
	if err != nil {
		return User{}, err
	}

	if err := info.CommonResponse.Error(); err != nil {
		return User{}, err
	}

	return info.User, nil
}
