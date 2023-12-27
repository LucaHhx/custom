package enum

const (
	FrpcTypeTcp   = "tcp"
	FrpcTypeXTcp  = "http"
	FrpcTypeUdp   = "udp"
	FrpcTypeHttps = "https"
	FrpcLocalIP   = "127.0.0.1"
)
const (
	FrpcPluginHttps2Http = "https2http"
)

var FrpcTypeMap = map[int]string{
	1: FrpcTypeTcp,
	2: FrpcTypeXTcp,
	3: FrpcTypeHttps,
}

const (
	FrpcStatusStop  = 0
	FrpcStatusRun   = 1
	FrpcStatusClose = 9
)

var FrpcStatusMap = map[string]int{
	"":        FrpcStatusStop,
	"running": FrpcStatusRun,
	"close":   FrpcStatusClose,
}

const FrpcClientMapKey = "frpc_client_map_key"
