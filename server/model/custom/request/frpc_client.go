package request

import (
	"custom-server/model/common/request"
	"custom-server/model/custom"
	"time"
)

type FrpcClientSearch struct {
	custom.FrpcClient
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	request.PageInfo
	IsMap bool `json:"isMap" form:"isMap"`
}
