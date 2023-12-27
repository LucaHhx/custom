package request

import (
	"custom-server/model/common/request"
	"custom-server/model/system"
)

type SysOperationRecordSearch struct {
	system.SysOperationRecord
	request.PageInfo
}
