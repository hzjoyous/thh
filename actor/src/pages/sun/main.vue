<template>
  <n-card style="margin:0 auto">
    <n-calendar
        @update:value="handleUpdateValue"
        #="{ year, month, date }"
        v-model:value="value"
        :is-date-disabled="isDateDisabled"
    >
      {{ year }}-{{ month }}-{{ date }}
    </n-calendar>
  </n-card>
</template>
<script>
import {NButton, NCalendar, NResult,NCard, useMessage} from 'naive-ui';
import {ref} from 'vue'
import {addDays, isYesterday} from 'date-fns'

export default {
  components: {
    NResult, NButton, NCalendar,NCard
  },
  setup() {
    const message = useMessage()

    return {
      success() {
        message.success(
            `还挺大`
        )
      },
      value: ref(addDays(Date.now(), 1).valueOf()),
      handleUpdateValue(_, {year, month, date}) {
        message.success(`${year}-${month}-${date}`)
      },
      isDateDisabled(timestamp) {
        return isYesterday(timestamp);
      }
    }
  }
}
</script>

<style>
.n-card {
  max-width: 1000px;
}
</style>