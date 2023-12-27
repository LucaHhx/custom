package request

import (
	"custom-server/model/common/request"
	"custom-server/model/system"
)

type ChatGptRequest struct {
	system.ChatGpt
	request.PageInfo
}
