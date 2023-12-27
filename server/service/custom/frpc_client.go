package custom

import (
	"custom-server/global"
	"custom-server/model/common/request"
	"custom-server/model/custom"
	customReq "custom-server/model/custom/request"
)

type FrpcClientService struct {
}

// CreateFrpcClient 创建frpcClient表记录
// Author [piexlmax](https://github.com/piexlmax)
func (frpcClientService *FrpcClientService) CreateFrpcClient(frpcClient *custom.FrpcClient) (err error) {
	frpcClient.Url = frpcClient.GetUrl()
	err = global.GVA_DB.Create(frpcClient).Error
	return err
}

// DeleteFrpcClient 删除frpcClient表记录
// Author [piexlmax](https://github.com/piexlmax)
func (frpcClientService *FrpcClientService) DeleteFrpcClient(frpcClient custom.FrpcClient) (err error) {
	err = global.GVA_DB.Delete(&frpcClient).Error
	return err
}

// DeleteFrpcClientByIds 批量删除frpcClient表记录
// Author [piexlmax](https://github.com/piexlmax)
func (frpcClientService *FrpcClientService) DeleteFrpcClientByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]custom.FrpcClient{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateFrpcClient 更新frpcClient表记录
// Author [piexlmax](https://github.com/piexlmax)
func (frpcClientService *FrpcClientService) UpdateFrpcClient(frpcClient custom.FrpcClient) (err error) {
	frpcClient.Url = frpcClient.GetUrl()
	err = global.GVA_DB.Save(&frpcClient).Error
	return err
}

// GetFrpcClient 根据id获取frpcClient表记录
// Author [piexlmax](https://github.com/piexlmax)
func (frpcClientService *FrpcClientService) GetFrpcClient(id uint) (frpcClient custom.FrpcClient, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&frpcClient).Error
	frpcClient.GetUrl()
	return
}

// GetFrpcClientInfoList 分页获取frpcClient表记录
// Author [piexlmax](https://github.com/piexlmax)
func (frpcClientService *FrpcClientService) GetFrpcClientInfoList(info customReq.FrpcClientSearch) (list []custom.FrpcClient, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&custom.FrpcClient{})
	var frpcClients []custom.FrpcClient
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.User != "" {
		db = db.Where("user LIKE ?", "%"+info.User+"%")
	}
	if info.ServerAddr != "" {
		db = db.Where("server_addr LIKE ?", "%"+info.ServerAddr+"%")
	}
	if info.WebAddr != "" {
		db = db.Where("web_addr LIKE ?", "%"+info.WebAddr+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&frpcClients).Error
	for i, client := range frpcClients {
		frpcClients[i].Url = client.GetUrl()
	}
	return frpcClients, total, err
}
