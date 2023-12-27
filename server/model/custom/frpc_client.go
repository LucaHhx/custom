// 自动生成模板FrpcClient
package custom

import (
	"custom-server/enum"
	"custom-server/global"
	"custom-server/utils/http"
	"fmt"
)

// frpcClient表 结构体  FrpcClient
type FrpcClient struct {
	global.GVA_MODEL
	User        string `json:"user" form:"user" gorm:"column:user;comment:用户;size:255;"`                         //用户
	ServerAddr  string `json:"serverAddr" form:"serverAddr" gorm:"column:server_addr;comment:服务器ip;size:255;"`   //服务器ip
	ServerPort  int    `json:"serverPort" form:"serverPort" gorm:"column:server_port;comment:服务器端口;size:32;"`    //服务器端口
	WebAddr     string `json:"webAddr" form:"webAddr" gorm:"column:web_addr;comment:网页ip;size:255;"`             //网页ip
	WebPort     int    `json:"webPort" form:"webPort" gorm:"column:web_port;comment:网页端口;size:32;"`              //网页端口
	WebMapPort  int    `json:"webMapPort" form:"webMapPort" gorm:"column:web_map_port;comment:网页映射端口;size:32;"`  //网页映射端口
	WebUser     string `json:"webUser" form:"webUser" gorm:"column:web_user;comment:网页用户;size:255;"`             //网页用户
	WebPassword string `json:"webPassword" form:"webPassword" gorm:"column:web_password;comment:网页密码;size:255;"` //网页密码
	Url         string `json:"url" form:"url" gorm:"column:url;comment:网页地址;size:255;"`                          //网页地址
	Config      string `json:"config" form:"config" gorm:"column:config;comment:额外配置;type:text;"`                //额外配置
}

// TableName frpcClient表 FrpcClient自定义表名 frpc_client
func (f *FrpcClient) TableName() string {
	return "frpc_client"
}

func (f *FrpcClient) GetUrl() string {
	if f.Url == "" {
		return fmt.Sprintf("http://%s:%d", f.ServerAddr, f.WebMapPort)
	} else {
		return fmt.Sprintf("https://%s", f.Url)
	}
}

func (f *FrpcClient) GetConfig() (str string, err error) {
	str += fmt.Sprintf("user = \"%s\"\n", f.User)
	str += fmt.Sprintf("serverAddr = \"%s\"\n", f.ServerAddr)
	str += fmt.Sprintf("serverPort = %d\n", f.ServerPort)
	str += fmt.Sprintf("webServer.addr = \"%s\"\n", f.WebAddr)
	str += fmt.Sprintf("webServer.port = %d\n", f.WebPort)
	str += fmt.Sprintf("webServer.user = \"%s\"\n", f.WebUser)
	str += fmt.Sprintf("webServer.password = \"%s\"\n", f.WebPassword)
	str += f.Config

	//[[proxies]]
	//	name = "Frpc-https"
	//	type = "https"
	//	localIP = "127.0.0.1"
	//	customDomains = ["frp.lucah.top"]
	//	[proxies.plugin]
	//	type = "https2http"
	//	localAddr = "127.0.0.1:7401"
	//	crtPath = "./certs/lucah_cert.pem"
	//	keyPath = "./certs/lucah_key.pem"
	//	hostHeaderRewrite = "127.0.0.1"
	//	requestHeaders.set.x-from-where = "frp"

	str += fmt.Sprintf("\n[[proxies]]\n")
	str += fmt.Sprintf("name = \"%s\"\n", "Frpc")
	str += fmt.Sprintf("type = \"%s\"\n", enum.FrpcTypeHttps)
	str += fmt.Sprintf("localIP = \"%s\"\n", enum.FrpcLocalIP)
	str += fmt.Sprintf("customDomains = [\"%s\"]\n", f.Url)
	str += fmt.Sprintf("[proxies.plugin]\n")
	str += fmt.Sprintf("type = \"%s\"\n", enum.FrpcPluginHttps2Http)
	str += fmt.Sprintf("localAddr = \"%s\"\n", fmt.Sprintf("%s:%d", f.WebAddr, f.WebPort))
	str += fmt.Sprintf("crtPath = \"%s\"\n", "./certs/lucah_cert.pem")
	str += fmt.Sprintf("keyPath = \"%s\"\n", "./certs/lucah_key.pem")
	str += fmt.Sprintf("hostHeaderRewrite = \"%s\"\n", enum.FrpcLocalIP)
	str += fmt.Sprintf("requestHeaders.set.x-from-where = \"%s\"\n", "frp")

	var list []FrpcClientConfig
	err = global.GVA_DB.Where("client_id = ?", f.ID).Find(&list).Error
	if err != nil {
		return
	}
	for _, v := range list {
		str += fmt.Sprintf("\n%s\n", v.GetConfig())
	}
	return
}

func (f *FrpcClient) UploadConfig() (ack *http.SendAck, err error) {
	var config string
	config, err = f.GetConfig()
	if err != nil {
		return
	}
	ack, err = http.GetHttpConn(
		http.SetMethod(enum.MethodPut),
		http.SetUrl(f.GetUrl()),
		http.SetPath("/api/config"),
		http.SetBasicAuth(f.WebUser, f.WebPassword),
		http.SetData([]byte(config)),
	).Send()
	return
}

func (f *FrpcClient) ReloadConfig() (ack *http.SendAck, err error) {
	ack, err = http.GetHttpConn(
		http.SetMethod(enum.MethodGet),
		http.SetUrl(f.GetUrl()),
		http.SetPath("/api/reload"),
		http.SetBasicAuth(f.WebUser, f.WebPassword),
	).Send()
	return
}
