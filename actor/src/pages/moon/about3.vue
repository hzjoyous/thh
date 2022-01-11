<template>
  <n-result status="404" title="！" description="这么大" size="huge">
    <template #footer>
      <n-button>哦</n-button>
    </template>
  </n-result>
</template>
<script setup>
import {NResult,NButton} from 'naive-ui'
</script>
