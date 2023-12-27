package request

import (
	"custom-server/model/common/request"
	"custom-server/model/system"
)

type SysDictionaryDetailSearch struct {
	system.SysDictionaryDetail
	request.PageInfo
}
