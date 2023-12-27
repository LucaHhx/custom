import service from '@/utils/request'

// @Tags FrpcClient
// @Summary 创建frpcClient表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FrpcClient true "创建frpcClient表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /frpcClient/createFrpcClient [post]
export const createFrpcClient = (data) => {
  return service({
    url: '/frpcClient/createFrpcClient',
    method: 'post',
    data
  })
}

// @Tags FrpcClient
// @Summary 删除frpcClient表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FrpcClient true "删除frpcClient表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /frpcClient/deleteFrpcClient [delete]
export const deleteFrpcClient = (data) => {
  return service({
    url: '/frpcClient/deleteFrpcClient',
    method: 'delete',
    data
  })
}

// @Tags FrpcClient
// @Summary 批量删除frpcClient表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除frpcClient表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /frpcClient/deleteFrpcClient [delete]
export const deleteFrpcClientByIds = (data) => {
  return service({
    url: '/frpcClient/deleteFrpcClientByIds',
    method: 'delete',
    data
  })
}

// @Tags FrpcClient
// @Summary 更新frpcClient表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FrpcClient true "更新frpcClient表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /frpcClient/updateFrpcClient [put]
export const updateFrpcClient = (data) => {
  return service({
    url: '/frpcClient/updateFrpcClient',
    method: 'put',
    data
  })
}

// @Tags FrpcClient
// @Summary 用id查询frpcClient表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.FrpcClient true "用id查询frpcClient表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /frpcClient/findFrpcClient [get]
export const findFrpcClient = (params) => {
  return service({
    url: '/frpcClient/findFrpcClient',
    method: 'get',
    params
  })
}

export const uploadFrpcClient = (params) => {
  return service({
    url: '/frpcClient/uploadFrpcClient',
    method: 'get',
    params
  })
}

// @Tags FrpcClient
// @Summary 分页获取frpcClient表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取frpcClient表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /frpcClient/getFrpcClientList [get]
export const getFrpcClientList = (params) => {
  return service({
    url: '/frpcClient/getFrpcClientList',
    method: 'get',
    params
  })
}
