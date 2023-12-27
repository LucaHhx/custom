package custom

import (
	"custom-server/api/v1"
	"custom-server/middleware"
	"github.com/gin-gonic/gin"
)

type FrpcClientConfigRouter struct {
}

// InitFrpcClientConfigRouter 初始化 frpcClientConfig表 路由信息
func (s *FrpcClientConfigRouter) InitFrpcClientConfigRouter(Router *gin.RouterGroup) {
	frpcClientConfigRouter := Router.Group("frpcClientConfig").Use(middleware.OperationRecord())
	frpcClientConfigRouterWithoutRecord := Router.Group("frpcClientConfig")
	var frpcClientConfigApi = v1.ApiGroupApp.CustomApiGroup.FrpcClientConfigApi
	{
		frpcClientConfigRouter.POST("createFrpcClientConfig", frpcClientConfigApi.CreateFrpcClientConfig)             // 新建frpcClientConfig表
		frpcClientConfigRouter.DELETE("deleteFrpcClientConfig", frpcClientConfigApi.DeleteFrpcClientConfig)           // 删除frpcClientConfig表
		frpcClientConfigRouter.DELETE("deleteFrpcClientConfigByIds", frpcClientConfigApi.DeleteFrpcClientConfigByIds) // 批量删除frpcClientConfig表
		frpcClientConfigRouter.PUT("updateFrpcClientConfig", frpcClientConfigApi.UpdateFrpcClientConfig)              // 更新frpcClientConfig表
	}
	{
		frpcClientConfigRouterWithoutRecord.GET("findFrpcClientConfig", frpcClientConfigApi.FindFrpcClientConfig)       // 根据ID获取frpcClientConfig表
		frpcClientConfigRouterWithoutRecord.GET("getFrpcClientConfigList", frpcClientConfigApi.GetFrpcClientConfigList) // 获取frpcClientConfig表列表
	}
}
