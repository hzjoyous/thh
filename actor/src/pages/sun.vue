<template>
  <n-layout position="absolute">
    <n-layout-header position="absolute"
                     style="height: 32px; padding: 8px;"
                     bordered
    >
      <router-link to="/sun/index">index</router-link>
      |
      <router-link to="/sun/sunMain">main</router-link>
      |
      <router-link to="/sun/about">about</router-link>
    </n-layout-header>

    <n-layout-content position="absolute" content-style="padding: 1px;" style="top: 64px; bottom: 64px;">
      <router-view></router-view>
    </n-layout-content>
  </n-layout>
</template>

<script>
import {NDataTable, NLayout, NLayoutContent, NLayoutFooter, NLayoutHeader, NLayoutSider, NSpace} from 'naive-ui';
import {defineComponent} from "vue";
import {useIsMobile, useIsSmallDesktop, useIsTablet} from '../utils/composables';


export default defineComponent({
  components: {
    NSpace,
    NLayout,
    NLayoutFooter,
    NLayoutSider,
    NLayoutContent,
    NLayoutHeader,
    NDataTable
  },
  setup() {
    const isMobileRef = useIsMobile()
    const isTabletRef = useIsTablet()
    const isSmallDesktop = useIsSmallDesktop()
    return {
      isMobile: isMobileRef,
      isTabletRef: isTabletRef,
      isSmallDesktop: isSmallDesktop,
    }
  }
})
</script>
<style>
</style>