# readme.md

- [ ] 简化 如类型转换，废弃helper 采用范围更大的 cast
- [ ] 优化启动逻辑，避免在全逻辑命令下的的无意义初始化
  - 方案一 逻辑懒加载
  - 方案二 入口分离
  - 方案三 抽离全逻辑，创建新项目

## 一个简易的gin项目文件

参考了 `laravel` 目录

```
├── app
│   ├── http
│   │   └── controllers
│   │       ├── api
│   │       │   └── about_controller.go
│   │       ├── base_controller.go     
│   │       └── web
│   │           └── welcome_controller.go
│   └── middlewares
│       └── simple_logger.go
├── bootstrap
│   └── bootstrap.go
├── config
│   ├── app.go
│   ├── config.go
│   └── database.go
├── go.mod
├── go.sum
├── main.go
├── models
│   ├── db.go
│   ├── model.go
│   └── user
│       └── user.go
├── readme.md
├── resources
│   └── views
│       └── welcome.tmpl
├── routes
│   ├── api.go
│   ├── route.go
│   └── web.go
├── tmp
│   ├── build-errors.log
│   ├── main.exe
│   ├── runner-build.exe
│   ├── runner-build.exe~
│   └── tmp.db
└── util
    └── types
        └── types.go
```

# how to run 

## ~~如果需要生成静态文件~~
```
statik -f -src=web/  -include=*.jpg,*.txt,*.html,*.css,*.js 
```

如果你想使用热加载进行开发。

[air](https://github.com/cosmtrek/air)

> [realize](https://github.com/oxequa/realize)
> [fresh](https://github.com/gravityblast/fresh)

```
go get -u github.com/cosmtrek/air
air 
realize start # 这个过于老
fresh # 这个也过于老
go mod tidy            
```



```other
https://aip.baidubce.com
```


```text
go  test  ./...   
// array_test.go 依赖 array.go 和 types.go
go test ./helpers/array_test.go ./helpers/array.go ./helpers/types.go 
```

```shell
cobra add newCommond
```

```text
node node_modules/esbuild/install.js  
npm run dev  -- --host 0.0.0.0

git config user.name 'github用户名'  
git config user.email '邮箱'  

gofmt -w .
```



```
go build -ldflags="-w -s" .
```
