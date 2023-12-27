import service from '@/utils/request'

// @Tags FrpcClientConfig
// @Summary 创建frpcClientConfig表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FrpcClientConfig true "创建frpcClientConfig表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /frpcClientConfig/createFrpcClientConfig [post]
export const createFrpcClientConfig = (data) => {
  return service({
    url: '/frpcClientConfig/createFrpcClientConfig',
    method: 'post',
    data
  })
}

// @Tags FrpcClientConfig
// @Summary 删除frpcClientConfig表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FrpcClientConfig true "删除frpcClientConfig表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /frpcClientConfig/deleteFrpcClientConfig [delete]
export const deleteFrpcClientConfig = (data) => {
  return service({
    url: '/frpcClientConfig/deleteFrpcClientConfig',
    method: 'delete',
    data
  })
}

// @Tags FrpcClientConfig
// @Summary 批量删除frpcClientConfig表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除frpcClientConfig表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /frpcClientConfig/deleteFrpcClientConfig [delete]
export const deleteFrpcClientConfigByIds = (data) => {
  return service({
    url: '/frpcClientConfig/deleteFrpcClientConfigByIds',
    method: 'delete',
    data
  })
}

// @Tags FrpcClientConfig
// @Summary 更新frpcClientConfig表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FrpcClientConfig true "更新frpcClientConfig表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /frpcClientConfig/updateFrpcClientConfig [put]
export const updateFrpcClientConfig = (data) => {
  return service({
    url: '/frpcClientConfig/updateFrpcClientConfig',
    method: 'put',
    data
  })
}

// @Tags FrpcClientConfig
// @Summary 用id查询frpcClientConfig表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.FrpcClientConfig true "用id查询frpcClientConfig表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /frpcClientConfig/findFrpcClientConfig [get]
export const findFrpcClientConfig = (params) => {
  return service({
    url: '/frpcClientConfig/findFrpcClientConfig',
    method: 'get',
    params
  })
}

// @Tags FrpcClientConfig
// @Summary 分页获取frpcClientConfig表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取frpcClientConfig表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /frpcClientConfig/getFrpcClientConfigList [get]
export const getFrpcClientConfigList = (params) => {
  return service({
    url: '/frpcClientConfig/getFrpcClientConfigList',
    method: 'get',
    params
  })
}
