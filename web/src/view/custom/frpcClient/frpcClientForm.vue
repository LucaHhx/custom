<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="用户:" prop="user">
          <el-input v-model="formData.user" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="服务器ip:" prop="serverAddr">
          <el-input v-model="formData.serverAddr" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="服务器端口:" prop="serverPort">
          <el-input v-model.number="formData.serverPort" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="网页ip:" prop="webAddr">
          <el-input v-model="formData.webAddr" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="网页端口:" prop="webPort">
          <el-input v-model.number="formData.webPort" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="网页映射端口:" prop="webMapPort">
          <el-input v-model.number="formData.webMapPort" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="网页用户:" prop="webUser">
          <el-input v-model="formData.webUser" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="网页密码:" prop="webPassword">
          <el-input v-model="formData.webPassword" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="网页地址:" prop="url">
          <el-input v-model="formData.url" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="额外配置:" prop="config">
          <RichEdit v-model="formData.config"/>
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
  createFrpcClient,
  updateFrpcClient,
  findFrpcClient
} from '@/api/frpcClient'

defineOptions({
    name: 'FrpcClientForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
// 富文本组件
import RichEdit from '@/components/richtext/rich-edit.vue'

const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
            user: '',
            serverAddr: '',
            serverPort: 0,
            webAddr: '',
            webPort: 0,
            webMapPort: 0,
            webUser: '',
            webPassword: '',
            url: '',
            config: '',
        })
// 验证规则
const rule = reactive({
               user : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               serverAddr : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               serverPort : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               webAddr : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               webPort : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               webMapPort : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findFrpcClient({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.refrpcClient
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createFrpcClient(formData.value)
               break
             case 'update':
               res = await updateFrpcClient(formData.value)
               break
             default:
               res = await createFrpcClient(formData.value)
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
