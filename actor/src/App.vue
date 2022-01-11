<template>
  <!--  <img alt="Vue logo" src="./assets/logo.png" />-->
  <!--  <HelloWorld msg="Hello Vue 3 + Vite" />-->
  <n-message-provider>
    <n-config-provider :locale="zhCN" :date-locale="dateZhCN">
      <n-layout position="absolute">
        <n-layout-header position="absolute" style="height: 64px; padding: 24px;" bordered>
          <n-button type="default">
            <router-link to="/sun">sun</router-link>
          </n-button>
          <n-button>
            <router-link to="/moon">moon</router-link>
          </n-button>

        </n-layout-header>
        <n-layout position="absolute" style="top: 64px; ">
          <router-view></router-view>
        </n-layout>

      </n-layout>
<!--      <n-affix-->
<!--          :top="120"-->
<!--          :trigger-top="60"-->
<!--      >-->
<!--        <n-tag>底部触发距离 60px</n-tag>-->
<!--      </n-affix>-->
    </n-config-provider>
  </n-message-provider>
</template>

<script>
import {NTag,NAffix,dateZhCN, NButton, NConfigProvider, NLayout, NLayoutHeader, NMessageProvider, NSpace, zhCN} from 'naive-ui'
import {defineComponent} from 'vue'

// import HelloWorld from './components/HelloWorld.vue'

// This starter template is using Vue 3 experimental <script setup> SFCs
// Check out https://github.com/vuejs/rfcs/blob/master/active-rfcs/0040-script-setup.md
export default defineComponent({
  components: {
    NConfigProvider, NSpace, NLayout, NButton, NLayoutHeader, NMessageProvider,NAffix,NTag
  },
  setup() {
    return {
      zhCN,
      dateZhCN
    }
  }
})
</script>

<style>
@import "app.css";
</style>
