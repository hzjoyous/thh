import{d as e,N as t,a as n,b as o,c as a,e as s,f as l,g as r,h as i,C as u,i as c,j as d,z as p,k as m,l as g,R as h,r as f,o as y,m as b,w as N,n as v,p as S,u as k,q as w,s as _,t as x,v as M,x as D,y as L,A as I,B as C,D as T,E as $,F as J,G as A,H as z,I as O,J as R,K as j,L as P,M as Y,O as q,P as B,Q as H,S as U,T as E,U as F,V,W as Z,X as G,Y as K,Z as Q,_ as W,$ as X,a0 as ee,a1 as te,a2 as ne,a3 as oe,a4 as ae,a5 as se,a6 as le,a7 as re,a8 as ie,a9 as ue,aa as ce}from"./vendor.b0cdd406.js";!function(){const e=document.createElement("link").relList;if(!(e&&e.supports&&e.supports("modulepreload"))){for(const e of document.querySelectorAll('link[rel="modulepreload"]'))t(e);new MutationObserver((e=>{for(const n of e)if("childList"===n.type)for(const e of n.addedNodes)"LINK"===e.tagName&&"modulepreload"===e.rel&&t(e)})).observe(document,{childList:!0,subtree:!0})}function t(e){if(e.ep)return;e.ep=!0;const t=function(e){const t={};return e.integrity&&(t.integrity=e.integrity),e.referrerpolicy&&(t.referrerPolicy=e.referrerpolicy),"use-credentials"===e.crossorigin?t.credentials="include":"anonymous"===e.crossorigin?t.credentials="omit":t.credentials="same-origin",t}(e);fetch(e.href,t)}}();const de=e({components:{NConfigProvider:t,NSpace:n,NLayout:o,NButton:a,NLayoutHeader:s,NMessageProvider:l,NAffix:r,NTag:i,CashIcon:u,NIcon:c,NDropdown:d},setup:()=>({zhCN:p,dateZhCN:m,options:[{label:"滨海湾金沙，新加坡",key:"marina bay sands",disabled:!0},{label:()=>g(h,{to:{path:"/moon/"}},{default:()=>"Moon"}),key:"Moon"},{label:()=>g(h,{to:{path:"/sun/"}},{default:()=>"sun"}),key:"sun"}],handleSelect(e){message.info(String(e))}})}),pe={style:{right:"60px",bottom:"100px",position:"absolute"}};function me(){const e=w();return k((()=>"xs"===e.value))}function ge(){const e=w();return k((()=>"s"===e.value))}function he(){const e=w();return k((()=>"m"===e.value))}de.render=function(e,t,n,o,a,s){const l=f("router-view"),r=f("n-layout"),i=f("cash-icon"),u=f("n-icon"),c=f("n-button"),d=f("n-dropdown"),p=f("n-config-provider"),m=f("n-message-provider");return y(),b(m,null,{default:N((()=>[v(p,{locale:e.zhCN,"date-locale":e.dateZhCN},{default:N((()=>[v(r,{position:"absolute"},{default:N((()=>[v(r,{position:"absolute",style:{top:"20px"}},{default:N((()=>[v(l)])),_:1}),S("div",pe,[v(d,{trigger:"hover",options:e.options,onSelect:e.handleSelect},{default:N((()=>[v(c,{circle:"",size:"large"},{icon:N((()=>[v(u,null,{default:N((()=>[v(i)])),_:1})])),_:1})])),_:1},8,["options","onSelect"])])])),_:1})])),_:1},8,["locale","date-locale"])])),_:1})};const fe=e({components:{NSpace:n,NLayout:o,NLayoutFooter:_,NLayoutSider:x,NLayoutContent:M,NLayoutHeader:s,NDataTable:D},setup:()=>({isMobile:me(),isTabletRef:ge(),isSmallDesktop:he()})}),ye=L("index"),be=L(" | "),Ne=L("main"),ve=L(" | "),Se=L("about");fe.render=function(e,t,n,o,a,s){const l=f("router-link"),r=f("n-layout-header"),i=f("router-view"),u=f("n-layout-content"),c=f("n-layout");return y(),b(c,{position:"absolute"},{default:N((()=>[v(r,{position:"absolute",style:{height:"32px",padding:"8px"},bordered:""},{default:N((()=>[v(l,{to:"/sun/index"},{default:N((()=>[ye])),_:1}),be,v(l,{to:"/sun/sunMain"},{default:N((()=>[Ne])),_:1}),ve,v(l,{to:"/sun/about"},{default:N((()=>[Se])),_:1})])),_:1}),v(u,{position:"absolute","content-style":"padding: 1px;",style:{top:"64px",bottom:"64px"}},{default:N((()=>[v(i)])),_:1})])),_:1})};const ke={components:{NResult:I,NButton:a,NCalendar:C,NCard:T},setup(){const e=$();return{success(){e.success("还挺大")},value:J(A(Date.now(),1).valueOf()),handleUpdateValue(t,{year:n,month:o,date:a}){e.success(`${n}-${o}-${a}`)},isDateDisabled:e=>z(e)}}};ke.render=function(e,t,n,o,a,s){const l=f("n-calendar"),r=f("n-card");return y(),b(r,{style:{margin:"0 auto"}},{default:N((()=>[v(l,{"onUpdate:value":[o.handleUpdateValue,t[0]||(t[0]=e=>o.value=e)],value:o.value,"is-date-disabled":o.isDateDisabled},{default:N((({year:e,month:t,date:n})=>[L(O(e)+"-"+O(t)+"-"+O(n),1)])),_:1},8,["onUpdate:value","value","is-date-disabled"])])),_:1})};const we={components:{NResult:I,NButton:a,NCalendar:C},setup(){const e=$();return{success(){e.success("一切尽在不言中")}}}},_e=L("哦");function xe(e){return()=>g(c,null,{default:()=>g(e)})}we.render=function(e,t,n,o,a,s){const l=f("n-button"),r=f("n-result");return y(),b(r,{status:"404",title:"！",description:"接受真相就是这么简单",size:"huge"},{footer:N((()=>[v(l,{onClick:o.success},{default:N((()=>[_e])),_:1},8,["onClick"])])),_:1})};const Me={setup(e,t){const n=function(e,t){if(void 0===t.belongMenu||!0!==t.belongMenu)return[];let o=[];if(!(void 0!==t.children&&t.children.length>0)){let n;return n="/"!==e.charAt(e.length-1)&&"/"!==t.path.charAt(0)?e+"/"+t.path:e+t.path,o.push({label:()=>g(h,{to:{path:n}},{default:()=>t.name}),key:n,icon:xe(P)}),o}for(let a in t.children)o.push(...n(t.path,t.children[a]));return o};let o=[];for(let a=0;a<pt.length;a++)o.push(...n("",pt[a]));o.push({label:()=>g(h,{to:{path:"/moon/moonabout2"}},{default:()=>"回家"}),key:"hear-the-wind-sing",icon:xe(P)},{label:"1973年的弹珠玩具",key:"pinball-1973",icon:xe(P),children:[{label:"鼠",key:"rat"}]});return{isMobile:me(),isTabletRef:ge(),isSmallDesktop:he(),menuOptions:o,options:o,collapsed:J(!0),renderMenuLabel:e=>e.path,renderMenuIcon:e=>"sheep-man"===e.key||("food"===e.key?null:g(c,null)),expandIcon:()=>g(c,null)}},components:{NSpace:n,NLayout:o,NLayoutFooter:_,NLayoutSider:x,NLayoutContent:M,NLayoutHeader:s,NH2:R,NMenu:j,NDropdown:d,NButton:a}},De=L("找个地方休息"),Le=L(" 世事洞明皆学问 | isMobile:"),Ie=L(" |isTabletRef:"),Ce=L(" |isSmallDesktop:");Me.render=function(e,t,n,o,a,s){const l=f("n-button"),r=f("n-dropdown"),i=f("n-layout"),u=f("n-layout-header"),c=f("n-menu"),d=f("n-layout-sider"),p=f("router-view"),m=f("n-layout-footer");return y(),b(i,{position:"absolute"},{default:N((()=>[v(u,{style:{height:"64px",padding:"18px"},bordered:"",position:"absolute"},{default:N((()=>[v(i,{style:{float:"right"}},{default:N((()=>[v(r,{trigger:"hover",options:o.options},{default:N((()=>[v(l,null,{default:N((()=>[De])),_:1})])),_:1},8,["options"])])),_:1})])),_:1}),v(i,{position:"absolute",style:{top:"64px",bottom:"64px"},"has-sider":""},{default:N((()=>[Y(v(d,{"content-style":"padding-top: 24px;","native-scrollbar":!1,bordered:""},{default:N((()=>[v(c,{options:o.menuOptions},null,8,["options"])])),_:1},512),[[q,!o.isMobile]]),v(i,{"content-style":"padding: 24px;","native-scrollbar":!1},{default:N((()=>[v(p)])),_:1})])),_:1}),v(m,{bordered:"",position:"absolute",style:{height:"64px",padding:"24px"}},{default:N((()=>[Le,S("span",null,O(o.isMobile),1),Ie,S("span",null,O(o.isTabletRef),1),Ce,S("span",null,O(o.isSmallDesktop),1)])),_:1})])),_:1})};const Te=e({components:{NDynamicInput:B},setup:()=>({value:J(["","",""])}),methods:{add(e){console.log(e)},remove(e){console.log(e)},update(e){console.log(e)}}});Te.render=function(e,t,n,o,a,s){const l=f("n-dynamic-input");return y(),H(U,null,[v(l,{value:e.value,"onUpdate:value":t[0]||(t[0]=t=>e.value=t),placeholder:"请输入",min:3,max:6,"on-create":e.add,"on-remove":e.remove,"on-update":e.update},null,8,["value","on-create","on-remove","on-update"]),S("pre",null,"  "+O(JSON.stringify(e.value,0,2))+"\r\n  ",1)],64)};var $e=Te;const Je=e({components:{NDynamicInput:B},setup:()=>({sessionStorageData:J(""),localStorageData:J("")}),methods:{set(){let e=localStorage.getItem("tmp");console.log(e),localStorage.setItem("tmp",(Number(e)+1).toString());let t=sessionStorage.getItem("tmp");sessionStorage.setItem("tmp",(Number(t)+1).toString()),this.sessionStorageData=JSON.stringify(sessionStorage),this.localStorageData=JSON.stringify(localStorage)},showNew(e){this.sessionStorageData=JSON.stringify(sessionStorage),this.localStorageData=JSON.stringify(localStorage)}}}),Ae=S("h1",null,"sessionStorage",-1),ze=S("h1",null,"localStorage",-1),Oe=S("h1",null,"cookie",-1),Re=S("h1",null,"session",-1),je=S("h2",null,"tmp in sessionStorage/localStorage 如果存在+1 ",-1);Je.render=function(e,t,n,o,a,s){return y(),H(U,null,[Ae,S("p",null,O(e.sessionStorageData),1),ze,S("p",null,O(e.localStorageData),1),Oe,Re,je,S("button",{onClick:t[0]||(t[0]=(...t)=>e.set&&e.set(...t))}," 设置"),S("button",{onClick:t[1]||(t[1]=(...t)=>e.showNew&&e.showNew(...t))}," 展示最新")],64)};var Pe=Je;const Ye=({sendMail:e})=>[{title:"Name",key:"name"},{title:"Age",key:"age"},{title:"Address",key:"address"},{title:"Tags",key:"tags",render:e=>e.tags.map((e=>g(i,{style:{marginRight:"6px"},type:"info"},{default:()=>e})))},{title:"Action",key:"actions",render:t=>g(a,{size:"small",onClick:()=>e(t)},{default:()=>"Send Email"})}],qe=e({components:{NDataTable:D},setup(){const e=$();return{data:[{key:0,name:"John Brown",age:32,address:"New York No. 1 Lake Park",tags:["nice","developer"]},{key:1,name:"Jim Green",age:42,address:"London No. 1 Lake Park",tags:["wow"]},{key:2,name:"Joe Black",age:32,address:"Sidney No. 1 Lake Park",tags:["cool","teacher"]}],columns:Ye({sendMail(t){e.info("send mail to "+t.name)}}),pagination:{pageSize:10}}}});qe.render=function(e,t,n,o,a,s){const l=f("n-data-table");return y(),b(l,{columns:e.columns,data:e.data,pagination:e.pagination},null,8,["columns","data","pagination"])};var Be=qe;const He={components:{NCarousel:E},setup(){}},Ue=S("img",{class:"carousel-img",src:"https://s.anw.red/fav/1623979004.jpg!/fw/600/quality/77/ignore-error/true"},null,-1),Ee=S("img",{class:"carousel-img",src:"https://s.anw.red/news/1623372884.jpg!/both/800x450/quality/78/progressive/true/ignore-error/true"},null,-1),Fe=S("img",{class:"carousel-img",src:"https://s.anw.red/news/1623177220.jpg!/both/800x450/quality/78/progressive/true/ignore-error/true"},null,-1),Ve=S("img",{class:"carousel-img",src:"https://s.anw.red/news/1623152423.jpg!/both/800x450/quality/78/progressive/true/ignore-error/true"},null,-1);He.render=function(e,t,n,o,a,s){const l=f("n-carousel");return y(),b(l,{autoplay:""},{default:N((()=>[Ue,Ee,Fe,Ve])),_:1})};var Ze=He;const Ge=L("哦");var Ke={setup:e=>(e,t)=>(y(),b(F(I),{status:"404",title:"！",description:"这么大",size:"huge"},{footer:N((()=>[v(F(a),null,{default:N((()=>[Ge])),_:1})])),_:1}))};const Qe=L("看看别的");var We={setup:e=>(e,t)=>(y(),b(F(V),{size:"large",description:"可以是大的"},{extra:N((()=>[v(F(a),{size:"small"},{default:N((()=>[Qe])),_:1})])),_:1}))};var Xe={setup:e=>(e,t)=>(y(),b(F(Z),null,{default:N((()=>[v(F(n),null,{default:N((()=>[v(F(G),{width:"100",src:"https://07akioni.oss-cn-beijing.aliyuncs.com/07akioni.jpeg"}),v(F(G),{width:"100",src:"https://gw.alipayobjects.com/zos/antfincdn/aPkFc8Sj7n/method-draw-image.svg"})])),_:1})])),_:1}))};const et={components:{NRow:K,NCol:Q,NStatistic:W,NIcon:c}},tt=L("/ 100"),nt=L("1,234,123");et.render=function(e,t,n,o,a,s){const l=f("n-icon"),r=f("n-statistic"),i=f("n-col"),u=f("n-row");return y(),b(u,null,{default:N((()=>[v(i,{span:12},{default:N((()=>[v(r,{label:"统计数据",value:99},{prefix:N((()=>[v(l)])),suffix:N((()=>[tt])),_:1})])),_:1}),v(i,{span:12},{default:N((()=>[v(r,{label:"活跃用户"},{default:N((()=>[nt])),_:1})])),_:1})])),_:1})};var ot=et;const at=e({name:"",props:{},components:{NRow:K,NCol:Q,NStatistic:W,NIcon:c},mounted(){this.$formula(document.getElementById("katexData"))},setup(){const e=new X;e.use((function(e,t){console.log(e),console.log(t)}));return{result:e.render("\n# markdown-it test!\n\n## two\n\n### three\n\n#### four\n\n##### five\n\n\n```sql\ninsert into dogs values(1,2,3);\n```\n- [ ] 111\n- [ ] 112\n\n[baidu](https://www.baidu.com)\n\n|12|12|\n|---|---|\n|1 1| 112|\n\n$\\sqrt{3x-1}+(1+x)^2$\n\n    ")}}}),st=["innerHTML"],lt=S("div",{class:"mermaid"}," sequenceDiagram Alice->>+John: Hello John, how are you? Alice->>+John: John, can you hear me? John--\x3e>-Alice: Hi Alice, I can hear you! John--\x3e>-Alice: I feel great! ",-1),rt=S("p",{id:"katexData"}," $\\int$ ",-1);at.render=function(e,t,n,o,a,s){return y(),H(U,null,[S("div",{innerHTML:e.result},null,8,st),lt,rt],64)};var it=at;const ut=e({components:{NDynamicInput:B,NTree:ee,NInput:te,NSpace:n},setup(){let e=[];for(let t=1;t<=70;t++){let n={key:t.toString(),label:"deep1"+t,children:[]};for(let e=1;e<=10;e++){let o={key:t.toString()+":"+e.toString(),label:"deep2"+e,children:[]};for(let n=1;n<=2;n++){let a={key:t.toString()+":"+e.toString()+":"+n.toString(),label:"deep3"+n,children:[]};o.children.push(a)}n.children.push(o)}e.push(n)}return{data:e,pattern:J("")}}});ut.render=function(e,t,n,o,a,s){const l=f("n-input"),r=f("n-tree"),i=f("n-space");return y(),b(i,{vertical:"",size:12},{default:N((()=>[v(l,{value:e.pattern,"onUpdate:value":t[0]||(t[0]=t=>e.pattern=t),placeholder:"搜索"},null,8,["value"]),v(r,{pattern:e.pattern,data:e.data,"block-line":""},null,8,["pattern","data"])])),_:1})};var ct=ut;const dt={components:{NResult:I,NButton:a,NCalendar:C,NCard:T,NTimelineItem:ne,NTimeline:oe},setup(){const e=$();let t=[],n=ae(),o=ae(ae().format("YYYY-01-01"));for(let a=1;a<12;a++){o.add(1,"months");let e="warning",a="dashed";console.log(o.format("M"),n.format("M")),parseInt(o.format("M"))>parseInt(n.format("M"))&&(e="success",a="default");let s=o.format("YYYY-MM-DD");t.push({title:s,time:s,type:e,lineType:a})}return t.sort((function(e,t){return e.time>t.time?-1:1})),t.push({title:"start"}),t.unshift({title:"end",type:"success"}),{dayInfoList:t,success(){e.success("还挺大")}}}};dt.render=function(e,t,n,o,a,s){const l=f("n-timeline-item"),r=f("n-timeline"),i=f("n-card");return y(),b(i,{style:{margin:"0 auto"}},{default:N((()=>[v(r,{size:"large"},{default:N((()=>[(y(!0),H(U,null,se(o.dayInfoList,(e=>(y(),b(l,{type:e.type,title:e.title,content:e.content,time:e.time,"line-type":e.lineType},null,8,["type","title","content","time","line-type"])))),256))])),_:1})])),_:1})};var pt=[{path:"/:catchAll(.*)*",name:"",redirect:"/sun/index"},{path:"/sun",component:fe,children:[{name:"",path:"",component:ke},{name:"index",path:"index",component:dt},{name:"sunMain",path:"sunMain",component:ke},{name:"about",path:"about",component:we}]},{belongMenu:!0,path:"/moon",component:Me,children:[{name:"",path:"",component:Be,belongMenu:!1},{name:"moonabout1",path:"moonabout1",component:Be,belongMenu:!0},{name:"moonabout2",path:"moonabout2",component:Ze,belongMenu:!0},{name:"moonabout3",path:"moonabout3",component:Ke,belongMenu:!0},{name:"moonabout4",path:"moonabout4",component:We,belongMenu:!0},{name:"moonabout5",path:"moonabout5",component:Xe,belongMenu:!0},{name:"moonabout6",path:"moonabout6",component:ot,belongMenu:!0},{name:"markdown",path:"markdown",component:it,belongMenu:!0},{name:"listTodo",path:"listTodo",component:$e,belongMenu:!0},{name:"tmp",path:"tmp",component:Pe,belongMenu:!0},{name:"tree",path:"tree",component:ct,belongMenu:!0}]}];const mt={delimiters:[{left:"$$",right:"$$",display:!0},{left:"$",right:"$",display:!1},{left:"\\(",right:"\\)",display:!1},{left:"\\[",right:"\\]",display:!0}],throwOnError:!1},gt=le({history:re(),routes:pt}),ht=ie(de);ht.config.globalProperties.$axios=ue,ht.config.globalProperties.$formula=function(e){ce(e,mt)},ht.use(gt),ht.mount("#app");
