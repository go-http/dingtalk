package dingtalk

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"net/http"
)

// 钉钉企业内部机器人请求结构
type RobotRequest struct {
	MsgType string
	Text    struct {
		Content string
	}

	MsgId    string
	CreateAt int

	ConversationType  string
	ConversationId    string
	ConversationTitle string
	SenderId          string
	SenderNick        string
	SenderCorpId      string
	SenderStaffId     string
	ChatbotUserId     string
	AtUsers           []struct {
		DingtalkId string
		StaffId    string
	}
}

func IsValidateRobotRequest(appSecret string, r *http.Request) bool {
	s := r.Header.Get("timestamp") + "\n" + appSecret

	mac := hmac.New(sha256.New, []byte(appSecret))
	mac.Write([]byte(s))
	expectMac := mac.Sum(nil)

	sign := base64.StdEncoding.EncodeToString(expectMac)
	return sign == r.Header.Get("sign")
}
