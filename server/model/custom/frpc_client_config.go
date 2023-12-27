// 自动生成模板FrpcClientConfig
package custom

import (
	"custom-server/enum"
	"custom-server/global"
	"fmt"
)

// frpcClientConfig表 结构体  FrpcClientConfig
type FrpcClientConfig struct {
	global.GVA_MODEL
	ClientId   int    `json:"clientId" form:"clientId" gorm:"column:client_id;comment:客户端编号;size:32;"`      //客户端编号
	ChannelTyp int    `json:"channelTyp" form:"channelTyp" gorm:"column:channel_typ;comment:通道类型;size:32;"` //通道类型
	Name       string `json:"name" form:"name" gorm:"column:name;comment:配置名称;size:255;"`                   //配置名称
	LocalIp    string `json:"localIp" form:"localIp" gorm:"column:local_ip;comment:本地ip;size:255;"`         //本地ip
	LocalPort  int    `json:"localPort" form:"localPort" gorm:"column:local_port;comment:本地端口;size:32;"`    //本地端口
	RemotePort int    `json:"remotePort" form:"remotePort" gorm:"column:remote_port;comment:远程端口;size:32;"` //远程端口
	Status     int    `json:"status" form:"status" gorm:"column:status;comment:状态;size:8;"`                 //状态
	Config     string `json:"config" form:"config" gorm:"column:config;comment:其他配置;type:text;"`            //其他配置
	Url        string `json:"url" form:"url" gorm:"column:url;comment:网页地址;size:255;"`                      //网页地址
}

// TableName frpcClientConfig表 FrpcClientConfig自定义表名 frpc_client_config
func (t *FrpcClientConfig) TableName() string {
	return "frpc_client_config"
}

func (t *FrpcClientConfig) GetConfig() (str string) {
	//[[proxies]]
	//name = "frpc"
	//type = "tcp"
	//localIP = "127.0.0.1"
	//localPort = 7401
	//remotePort = 47401

	str += fmt.Sprintf("[[proxies]]\n")
	str += fmt.Sprintf("name = \"%s\"\n", t.Name)
	str += fmt.Sprintf("type = \"%s\"\n", enum.FrpcTypeMap[t.ChannelTyp])
	if t.LocalIp != "" {
		str += fmt.Sprintf("localIP = \"%s\"\n", t.LocalIp)
	}
	if t.LocalPort != 0 {
		str += fmt.Sprintf("localPort = %d\n", t.LocalPort)
	}
	if t.RemotePort != 0 {
		str += fmt.Sprintf("remotePort = %d\n", t.RemotePort)
	}
	if t.Url != "" {
		str += fmt.Sprintf("customDomains = [\"%s\"]\n", t.Url)
	}
	str += t.Config
	return
}

type FrpcItem struct {
	Name       string `json:"name"`
	Type       string `json:"type"`
	Status     string `json:"status"`
	Err        string `json:"err"`
	LocalAddr  string `json:"local_addr"`
	Plugin     string `json:"plugin"`
	RemoteAddr string `json:"remote_addr"`
}

type FrpcClientConfigInfo struct {
	Tcp   []FrpcItem `json:"tcp"`
	Https []FrpcItem `json:"https"`
}
