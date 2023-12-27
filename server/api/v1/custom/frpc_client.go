package custom

import (
	"custom-server/enum"
	"custom-server/global"
	"custom-server/model/common/request"
	"custom-server/model/common/response"
	"custom-server/model/custom"
	customReq "custom-server/model/custom/request"
	"custom-server/service"
	"custom-server/utils"
	"custom-server/utils/http"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type FrpcClientApi struct {
}

var frpcClientService = service.ServiceGroupApp.CustomServiceGroup.FrpcClientService

// CreateFrpcClient 创建frpcClient表
// @Tags FrpcClient
// @Summary 创建frpcClient表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body custom.FrpcClient true "创建frpcClient表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /frpcClient/createFrpcClient [post]
func (frpcClientApi *FrpcClientApi) CreateFrpcClient(c *gin.Context) {
	var frpcClient custom.FrpcClient
	err := c.ShouldBindJSON(&frpcClient)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"User":       {utils.NotEmpty()},
		"ServerAddr": {utils.NotEmpty()},
		"ServerPort": {utils.NotEmpty()},
		"WebAddr":    {utils.NotEmpty()},
		"WebPort":    {utils.NotEmpty()},
		"WebMapPort": {utils.NotEmpty()},
	}
	if err := utils.Verify(frpcClient, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := frpcClientService.CreateFrpcClient(&frpcClient); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteFrpcClient 删除frpcClient表
// @Tags FrpcClient
// @Summary 删除frpcClient表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body custom.FrpcClient true "删除frpcClient表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /frpcClient/deleteFrpcClient [delete]
func (frpcClientApi *FrpcClientApi) DeleteFrpcClient(c *gin.Context) {
	var frpcClient custom.FrpcClient
	err := c.ShouldBindJSON(&frpcClient)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := frpcClientService.DeleteFrpcClient(frpcClient); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteFrpcClientByIds 批量删除frpcClient表
// @Tags FrpcClient
// @Summary 批量删除frpcClient表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除frpcClient表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /frpcClient/deleteFrpcClientByIds [delete]
func (frpcClientApi *FrpcClientApi) DeleteFrpcClientByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := frpcClientService.DeleteFrpcClientByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateFrpcClient 更新frpcClient表
// @Tags FrpcClient
// @Summary 更新frpcClient表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body custom.FrpcClient true "更新frpcClient表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /frpcClient/updateFrpcClient [put]
func (frpcClientApi *FrpcClientApi) UpdateFrpcClient(c *gin.Context) {
	var frpcClient custom.FrpcClient
	err := c.ShouldBindJSON(&frpcClient)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"User":       {utils.NotEmpty()},
		"ServerAddr": {utils.NotEmpty()},
		"ServerPort": {utils.NotEmpty()},
		"WebAddr":    {utils.NotEmpty()},
		"WebPort":    {utils.NotEmpty()},
		"WebMapPort": {utils.NotEmpty()},
	}
	if err := utils.Verify(frpcClient, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := frpcClientService.UpdateFrpcClient(frpcClient); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindFrpcClient 用id查询frpcClient表
// @Tags FrpcClient
// @Summary 用id查询frpcClient表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query custom.FrpcClient true "用id查询frpcClient表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /frpcClient/findFrpcClient [get]
func (frpcClientApi *FrpcClientApi) FindFrpcClient(c *gin.Context) {
	var frpcClient custom.FrpcClient
	err := c.ShouldBindQuery(&frpcClient)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if refrpcClient, err := frpcClientService.GetFrpcClient(frpcClient.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"refrpcClient": refrpcClient}, c)
	}
}

func (frpcClientApi *FrpcClientApi) UploadFrpcClient(c *gin.Context) {
	var frpcClient custom.FrpcClient
	err := c.ShouldBindQuery(&frpcClient)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if refrpcClient, err := frpcClientService.GetFrpcClient(frpcClient.ID); err != nil {
		global.GVA_LOG.Error("上传失败!", zap.Error(err))
		response.FailWithMessage("上传失败", c)
	} else {
		var config *http.SendAck
		config, err = refrpcClient.UploadConfig()
		if err != nil {
			response.FailWithMessage(err.Error(), c)
		}
		config, err = refrpcClient.ReloadConfig()
		if err != nil {
			response.FailWithMessage(config.Data, c)
		}
		response.OkWithData(config, c)
	}
}

// GetFrpcClientList 分页获取frpcClient表列表
// @Tags FrpcClient
// @Summary 分页获取frpcClient表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query customReq.FrpcClientSearch true "分页获取frpcClient表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /frpcClient/getFrpcClientList [get]
func (frpcClientApi *FrpcClientApi) GetFrpcClientList(c *gin.Context) {
	var pageInfo customReq.FrpcClientSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := frpcClientService.GetFrpcClientInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		if pageInfo.IsMap {
			clientMap := make(map[uint]string)
			for _, client := range list {
				clientMap[client.ID] = client.User
			}
			global.BlackCache.SetNoExpire(enum.FrpcClientMapKey, clientMap)
			response.OkWithDetailed(clientMap, "获取成功", c)
		} else {
			response.OkWithDetailed(response.PageResult{
				List:     list,
				Total:    total,
				Page:     pageInfo.Page,
				PageSize: pageInfo.PageSize,
			}, "获取成功", c)
		}
	}
}
