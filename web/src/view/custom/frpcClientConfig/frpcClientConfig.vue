<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" :rules="searchRule" @keyup.enter="onSubmit">
        <el-form-item label="客户端编号" prop="clientId">
          <el-select v-model="searchInfo.clientId" :clearable="true" filterable placeholder="Please select">
            <el-option v-for="v,k in clientIdOpts" :key="k" :label="v" :value="parseInt(k)" />
          </el-select>

        </el-form-item>
        <el-form-item label="通道类型" prop="channelTyp">
          <el-select v-model="searchInfo.channelTyp" :clearable="true" filterable placeholder="Please select">
            <el-option v-for="v,k in channelTypOpts" :key="k" :label="v" :value="parseInt(k)" />
          </el-select>

        </el-form-item>
        <el-form-item label="配置名称" prop="name">
          <el-input v-model="searchInfo.name" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="本地ip" prop="localIp">
          <el-input v-model="searchInfo.localIp" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="本地端口" prop="localPort">

          <el-input v-model.number="searchInfo.localPort" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="状态" prop="status">
          <el-select v-model="searchInfo.status" :clearable="true" filterable placeholder="Please select">
            <el-option v-for="v,k in statusOpts" :key="k" :label="v" :value="parseInt(k)" />
          </el-select>

        </el-form-item>
        <el-form-item label="创建日期" prop="createdAt">
          <template #label>
            <span>
              创建日期
              <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）">
                <el-icon><QuestionFilled /></el-icon>
              </el-tooltip>
            </span>
          </template>
          <el-date-picker v-model="searchInfo.startCreatedAt" type="datetime" placeholder="开始日期" :disabled-date="time=> searchInfo.endCreatedAt ? time.getTime() > searchInfo.endCreatedAt.getTime() : false" />
          —
          <el-date-picker v-model="searchInfo.endCreatedAt" type="datetime" placeholder="结束日期" :disabled-date="time=> searchInfo.startCreatedAt ? time.getTime() < searchInfo.startCreatedAt.getTime() : false" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button type="primary" icon="plus" @click="openDialog">新增</el-button>
        <el-popover v-model:visible="deleteVisible" :disabled="!multipleSelection.length" placement="top" width="160">
          <p>确定要删除吗？</p>
          <div style="text-align: right; margin-top: 8px;">
            <el-button type="primary" link @click="deleteVisible = false">取消</el-button>
            <el-button type="primary" @click="onDelete">确定</el-button>
          </div>
          <template #reference>
            <el-button icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="deleteVisible = true">删除</el-button>
          </template>
        </el-popover>
      </div>
      <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        border
        @selection-change="handleSelectionChange"
      >
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" label="日期" width="180" fixed="right">
          <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <!-- <el-table-column align="left" label="ID" show-overflow-tooltip prop="ID" width="80" fixed="left" /> -->
        <el-table-column align="left" label="客户端编号" prop="clientId" width="120">
          <template #default="scope">
            <el-tag size="small" :type="statusTags[scope.row.clientId]" effect="dark">
              {{ clientIdOpts[scope.row.clientId] }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column align="left" label="通道类型" prop="channelTyp" width="120">
          <template #default="scope">
            <el-tag size="small" :type="statusTags[scope.row.channelTyp]" effect="dark">
              {{ channelTypOpts[scope.row.channelTyp] }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column align="left" label="配置名称" show-overflow-tooltip prop="name" width="120" />
        <el-table-column align="left" label="网页地址" show-overflow-tooltip prop="url" width="200">
          <template #default="scope">
            <a :href="'https://'+scope.row.url" target="_blank">{{ scope.row.url }}</a>
          </template>
        </el-table-column>
        <el-table-column align="left" label="本地ip" show-overflow-tooltip prop="localIp" width="120" />
        <el-table-column align="left" label="本地端口" show-overflow-tooltip prop="localPort" width="120" />
        <el-table-column align="left" label="远程端口" show-overflow-tooltip prop="remotePort" width="120" />
        <el-table-column align="left" label="状态" prop="status" width="120">
          <template #default="scope">
            <el-tag size="small" :type="statusTags[scope.row.status]" effect="dark">
              {{ statusOpts[scope.row.status] }}
            </el-tag>
          </template>
        </el-table-column>
        <el-table-column align="left" label="其他配置" show-overflow-tooltip prop="config" width="120" />
        <el-table-column align="left" label="操作" min-width="120" fixed="right">
          <template #default="scope">
            <el-button type="primary" link icon="edit" class="table-button" @click="updateFrpcClientConfigFunc(scope.row)">变更</el-button>
            <el-button type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>
            <el-button type="primary" link icon="delete" @click="copRow(scope.row)">拷贝</el-button>
          </template>
        </el-table-column>
      </el-table>
      <div class="gva-pagination">
        <el-pagination
          layout="total, sizes, prev, pager, next, jumper"
          :current-page="page"
          :page-size="pageSize"
          :page-sizes="[10, 30, 50, 100]"
          :total="total"
          @current-change="handleCurrentChange"
          @size-change="handleSizeChange"
        />
      </div>
    </div>
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" :title="type==='create'?'添加':'修改'" destroy-on-close>
      <el-scrollbar height="500px">
        <el-form ref="elFormRef" :model="formData" label-position="right" :rules="rule" label-width="100px" :inline="true">
          <el-form-item label="客户端编号:" prop="clientId">
            <el-select v-model="formData.clientId" placeholder="Please select" style="width:100%" :clearable="true">
              <el-option v-for="v,k in clientIdOpts" :key="k" :label="v" :value="parseInt(k)" />
            </el-select>
          </el-form-item>
          <el-form-item label="通道类型:" prop="channelTyp">
            <el-select v-model="formData.channelTyp" placeholder="Please select" style="width:100%" :clearable="true">
              <el-option v-for="v,k in channelTypOpts" :key="k" :label="v" :value="parseInt(k)" />
            </el-select>
          </el-form-item>
          <br>
          <el-form-item label="配置名称:" prop="name">
            <el-input v-model="formData.name" :clearable="true" placeholder="请输入配置名称" />
          </el-form-item>
          <br>
          <el-form-item label="本地ip:" prop="localIp">
            <el-input v-model="formData.localIp" placeholder="请输入本地ip" style="width:200px" />:
            <el-input v-model.number="formData.localPort" placeholder="请输入本地端口" style="width:100px" />=>
            <el-input v-model.number="formData.remotePort" placeholder="请输入远程端口" style="width:100px" />
          </el-form-item>
          <br>
          <el-form-item label="状态:" disabled prop="status">
            <el-select v-model="formData.status" disabled placeholder="Please select" style="width:100%" :clearable="true">
              <el-option v-for="v,k in statusOpts" :key="k" :label="v" :value="parseInt(k)" />
            </el-select>
          </el-form-item>
          <el-form-item label="网页地址:" prop="url">
            <el-input v-model="formData.url" :clearable="true" placeholder="请输入网页地址" />
          </el-form-item>
          <el-form-item label="其他配置:" prop="config" style="width:80%">
            <el-input v-model="formData.config" :clearable="true" placeholder="请输入其他配置" type="textarea" :autosize="{ minRows: 10, maxRows: 20 }" style="width:100%" />
          </el-form-item>
        </el-form>
      </el-scrollbar>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeDialog">取 消</el-button>
          <el-button type="primary" @click="enterDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import {
  getFrpcClientList
} from '@/api/frpcClient'

const statusTags = { 0: 'warning', 1: 'success', 2: 'info', 3: 'primary', 4: 'danger', 9: 'danger' }
let clientIdOpts = { 1: 'Normal', 2: 'Close' }
const channelTypOpts = { 1: 'Tcp', 2: 'Xtcp', 3: 'Https' }
const statusOpts = { 0: 'Stop', 1: 'Runing', 9: 'Close' }

const getClientIdOpts = async() => {
  const table = await getFrpcClientList({ pageSize: 9999, isMap: true })
  console.log(table)
  if (table.code === 0) {
    clientIdOpts = table.data
  }
}
getClientIdOpts()

import {
  createFrpcClientConfig,
  deleteFrpcClientConfig,
  deleteFrpcClientConfigByIds,
  updateFrpcClientConfig,
  findFrpcClientConfig,
  getFrpcClientConfigList
} from '@/api/frpcClientConfig'

// 富文本组件
import RichEdit from '@/components/richtext/rich-edit.vue'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict, ReturnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

defineOptions({
  name: 'FrpcClientConfig'
})

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
  name: '',
  localIp: '127.0.0.1',
  localPort: 22,
  remotePort: 22,
  config: '',
  channelTyp: 1
})

