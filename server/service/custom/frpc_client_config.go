package custom

import (
	"custom-server/enum"
	"custom-server/global"
	"custom-server/model/common/request"
	"custom-server/model/custom"
	customReq "custom-server/model/custom/request"
	"custom-server/utils/helper"
	"custom-server/utils/http"
	"fmt"
	"github.com/goccy/go-json"
	"time"
)

type FrpcClientConfigService struct {
}

// CreateFrpcClientConfig 创建frpcClientConfig表记录
// Author [piexlmax](https://github.com/piexlmax)
func (frpcClientConfigService *FrpcClientConfigService) CreateFrpcClientConfig(frpcClientConfig *custom.FrpcClientConfig) (err error) {
	frpcClientConfig.CreatedAt = time.Now()
	err = global.GVA_DB.Create(frpcClientConfig).Error
	return err
}

// DeleteFrpcClientConfig 删除frpcClientConfig表记录
// Author [piexlmax](https://github.com/piexlmax)
func (frpcClientConfigService *FrpcClientConfigService) DeleteFrpcClientConfig(frpcClientConfig custom.FrpcClientConfig) (err error) {
	err = global.GVA_DB.Delete(&frpcClientConfig).Error
	return err
}

// DeleteFrpcClientConfigByIds 批量删除frpcClientConfig表记录
// Author [piexlmax](https://github.com/piexlmax)
func (frpcClientConfigService *FrpcClientConfigService) DeleteFrpcClientConfigByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]custom.FrpcClientConfig{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateFrpcClientConfig 更新frpcClientConfig表记录
// Author [piexlmax](https://github.com/piexlmax)
func (frpcClientConfigService *FrpcClientConfigService) UpdateFrpcClientConfig(frpcClientConfig custom.FrpcClientConfig) (err error) {
	err = global.GVA_DB.Save(&frpcClientConfig).Error
	return err
}

// GetFrpcClientConfig 根据id获取frpcClientConfig表记录
// Author [piexlmax](https://github.com/piexlmax)
func (frpcClientConfigService *FrpcClientConfigService) GetFrpcClientConfig(id uint) (frpcClientConfig custom.FrpcClientConfig, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&frpcClientConfig).Error
	return
}

// GetFrpcClientConfigInfoList 分页获取frpcClientConfig表记录
// Author [piexlmax](https://github.com/piexlmax)
func (frpcClientConfigService *FrpcClientConfigService) GetFrpcClientConfigInfoList(info customReq.FrpcClientConfigSearch) (list []custom.FrpcClientConfig, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&custom.FrpcClientConfig{})
	var frpcClientConfigs []custom.FrpcClientConfig
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.ClientId != 0 {
		db = db.Where("client_id = ?", info.ClientId)
	}
	if info.ChannelTyp != 0 {
		db = db.Where("channel_typ = ?", info.ChannelTyp)
	}
	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	if info.LocalIp != "" {
		db = db.Where("local_ip LIKE ?", "%"+info.LocalIp+"%")
	}
	if info.LocalPort != 0 {
		db = db.Where("local_port = ?", info.LocalPort)
	}
	if info.Status != 0 {
		db = db.Where("status = ?", info.Status)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&frpcClientConfigs).Error
	var status map[string]int
	status, err = frpcClientConfigService.GetFrpcClientConfigInfo(
		helper.ArrayDistinct(frpcClientConfigs, func(f custom.FrpcClientConfig) int {
			return f.ClientId
		})...,
	)
	if err != nil {
		return nil, 0, err
	}

	value, ok := global.BlackCache.Get(enum.FrpcClientMapKey)
	if !ok {
		return frpcClientConfigs, total, err
	}

	for i, v := range frpcClientConfigs {
		v.Status = status[fmt.Sprintf("%s-%s", enum.FrpcTypeMap[v.ChannelTyp], fmt.Sprintf("%s.%s", value.(map[uint]string)[uint(v.ClientId)], v.Name))]
		frpcClientConfigs[i] = v
	}

	return frpcClientConfigs, total, err
}

func (frpcClientConfigService *FrpcClientConfigService) GetFrpcClientConfigInfo(ids ...int) (status map[string]int, err error) {
	var ack *http.SendAck
	var tcpList []custom.FrpcItem
	for _, id := range ids {
		var frpcClient custom.FrpcClient
		err = global.GVA_DB.Model(&custom.FrpcClient{}).Where("id = ?", id).First(&frpcClient).Error
		if err != nil {
			return
		}
		ack, err = http.GetHttpConn(
			http.SetMethod(enum.MethodGet),
			http.SetUrl(frpcClient.GetUrl()),
			http.SetPath("/api/status"),
			http.SetBasicAuth(frpcClient.WebUser, frpcClient.WebPassword),
		).Send()
		if err != nil {
			return
		}
		info := custom.FrpcClientConfigInfo{}
		err = json.Unmarshal([]byte(ack.Data), &info)
		if err != nil {
			return
		}
		tcpList = append(tcpList, info.Tcp...)
		tcpList = append(tcpList, info.Https...)
	}
	status = helper.ArrayToMap(tcpList, func(t custom.FrpcItem) (string, int) {
		return fmt.Sprintf("%s-%s", t.Type, t.Name), enum.FrpcStatusMap[t.Status]
	})
	return

}
