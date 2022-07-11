<template>
  <n-card style="margin:0 auto">

    <n-timeline :size="'large'">
      <n-timeline-item v-for="timeInfo in dayInfoList" :type="timeInfo.type"
                       :title="timeInfo.title"
                       :content="timeInfo.content"
                       :time="timeInfo.time"
                       :line-type="timeInfo.lineType"
      />
    </n-timeline>

  </n-card>
</template>
<script>
import {NButton, NCalendar, NCard, NResult, NTimeline, NTimelineItem, useMessage} from 'naive-ui';
import moment from "moment"

export default {
  components: {
    NResult, NButton, NCalendar, NCard, NTimelineItem, NTimeline
  },
  setup() {
    const message = useMessage()
    let dayInfoList = [];
    let nowT = moment()
    let t = moment(moment().format("YYYY-01-01"))
    for (let i = 1; i < 12; i++) {
      t.add(1, "months")
      let type = 'warning'
      let lineType = 'dashed'
      console.log(t.format('M'), nowT.format('M'))
      if (parseInt(t.format('M')) > parseInt(nowT.format('M'))) {
        type = 'success'
        lineType = 'default'
      }
      let timeInfo = t.format('YYYY-MM-DD')
      dayInfoList.push({
        title: timeInfo,
        time: timeInfo,
        // content: timeInfo,
        type: type,
        lineType: lineType
      })
    }
    dayInfoList.sort(function (item1, item2) {
      return item1.time > item2.time ? -1 : 1
    })

    dayInfoList.push({title: "start"})
    dayInfoList.unshift({title: "end", type: "success"});
    return {
      dayInfoList: dayInfoList,
      success() {
        message.success(
            `还挺大`
        )
      },
    }
  }
}
</script>

<style>
.n-card {
  max-width: 1000px;
}
</style>