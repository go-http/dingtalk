package dingtalk

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// API地址
const (
	DEPARTMENT_LIST_URL = DINGTALK_API_URL + "/department/list"
)

// 部门
type Department struct {
	Id              int
	Name            string
	ParentId        int
	CreateDeptGroup bool
	AutoAddUser     bool
	Ext             string
}

// 获取部门列表API的返回数据
type DepartmentListResponse struct {
	CommonResponse
	Department []Department
}

// 获取指定部门的子部门列表
func GetDepartmentList(id int, fetchChild bool) ([]Department, error) {
	param := url.Values{
		"access_token": {accessToken},
		"fetch_child":  {fmt.Sprintf("%t", fetchChild)},
	}

	if id != 0 {
		param.Set("id", fmt.Sprintf("%d", id))

	}

	resp, err := http.Get(DEPARTMENT_LIST_URL + "?" + param.Encode())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var info DepartmentListResponse
	err = json.NewDecoder(resp.Body).Decode(&info)
	if err != nil {
		return nil, err
	}

	if err := info.CommonResponse.Error(); err != nil {
		return nil, err
	}

	return info.Department, nil
}

// 获取所有部门列表
func GetAllDepartments() ([]Department, error) {
	return GetDepartmentList(0, true)
}

// 获取重复（部门名称）的部门信息
func DumpDuplicatedDepartments() (map[string]int, error) {
	deptCount := map[string]int{}

	depts, err := GetDepartmentList(0, true)
	if err != nil {
		return nil, err
	}

	for _, dept := range depts {
		deptCount[dept.Name] += 1
	}

	for dept, count := range deptCount {
		if count == 1 {
			delete(deptCount, dept)
		}
	}

	return deptCount, nil
}

// 获取部门名称-ID映射表
func GetDepartmentIdMap() (map[string]int, error) {
	// 获取所有部门信息
	depts, err := GetDepartmentList(0, true)
	if err != nil {
		return nil, err
	}

	//检查是否有重名部门
	deptCount := map[string]int{}
	nameIdMap := map[string]int{}
	for _, dept := range depts {
		deptCount[dept.Name] += 1
		nameIdMap[dept.Name] = dept.Id
	}

	for dept, count := range deptCount {
		if count == 1 {
			delete(deptCount, dept)
		}
	}

	if len(deptCount) > 0 {
		return nil, fmt.Errorf("部分部门重名，重名部门包括%s", deptCount)
	}

	return nameIdMap, nil
}
