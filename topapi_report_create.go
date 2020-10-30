package dingtalk

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
)

const (
	REPORT_CREATE_URL = DINGTALK_API_URL + "/topapi/report/create"
)

//填报的日志内容
type ReportContentToSend struct {
	Content     string `json:"content"`
	Key         string `json:"key"`
	Sort        int    `json:"sort"`
	Type        int    `json:"type"`
	ContentType string `json:"content_type"`
}

//CreateReport 创建日志，返回日志ID
func CreateReport(fromUserId, templateId string, contents []ReportContentToSend, toUserIds, toGroupIds []string) (string, error) {
	param := url.Values{
		"access_token": {accessToken},
	}

	var request struct {
		CreateReportParam struct {
			DdFrom     string                `json:"dd_from"`
			UserId     string                `json:"userid"`
			TemplateId string                `json:"template_id"`
			ToChat     bool                  `json:"to_chat"`
			ToUserIds  []string              `json:"to_userids"`
			ToCids     []string              `json:"to_cids"`
			Contents   []ReportContentToSend `json:"contents"`
		} `json:"create_report_param"`
	}

	request.CreateReportParam.UserId = fromUserId
	request.CreateReportParam.TemplateId = templateId
	request.CreateReportParam.ToUserIds = toUserIds
	request.CreateReportParam.ToCids = toGroupIds
	request.CreateReportParam.Contents = contents

	b, err := json.Marshal(request)
	if err != nil {
		return "", err
	}

	resp, err := http.Post(REPORT_CREATE_URL+"?"+param.Encode(), "application/json", bytes.NewReader(b))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var respInfo struct {
		Result string
		CommonResponse
	}
	err = json.NewDecoder(resp.Body).Decode(&respInfo)
	if err != nil {
		return "", err
	}

	if err := respInfo.CommonResponse.Error(); err != nil {
		return "", err
	}

	return respInfo.Result, nil
}
