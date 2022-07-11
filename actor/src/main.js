import {createApp} from 'vue'
import App from './App.vue'
import axios from 'axios'
import {createRouter, createWebHashHistory} from "vue-router";
import routes from './routes.js'
import Mermaid from 'mermaid'
import katex from 'katex'
// !import 不能忘记引入katex的样式
import 'katex/dist/katex.css'
// 引入katex下的自动渲染函数
import renderMathInElement from 'katex/contrib/auto-render/auto-render'

// 定义自动渲染的配置参数,这些参数根据你的需求进行修改，下面的参数是官网上抄下来的
const renderOption = {
    delimiters: [
        {left: '$$', right: '$$', display: true},
        {left: '$', right: '$', display: false},
        {left: '\\(', right: '\\)', display: false},
        {left: '\\[', right: '\\]', display: true}
    ],
    throwOnError : false
}
// 2. 定义一些路由
// 每个路由都需要映射到一个组件。
// 我们后面再讨论嵌套路由。


// 3. 创建路由实例并传递 `routes` 配置
// 你可以在这里输入更多的配置，但我们在这里
// 暂时保持简单
const router = createRouter({
    // 4. 内部提供了 history 模式的实现。为了简单起见，我们在这里使用 hash 模式。
    history: createWebHashHistory(),
    routes: routes, // `routes: routes` 的缩写
})

// let title;  // 用于临时存放原来的title内容
// window.onblur = function(){
//     // onblur时先存下原来的title,再更改title内容
//     title = document.title;
//     document.title = title+"🤔";
// };
// window.onfocus = function () {
//     // onfocus时原来的title不为空才替换回去
//     // 防止页面还没加载完成且onblur时title=undefined的情况
//     if(title) {
//         document.title = title;
//     }
// }

const app = createApp(App)
app.config.globalProperties.$axios = axios
app.config.globalProperties.$formula = function (dom) {
    renderMathInElement(dom, renderOption)
}
app.use(router)
app.mount('#app')

