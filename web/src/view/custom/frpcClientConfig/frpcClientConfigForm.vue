<template>
  <div>
    <div class="gva-form-box">
      <el-form ref="elFormRef" :model="formData" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="客户端编号:" prop="clientId">
          <el-select v-model="formData.clientId" placeholder="请选择" style="width:100%" :clearable="true">
            <el-option v-for="item in [32]" :key="item" :label="item" :value="item" />
          </el-select>
        </el-form-item>
        <el-form-item label="通道类型:" prop="channelTyp">
          <el-select v-model="formData.channelTyp" placeholder="请选择" style="width:100%" :clearable="true">
            <el-option v-for="item in [32]" :key="item" :label="item" :value="item" />
          </el-select>
        </el-form-item>
        <el-form-item label="配置名称:" prop="name">
          <el-input v-model="formData.name" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="本地ip:" prop="localIp">
          <el-input v-model="formData.localIp" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="本地端口:" prop="localPort">
          <el-input v-model.number="formData.localPort" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="远程端口:" prop="remotePort">
          <el-input v-model.number="formData.remotePort" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="状态:" prop="status">
          <el-select v-model="formData.status" placeholder="请选择" style="width:100%" :clearable="true">
            <el-option v-for="item in [8]" :key="item" :label="item" :value="item" />
          </el-select>
        </el-form-item>
        <el-form-item label="其他配置:" prop="config">
          <RichEdit v-model="formData.config" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import {
  createFrpcClientConfig,
  updateFrpcClientConfig,
  findFrpcClientConfig
} from '@/api/frpcClientConfig'

defineOptions({
  name: 'FrpcClientConfigForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
// 富文本组件
import RichEdit from '@/components/richtext/rich-edit.vue'

const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
  name: '',
  localIp: '',
  localPort: 0,
  remotePort: 0,
  config: '',
})
// 验证规则
const rule = reactive({
  clientId: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  }],
  channelTyp: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  }],
  name: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  }],
  localIp: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  }],
  localPort: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  }],
})

const elFormRef = ref()

// 初始化方法
const init = async() => {
  // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
  if (route.query.id) {
    const res = await findFrpcClientConfig({ ID: route.query.id })
    if (res.code === 0) {
      formData.value = res.data.refrpcClientConfig
      type.value = 'update'
    }
  } else {
    type.value = 'create'
  }
}

init()
// 保存按钮
const save = async() => {
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
    }
  })
}

// 返回按钮
const back = () => {
  router.go(-1)
}

</script>

<style>
</style>
