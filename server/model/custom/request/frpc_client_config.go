package request

import (
	"custom-server/model/common/request"
	"custom-server/model/custom"
	"time"
)

type FrpcClientConfigSearch struct {
	custom.FrpcClientConfig
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	request.PageInfo
}
