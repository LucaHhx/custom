package custom

import (
	"custom-server/global"
	"custom-server/model/common/request"
	"custom-server/model/common/response"
	"custom-server/model/custom"
	customReq "custom-server/model/custom/request"
	"custom-server/service"
	"custom-server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type FrpcClientConfigApi struct {
}

var frpcClientConfigService = service.ServiceGroupApp.CustomServiceGroup.FrpcClientConfigService

// CreateFrpcClientConfig 创建frpcClientConfig表
// @Tags FrpcClientConfig
// @Summary 创建frpcClientConfig表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body custom.FrpcClientConfig true "创建frpcClientConfig表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /frpcClientConfig/createFrpcClientConfig [post]
func (frpcClientConfigApi *FrpcClientConfigApi) CreateFrpcClientConfig(c *gin.Context) {
	var frpcClientConfig custom.FrpcClientConfig
	err := c.ShouldBindJSON(&frpcClientConfig)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"ClientId":   {utils.NotEmpty()},
		"ChannelTyp": {utils.NotEmpty()},
		"Name":       {utils.NotEmpty()},
	}
	if err := utils.Verify(frpcClientConfig, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := frpcClientConfigService.CreateFrpcClientConfig(&frpcClientConfig); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteFrpcClientConfig 删除frpcClientConfig表
// @Tags FrpcClientConfig
// @Summary 删除frpcClientConfig表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body custom.FrpcClientConfig true "删除frpcClientConfig表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /frpcClientConfig/deleteFrpcClientConfig [delete]
func (frpcClientConfigApi *FrpcClientConfigApi) DeleteFrpcClientConfig(c *gin.Context) {
	var frpcClientConfig custom.FrpcClientConfig
	err := c.ShouldBindJSON(&frpcClientConfig)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := frpcClientConfigService.DeleteFrpcClientConfig(frpcClientConfig); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteFrpcClientConfigByIds 批量删除frpcClientConfig表
// @Tags FrpcClientConfig
// @Summary 批量删除frpcClientConfig表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除frpcClientConfig表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /frpcClientConfig/deleteFrpcClientConfigByIds [delete]
func (frpcClientConfigApi *FrpcClientConfigApi) DeleteFrpcClientConfigByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := frpcClientConfigService.DeleteFrpcClientConfigByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateFrpcClientConfig 更新frpcClientConfig表
// @Tags FrpcClientConfig
// @Summary 更新frpcClientConfig表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body custom.FrpcClientConfig true "更新frpcClientConfig表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /frpcClientConfig/updateFrpcClientConfig [put]
func (frpcClientConfigApi *FrpcClientConfigApi) UpdateFrpcClientConfig(c *gin.Context) {
	var frpcClientConfig custom.FrpcClientConfig
	err := c.ShouldBindJSON(&frpcClientConfig)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"ClientId":   {utils.NotEmpty()},
		"ChannelTyp": {utils.NotEmpty()},
		"Name":       {utils.NotEmpty()},
	}
	if err := utils.Verify(frpcClientConfig, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := frpcClientConfigService.UpdateFrpcClientConfig(frpcClientConfig); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindFrpcClientConfig 用id查询frpcClientConfig表
// @Tags FrpcClientConfig
// @Summary 用id查询frpcClientConfig表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query custom.FrpcClientConfig true "用id查询frpcClientConfig表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /frpcClientConfig/findFrpcClientConfig [get]
func (frpcClientConfigApi *FrpcClientConfigApi) FindFrpcClientConfig(c *gin.Context) {
	var frpcClientConfig custom.FrpcClientConfig
	err := c.ShouldBindQuery(&frpcClientConfig)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if refrpcClientConfig, err := frpcClientConfigService.GetFrpcClientConfig(frpcClientConfig.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"refrpcClientConfig": refrpcClientConfig}, c)
	}
}

// GetFrpcClientConfigList 分页获取frpcClientConfig表列表
// @Tags FrpcClientConfig
// @Summary 分页获取frpcClientConfig表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query customReq.FrpcClientConfigSearch true "分页获取frpcClientConfig表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /frpcClientConfig/getFrpcClientConfigList [get]
func (frpcClientConfigApi *FrpcClientConfigApi) GetFrpcClientConfigList(c *gin.Context) {
	var pageInfo customReq.FrpcClientConfigSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := frpcClientConfigService.GetFrpcClientConfigInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
