<template>
  <n-row>
    <n-col :span="12">
      <n-statistic label="统计数据" :value="99">
        <template #prefix>
          <n-icon>

          </n-icon>
        </template>
        <template #suffix>/ 100</template>
      </n-statistic>
    </n-col>
    <n-col :span="12">
      <n-statistic label="活跃用户">1,234,123</n-statistic>
    </n-col>
  </n-row>
</template>

<script>
import {NRow,NCol,NStatistic,NIcon} from 'naive-ui'

export default {
  components: {
    NRow,NCol,NStatistic,NIcon
  }
}
</script>