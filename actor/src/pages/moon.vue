<template>
  <n-layout position="absolute">
    <n-layout-header style="height: 64px; padding: 18px;" bordered
                     position="absolute"
    >
      <n-layout style="float: right">
        <n-dropdown trigger="hover" :options="options" >
          <n-button>找个地方休息</n-button>
        </n-dropdown>
      </n-layout>

    </n-layout-header
    >
    <n-layout position="absolute" style="top: 64px; bottom: 64px;" has-sider>

      <n-layout-sider
          content-style="padding-top: 24px;"
          :native-scrollbar="false" bordered
          v-show="!isMobile"
      >
        <n-menu
            :options="menuOptions"
        />

      </n-layout-sider>
      <n-layout content-style="padding: 24px;" :native-scrollbar="false">
        <router-view></router-view>
      </n-layout>
    </n-layout>
    <n-layout-footer bordered
                     position="absolute"
                     style="height: 64px; padding: 24px;"
    >
      世事洞明皆学问 | isMobile:<span>{{ isMobile }}</span>
      |isTabletRef:<span>{{ isTabletRef }}</span>
      |isSmallDesktop:<span>{{ isSmallDesktop }}</span>
    </n-layout-footer>
  </n-layout>
</template>

<script>
import {h, ref} from 'vue'
import {RouterLink} from 'vue-router'
import {BookOutline as BookIcon,} from '@vicons/ionicons5'
import {
  NButton,
  NDropdown,
  NH2,
  NIcon,
  NLayout,
  NLayoutContent,
  NLayoutFooter,
  NLayoutHeader,
  NLayoutSider,
  NMenu,
  NSpace
} from 'naive-ui'
import routes from "../routes";
import {useIsMobile, useIsSmallDesktop, useIsTablet} from '../utils/composables';

function renderIcon(icon) {
  return () => h(NIcon, null, {default: () => h(icon)})
}

export default {
  setup(type, children) {

    const buildMenuOption = function (parentPath, route) {
      if (route.belongMenu === undefined || route.belongMenu !== true) {
        return []
      }
      let menuList = [];
      if (route.children !== undefined && route.children.length > 0) {
        for (let key in route.children) {
          menuList.push(...buildMenuOption(route.path, route.children[key]))
        }
      } else {
        let path
        if (parentPath.charAt(parentPath.length - 1) !== '/' && route.path.charAt(0) !== '/') {
          path = parentPath + '/' + route.path
        } else {
          path = parentPath + route.path
        }
        menuList.push({
          label: () =>
              h(
                  RouterLink,
                  {
                    to: {
                      path: path,
                    }
                  },
                  {default: () => route.name}
              )
          ,
          key: path,
          icon: renderIcon(BookIcon)
        })
        return menuList
      }
      return menuList;
    };


    let menuOptions = []
    for (let i = 0; i < routes.length; i++) {
      menuOptions.push(...buildMenuOption('', routes[i]))
    }
    menuOptions.push({
          label: () => h(
              RouterLink,
              {
                to: {
                  path: '/moon/moonabout2',
                }
              },
              {default: () => '回家'}),
          key: 'hear-the-wind-sing',
          icon: renderIcon(BookIcon)
        },
        {
          label: '1973年的弹珠玩具',
          key: 'pinball-1973',
          icon: renderIcon(BookIcon),
          children: [
            {
              label: '鼠',
              key: 'rat'
            }
          ]
        })
    const isMobileRef = useIsMobile()
    const isTabletRef = useIsTablet()
    const isSmallDesktop = useIsSmallDesktop()

    return {

      isMobile: isMobileRef,
      isTabletRef: isTabletRef,
      isSmallDesktop: isSmallDesktop,
      menuOptions: menuOptions,
      options: menuOptions,
      collapsed: ref(true),
      renderMenuLabel(option) {
        return option.path
      },
      renderMenuIcon(option) {
        // 渲染图标占位符以保持缩进
        if (option.key === 'sheep-man') return true
        // 返回 falsy 值，不再渲染图标及占位符
        if (option.key === 'food') return null
        return h(NIcon, null,)
      },
      expandIcon() {
        return h(NIcon, null,)
      }

    }

  },
  components: {
    NSpace, NLayout, NLayoutFooter, NLayoutSider, NLayoutContent, NLayoutHeader, NH2, NMenu, NDropdown, NButton
  }
}
</script>
<style>
</style>