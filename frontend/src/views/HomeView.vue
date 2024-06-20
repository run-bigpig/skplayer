<script setup>
import {onMounted, ref, watch} from "vue";
import axios from "axios";
import {useRoute} from "vue-router";

const route = useRoute()
const dataSource = ref([])
const loading = ref(false)
const getList = () => {
  loading.value = true
  axios.post("/api/list", {
    class: 0
    , page: 1
    , pageSize: 24
  }).then(res => {
    dataSource.value = res.data.data.list
    loading.value = false
  })
}
onMounted(() => {
  getList()
})
</script>

<template>
  <a-list
      :grid-props="{ gutter: [10, 10], xs: 24, sm: 12, md: 4, lg: 4, xl:4, xxl: 4 }"
      :data="dataSource"
      :bordered="false"
      :max-height="960"
      :scrollbar="false"
      :loading=loading
  >
    <template #item="{ item }">
      <a-list-item @click="$router.push('/play/'+item.id)">
        <a-image
            :src=item.img
            width="260"
            height="350"
            fit="cover"
            footer-position="outer"
            show-loader
        >
          <template #extra>
            <h4>{{ item.name }}</h4>
          </template>
        </a-image>
      </a-list-item>
    </template>
  </a-list>
</template>

<style scoped>

</style>