// 验证规则
const rule = reactive({
  clientId: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  },
  ],
  channelTyp: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  },
  ],
  name: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  },
  {
    whitespace: true,
    message: '不能只输入空格',
    trigger: ['input', 'blur'],
  }
  ],
})

const searchRule = reactive({
  createdAt: [
    { validator: (rule, value, callback) => {
      if (searchInfo.value.startCreatedAt && !searchInfo.value.endCreatedAt) {
        callback(new Error('请填写结束日期'))
      } else if (!searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt) {
        callback(new Error('请填写开始日期'))
      } else if (searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt && (searchInfo.value.startCreatedAt.getTime() === searchInfo.value.endCreatedAt.getTime() || searchInfo.value.startCreatedAt.getTime() > searchInfo.value.endCreatedAt.getTime())) {
        callback(new Error('开始日期应当早于结束日期'))
      } else {
        callback()
      }
    }, trigger: 'change' }
  ],
})

const elFormRef = ref()
const elSearchFormRef = ref()

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

// 重置
const onReset = () => {
  searchInfo.value = {}
  getTableData()
}

// 搜索
const onSubmit = () => {
  getClientIdOpts()
  elSearchFormRef.value?.validate(async(valid) => {
    if (!valid) return
    page.value = 1
    pageSize.value = 10
    getTableData()
  })
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async() => {
  const table = await getFrpcClientConfigList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async() => {
}

// 获取需要的字典 可能为空 按需保留
setOptions()

// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
  multipleSelection.value = val
}

// 删除行
const deleteRow = (row) => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    deleteFrpcClientConfigFunc(row)
  })
}

