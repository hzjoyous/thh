<template>
  <div v-html="result">
  </div>
</template>

<script>
import {NCol, NIcon, NRow, NStatistic} from 'naive-ui'
import MarkdownIt from "markdown-it"
import {defineComponent} from "vue"

export default defineComponent({
  name: "",
  props: {},
  components: {
    NRow, NCol, NStatistic, NIcon
  },
  setup() {
    const md = new MarkdownIt();
    const p = function (md, option) {
      console.log(md)
      console.log(option)
    };
    md.use(p)



    const result = md.render(`
# markdown-it test!

## two

### three

#### four

##### five


\`\`\`sql
insert into dogs values(1,2,3);
\`\`\`
- [ ] 111
- [ ] 112

[baidu](https://www.baidu.com)

|12|12|
|---|---|
|1 1| 112|

$\\sqrt{3x-1}+(1+x)^2$

    `);


    return {
      result: result
    }
  }
})
</script>