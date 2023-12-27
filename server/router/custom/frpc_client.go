package custom

import (
	"custom-server/api/v1"
	"custom-server/middleware"
	"github.com/gin-gonic/gin"
)

type FrpcClientRouter struct {
}

// InitFrpcClientRouter 初始化 frpcClient表 路由信息
func (s *FrpcClientRouter) InitFrpcClientRouter(Router *gin.RouterGroup) {
	frpcClientRouter := Router.Group("frpcClient").Use(middleware.OperationRecord())
	frpcClientRouterWithoutRecord := Router.Group("frpcClient")
	var frpcClientApi = v1.ApiGroupApp.CustomApiGroup.FrpcClientApi
	{
		frpcClientRouter.POST("createFrpcClient", frpcClientApi.CreateFrpcClient)             // 新建frpcClient表
		frpcClientRouter.DELETE("deleteFrpcClient", frpcClientApi.DeleteFrpcClient)           // 删除frpcClient表
		frpcClientRouter.DELETE("deleteFrpcClientByIds", frpcClientApi.DeleteFrpcClientByIds) // 批量删除frpcClient表
		frpcClientRouter.PUT("updateFrpcClient", frpcClientApi.UpdateFrpcClient)              // 更新frpcClient表
	}
	{
		frpcClientRouterWithoutRecord.GET("findFrpcClient", frpcClientApi.FindFrpcClient)       // 根据ID获取frpcClient表
		frpcClientRouterWithoutRecord.GET("getFrpcClientList", frpcClientApi.GetFrpcClientList) // 获取frpcClient表列表
		frpcClientRouterWithoutRecord.GET("uploadFrpcClient", frpcClientApi.UploadFrpcClient)   // 获取frpcClient表列表

	}
}
