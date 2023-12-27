package response

import "custom-server/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