const copRow = (row) => {
  formData.value = row
  formData.value.name = ''
  formData.value.ID = 0
  openDialog()
}

// 批量删除控制标记
const deleteVisible = ref(false)

// 多选删除
const onDelete = async() => {
  const ids = []
  if (multipleSelection.value.length === 0) {
    ElMessage({
      type: 'warning',
      message: '请选择要删除的数据'
    })
    return
  }
  multipleSelection.value &&
        multipleSelection.value.map(item => {
          ids.push(item.ID)
        })
  const res = await deleteFrpcClientConfigByIds({ ids })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '删除成功'
    })
    if (tableData.value.length === ids.length && page.value > 1) {
      page.value--
    }
    deleteVisible.value = false
    getTableData()
  }
}

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateFrpcClientConfigFunc = async(row) => {
  const res = await findFrpcClientConfig({ ID: row.ID })
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data.refrpcClientConfig
    dialogFormVisible.value = true
  }
}

// 删除行
const deleteFrpcClientConfigFunc = async(row) => {
  const res = await deleteFrpcClientConfig({ ID: row.ID })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '删除成功'
    })
    if (tableData.value.length === 1 && page.value > 1) {
      page.value--
    }
    getTableData()
  }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)

// 打开详情
const getDetails = async(row) => {
  // 打开弹窗
  console.log(row)
  const res = await findFrpcClientConfig({ ID: row.ID })
  if (res.code === 0) {
    formData.value = res.data.refrpcClientConfig
    openDetailShow()
  }
}

// 打开弹窗
const openDialog = () => {
  type.value = 'create'
  dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
  dialogFormVisible.value = false
  formData.value = {
    name: '',
    localIp: '127.0.0.1',
    localPort: 22,
    remotePort: 22,
    channelTyp: 1
  }
}
// 弹窗确定
const enterDialog = async() => {
  elFormRef.value?.validate(async(valid) => {
    if (!valid) return
    let res
    switch (type.value) {
      case 'create':
        res = await createFrpcClientConfig(formData.value)
        break
      case 'update':
        res = await updateFrpcClientConfig(formData.value)
        break
      default:
        res = await createFrpcClientConfig(formData.value)
        break
    }
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '创建/更改成功'
      })
      closeDialog()
      getTableData()
    }
  })
}

</script>

<style>

</style>
