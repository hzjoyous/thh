<template>
  <n-result status="404" title="！" description="接受真相就是这么简单" size="huge">
    <template #footer>
      <n-button @click="success">哦</n-button>
    </template>
  </n-result>
</template>
<script>
import {NResult, NButton, NCalendar, useMessage} from 'naive-ui'

export default {
  components: {
    NResult, NButton, NCalendar
  },
  setup() {
    const message = useMessage()

    return {
      success() {
        message.success(
            `一切尽在不言中`
        )
      }
    }
  }
}
</script>

