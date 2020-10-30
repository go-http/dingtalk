package dingtalk

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// API地址
const (
	REPORT_LIST_URL = DINGTALK_API_URL + "/topapi/report/list"
)

// 日志内容
type ReportContent struct {
	Sort  string
	Type  string
	Value string
	Key   string
}

// 日志
type Report struct {
	Remark       string
	Contents     []ReportContent
	TemplateName string `json:"template_name"`
	DeptName     string `json:"dept_name"`
	CreatorName  string `json:"creator_name"`
	CreatorId    string `json:"creator_id"`
	CreateTime   int    `json:"create_time"`
	ReportId     string `json:"report_id"`
}

func (r *Report) DeptNames() []string {
	return strings.Split(r.DeptName, ",")
}

const DAILY_REPORT = "日报"

// 获取指定时间段的日报
func GetDailyReports(start, end time.Time) ([]Report, error) {
	return GetReports(start, end, DAILY_REPORT)
}

// 获取指定时间段的所有日志
func GetAllReports(start, end time.Time) ([]Report, error) {
	return GetReports(start, end, "")
}

// 获取指定时间段、指定模板的日志
func GetReports(start, end time.Time, template string) ([]Report, error) {
	reports := make([]Report, 0, 300)
	cursor := 0
	hasMore := true

	for hasMore {
		resp, err := GetUserReportFromCursor(start, end, template, "", cursor)
		if err != nil {
			return nil, err
		}

		reports = append(reports, resp.Result.DataList...)

		hasMore = resp.Result.HasMore
		cursor = resp.Result.NextCursor
	}

	return reports, nil
}

// 获取日志API的返回数据
type ReportListResponse struct {
	CommonResponse
	Result struct {
		Size       int
		NextCursor int      `json:"next_cursor"`
		HasMore    bool     `json:"has_more"`
		DataList   []Report `json:"data_list"`
	}
}

// 获取指定起始点的日志
func GetUserReportFromCursor(start, end time.Time, template, userid string, cursor int) (*ReportListResponse, error) {
	param := url.Values{
		"access_token": {accessToken},
		"start_time":   {fmt.Sprintf("%d", start.Unix()*1000)},
		"end_time":     {fmt.Sprintf("%d", end.Unix()*1000)},
		"cursor":       {fmt.Sprintf("%d", cursor)},
		"size":         {"20"}, //直接用钉钉允许的最大值
	}

	if template != "" {
		param.Set("template_name", template)
	}

	if userid != "" {
		param.Set("userid", userid)
	}

	resp, err := http.Get(REPORT_LIST_URL + "?" + param.Encode())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var info ReportListResponse
	err = json.NewDecoder(resp.Body).Decode(&info)
	if err != nil {
		return nil, err
	}

	if err := info.CommonResponse.Error(); err != nil {
		return nil, err
	}

	return &info, nil
}
