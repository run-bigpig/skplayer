<script setup>
import {onMounted, ref} from "vue";
import axios from "axios";
import {Message} from "@arco-design/web-vue";
import {useSettingStore} from "@/stores/setting.js";
const store = useSettingStore()
const fieldNames = {value:'key',label:'name'}
const options = ref([])
const form = ref({
  source:''
})
const getSource = () => {
  axios.get('/api/getsource').then(res => {
     let data = res.data.data
     options.value = data.list.map(item => {
       return {
         key:item.key,
         name:item.name
       }
     })
     form.value.source = data.default.key
     store.setSetting(form.value)
  })
}
onMounted(()=>{
  getSource()
})

const onSubmit = () => {
 axios.post('/api/setting',form.value).then(res => {
     store.setSetting(form.value)
     Message.info({
       content:res.data.msg,
       showIcon: true
     })
 })
}

</script>

<template>
<a-card :style="{padding:'20px'}" :bordered="false">
      <a-form :model="form" layout="horizontal">
        <a-row>
          <a-form-item label="数据源" field="source">
            <a-select v-model="form.source" placeholder="请选择默认数据源" :options=options :field-names="fieldNames"></a-select>
          </a-form-item>
        </a-row>
        <a-row>
          <a-form-item>
            <a-button type="primary" @click="onSubmit">保存</a-button>
          </a-form-item>
        </a-row>
      </a-form>
</a-card>
</template>

<style scoped>

</style>