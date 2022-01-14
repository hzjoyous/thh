import sun from "./pages/sun.vue";
import sunMain from "./pages/sun/main.vue";
import about from "./pages/sun/about.vue";
import moon from "./pages/moon.vue";
import listTodo from "./pages/moon/listTodo.vue";
import tmp from "./pages/moon/tmp.vue";
import moonabout1 from "./pages/moon/about1.vue";
import moonabout2 from "./pages/moon/about2.vue";
import moonabout3 from "./pages/moon/about3.vue";
import moonabout4 from "./pages/moon/about4.vue";
import moonabout5 from "./pages/moon/about5.vue";
import moonabout6 from "./pages/moon/about6.vue";
import markdown from "./pages/moon/markdown.vue";

export default [
    // {path: '/', component: about},
    {path: '/:catchAll(.*)*', name: 'index', redirect: '/sun/sunMain'},
    // {path: '/:catchAll(.*)*', name: 'index', redirect: '/moon/tmp'},
    {
        path: '/sun', component: sun, children: [
            {name: '', path: '', component: sunMain},
            {name: 'sunMain', path: 'sunMain', component: sunMain},
            {name: 'about', path: 'about', component: about},
        ]
    },
    {
        belongMenu: true,
        path: '/moon', component: moon, children: [
            {name: '', path: '', component: moonabout1, belongMenu: false},
            {name: 'moonabout1', path: 'moonabout1', component: moonabout1, belongMenu: true},
            {name: 'moonabout2', path: 'moonabout2', component: moonabout2, belongMenu: true},
            {name: 'moonabout3', path: 'moonabout3', component: moonabout3, belongMenu: true},
            {name: 'moonabout4', path: 'moonabout4', component: moonabout4, belongMenu: true},
            {name: 'moonabout5', path: 'moonabout5', component: moonabout5, belongMenu: true},
            {name: 'moonabout6', path: 'moonabout6', component: moonabout6, belongMenu: true},
            {name: 'markdown', path: 'markdown', component: markdown, belongMenu: true},
            {name: 'listTodo', path: 'listTodo', component: listTodo, belongMenu: true},
            {name: 'tmp', path: 'tmp', component: tmp, belongMenu: true},
        ]
    },
]