package tests

import (
	"custom-server/core"
	"custom-server/global"
	"custom-server/model/custom"
	"fmt"
	"testing"
)

func TestPut(t *testing.T) {
	core.Init()
	frpcClient := custom.FrpcClient{}
	err := global.GVA_DB.Where("id = ?", 1).First(&frpcClient).Error
	if err != nil {
		fmt.Println("Error getting frpcClient:", err)
		return
	}
	s, err := frpcClient.UploadConfig()
	if err != nil {
		return
	}
	fmt.Printf("Code: %d, Ok: %t, Data: %s", s.Code, s.Ok, s.Data)

	config, err := frpcClient.ReloadConfig()
	if err != nil {
		return
	}
	fmt.Printf("Code: %d, Ok: %t, Data: %s", config.Code, config.Ok, config.Data)
	//data := []byte(frpcClient.GetConfig()) // 替换为你的数据
	//
	//req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/api/config", frpcClient.GetUrl()), bytes.NewBuffer(data))
	//if err != nil {
	//	fmt.Println("Error creating request:", err)
	//	return
	//}
	//
	//req.Header.Set("Content-Type", "application/json")
	//req.SetBasicAuth(frpcClient.WebUser, frpcClient.WebPassword)
	//
	//client := &http.Client{}
	//resp, err := client.Do(req)
	//if err != nil {
	//	fmt.Println("Error sending request:", err)
	//	return
	//}
	//defer resp.Body.Close()
	//
	//fmt.Println("Response status:", resp.Status)
}
