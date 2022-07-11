<template>
  <n-space vertical :size="12">
    <n-input v-model:value="pattern" placeholder="搜索"/>
    <n-tree :pattern="pattern" :data="data" block-line/>
  </n-space>
</template>

<script>
import {defineComponent, ref} from 'vue'
import {NDynamicInput, NInput, NSpace, NTree} from "naive-ui"

function createData(level = 4, baseKey = '') {
  if (!level) return undefined
  return Array.apply(null, {length: 2}).map((_, index) => {
    const key = '' + baseKey + level + index
    return {
      label: createLabel(level),
      key,
      children: createData(level - 1, key)
    }
  })
}

function createLabel(level) {
  if (level === 4) return '道生一'
  if (level === 3) return '一生二'
  if (level === 2) return '二生三'
  if (level === 1) return '三生万物'
}

export default defineComponent({
  components: {NDynamicInput, NTree, NInput, NSpace},
  setup() {
    let data = [];
    for (let deep1 = 1; deep1 <= 70; deep1++) {
      let tempNode1 = {
        key: deep1.toString(),
        label: "deep1" + deep1,
        children: []
      };
      for (let deep2 = 1; deep2 <= 10; deep2++) {
        let tempNode2 = {
          key: deep1.toString() + ':' + deep2.toString(),
          label: "deep2" + deep2,
          children: []
        }
        for (let deep3 = 1; deep3 <= 2; deep3++) {
          let tempNode3 = {
            key: deep1.toString() + ':' + deep2.toString()+":"+deep3.toString(),
            label: "deep3" + deep3,
            children: []
          }
          tempNode2.children.push(tempNode3)
        }
        tempNode1.children.push(tempNode2)
      }
      data.push(tempNode1)
    }
    return {
      data: data,
      pattern: ref(''),
    }
  }
})
</script>