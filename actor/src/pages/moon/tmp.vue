<template>
  <h1>sessionStorage</h1>
  <p>{{ sessionStorageData }}</p>
  <h1>localStorage</h1>
  <p>{{ localStorageData }}</p>
  <h1>cookie</h1>
  <h1>session</h1>
  <h2>tmp in sessionStorage/localStorage 如果存在+1 </h2>
  <button @click="set"> 设置</button>
  <button @click="showNew"> 展示最新</button>
</template>

<script>
import {defineComponent, ref} from 'vue'
import {NDynamicInput} from "naive-ui"

export default defineComponent({
  components: {NDynamicInput},
  setup() {
    let sessionStorageData = ref("")
    let localStorageData = ref("")

    return {
      sessionStorageData: sessionStorageData,
      localStorageData: localStorageData,
    }
  },
  methods: {
    set() {
      let localTmp = localStorage.getItem("tmp")
      console.log(localTmp)
      localStorage.setItem("tmp", (Number(localTmp) + 1).toString())
      let sessionTmp = sessionStorage.getItem("tmp")
      sessionStorage.setItem("tmp", (Number(sessionTmp) + 1).toString())
      this.sessionStorageData = JSON.stringify(sessionStorage)
      this.localStorageData = JSON.stringify(localStorage)
    },
    showNew(item) {
      this.sessionStorageData = JSON.stringify(sessionStorage)
      this.localStorageData = JSON.stringify(localStorage)
    }
  }
})
</script>