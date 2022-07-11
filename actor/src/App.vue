<template>
  <!--  <img alt="Vue logo" src="./assets/logo.png" />-->
  <!--  <HelloWorld msg="Hello Vue 3 + Vite" />-->
  <n-message-provider>
    <n-config-provider :locale="zhCN" :date-locale="dateZhCN">
      <n-layout position="absolute">
<!--        <n-layout-header position="absolute" style="height: 64px; padding: 24px;" bordered>-->
<!--          <n-button type="default">-->
<!--            <router-link to="/sun">sun</router-link>-->
<!--          </n-button>-->
<!--          <n-button>-->
<!--            <router-link to="/moon">moon</router-link>-->
<!--          </n-button>-->

<!--        </n-layout-header>-->
<!--        <n-layout position="absolute" style="top: 64px; ">-->
        <n-layout position="absolute" style="top: 20px; ">
          <router-view></router-view>
        </n-layout>

        <div
            style="right: 60px;bottom: 100px;position:absolute"
        >
          <n-dropdown trigger="hover" :options="options" @select="handleSelect">
            <n-button circle size="large" >
              <template #icon>
                <n-icon><cash-icon /></n-icon>
              </template>
            </n-button>
          </n-dropdown>

        </div>
      </n-layout>
    </n-config-provider>
  </n-message-provider>
</template>

<script>
import {NTag,NAffix,dateZhCN,NIcon,NDropdown, NButton, NConfigProvider, NLayout, NLayoutHeader, NMessageProvider, NSpace, zhCN} from 'naive-ui'
import {defineComponent, h} from 'vue'
import { CashOutline as CashIcon } from '@vicons/ionicons5'
import {RouterLink} from "vue-router";


// import HelloWorld from './components/HelloWorld.vue'

// This starter template is using Vue 3 experimental <script setup> SFCs
// Check out https://github.com/vuejs/rfcs/blob/master/active-rfcs/0040-script-setup.md
export default defineComponent({
  components: {
    NConfigProvider, NSpace, NLayout, NButton, NLayoutHeader, NMessageProvider,NAffix,NTag,CashIcon,NIcon,NDropdown
  },
  setup() {
    return {
      zhCN,
      dateZhCN,
      options: [
        {
          label: '滨海湾金沙，新加坡',
          key: 'marina bay sands',
          disabled: true
        },
        {
          label: () => h(
              RouterLink,
              {
                to: {
                  path: '/moon/',
                }
              },
              {default: () => 'Moon'}),
          key: "Moon"
        },

        {
          label: () => h(
              RouterLink,
              {
                to: {
                  path: '/sun/',
                }
              },
              {default: () => 'sun'}),
          key: "sun"
        },
      ],
      handleSelect (key) {
        message.info(String(key))
      }
    }
  }
})
</script>

<style>
@import "app.css";
</style>
