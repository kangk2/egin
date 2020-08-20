## Go学习实践

本项目作为`Go`入门的实践项目, 将基于`Gin`框架完成一些封装. 最终目标实现一个`配置化的`,`能够快速上手的`, `规范目录结构的`脚手架.

1. GO安排 [这里](https://www.jianshu.com/p/ad57228c6e6a)
2. Go编程基础(视频) [这里](https://study.163.com/course/courseMain.htm?courseId=306002)
3. GOWeb基础(视频) [这里](https://study.163.com/course/courseMain.htm?courseId=328001)
4. 入门指南(文档) [这里](https://github.com/unknwon/the-way-to-go_ZH_CN)
5. Go语言圣经(文档) [这里](https://github.com/golang-china/gopl-zh)

### 目录结构定义
```bash
./
├── README.md
├── app.json 配置文件
├── config 配置目录
│   └── routes.go 路由配置
├── controller 控制器
│   └── user.go
├── go.mod 项目依赖, 类似 php composer.json
├── go.sum 依赖的锁定, 类似 php composer.lock
├── main.go 入口文件, 类似 php index.php
├── model 数据库模型
│   └── user.go 对应数据源 user 表
├── pkg 脚手架类库
│   ├── bootstrap.go 
│   ├── consts 常量
│   │   └── code.go 系统CODE码及信息
│   ├── db 数据库访问层
│   │   ├── db.go db连接相关
│   │   ├── model.go 模型 数据操作 CRUD
│   │   ├── sql.go sql构造
│   │   └── sql_test.go sql构造的测试文件
│   ├── lib 统一类库
│   │   ├── func.go
│   │   ├── http.go
│   │   └── string.go
│   ├── middleware http访问的中间件
│   │   └── logger.go
│   ├── route 路由解析注册
│   │   └── route.go
│   └── utils 通用工具
│       ├── config.go
│       └── logger.go
├── service 业务服务
```

#### Go的安装 

[查看](https://www.jianshu.com/p/ad57228c6e6a)

### 实践

我们将在一个通用脚手架的封装过程中不断学习GO语言.

#### 路由的配置化注册

我们将在 `conf/routes.go` 中集中定义项目路由, 项目启动时统一注册到`Gin`框架中

- [x] conf/routes.go 基础路由定义
- [ ] conf/routes/**.go 多文件定义

### 配置的管理

- [x] pkg/utils/config.go env,json等配置方式的支持

### 日志

- [x] pkg/utils/logger.go 基于logrus的日志管理, 标准输出/文件(分割)
- [ ] 日志输出到 es/mongo/redis 

### 控制器的封装

基于单一数据模型的`CRUD`通用控制器方法

- [ ] 底层通用方法

### 数据模型的封装

基于单表的`CRUD`方法

- [x] pkg/db/model.go 连接池 
- [x] pkg/db/model.go crud 方法封装 

### 数据验证

- [ ] 验证器
- [ ] 接口参数合法性自动验证

### 权限验证
- [ ] JWT
- [ ] AKSK

### 参考资料
- [database/sql资料](https://segmentfault.com/a/1190000003036452)
- [go/tools-github](https://github.com/golang/tools)
- [go/tools简明介绍](https://studygolang.com/articles/11837)
- [go/test详解](http://c.biancheng.net/view/124.html)
- [go/mod包管理](https://juejin.im/post/6844903798658301960)
- [go/doc项目文档](https://wiki.jikexueyuan.com/project/go-command-tutorial/0.5.html)
- [go/generate代码生成](https://juejin.im/post/6844903923166216200)
- [gin文档](https://learnku.com/docs/gin-gonic/2019)
- [Gin-github](https://github.com/gin-gonic/gin)
- [gorm文档](http://gorm.io/zh_CN/docs/index.html)
- [Spew变量打印](https://github.com/davecgh/go-spew)
- [fresh项目热重启](https://github.com/gravityblast/fresh)
- [logrus日志](https://juejin.im/post/6844904061393698823)

### 常用命令

一下命令均在项目根目录执行

```bash
go run main.go
go test -v ./pkg/db
godoc -http=:8888
fresh
```


