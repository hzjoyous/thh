<template>
  <n-dynamic-input v-model:value="value" placeholder="请输入" :min="3" :max="6" :on-create="add" :on-remove="remove"
                   :on-update="update"/>
  <pre>
  {{ JSON.stringify(value, 0, 2) }}
  </pre>
</template>

<script>
import {defineComponent, ref} from 'vue'
import {NDynamicInput} from "naive-ui"

export default defineComponent({
  components: {NDynamicInput},
  setup() {
    return {
      value: ref(['', '', ''])
    }
  },
  methods: {
    add(item) {
      console.log(item)
    },
    remove(item) {
      console.log(item)
    },
    update(item) {
      console.log(item)
    }
  }
})
</script>