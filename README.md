## Go学习实践

本项目作为`Go`入门的实践项目, 将基于`Gin`框架完成一些封装. 最终目标实现一个`配置化的`,`能够快速上手的`, `规范目录结构的`脚手架.

1. GO安装 [这里](https://www.jianshu.com/p/ad57228c6e6a)
2. Go编程基础(视频) [这里](https://study.163.com/course/courseMain.htm?courseId=306002)
3. GOWeb基础(视频) [这里](https://study.163.com/course/courseMain.htm?courseId=328001)
4. 入门指南(文档) [这里](https://github.com/unknwon/the-way-to-go_ZH_CN)
5. Go语言圣经(文档) [这里](https://github.com/golang-china/gopl-zh)

### 目录结构定义

```bash
./
├── README.md
├── app.json 运行时配置
├── config 启动性配置
│   ├── middlewares.go
│   ├── routes
│   │   └── user.go
│   └── routes.go
├── consumer 消费者
│   └── rabbmitmq_exmaple.go mq样例
├── controller 控制器
│   └── user.go
├── docs 文档
│   ├── consul.md
│   └── grafana-prometheus.md
├── go.mod 类目 composer.json
├── go.sum 类似 composer.lock
├── main.go 项目入口
├── model 数据模型
│   ├── common_config.go
│   └── user.go
├── pkg 脚手架-业务无关
│   ├── bootstrap.go 启动入口
│   ├── cache 缓存
│   │   ├── db.go
│   │   ├── hook.go
│   │   └── redis.go
│   ├── consts 常量
│   │   └── code.go 系统错误码及信息
│   ├── db mysql-model
│   │   ├── db.go db链接
│   │   ├── model.go 模型 crud
│   │   ├── sql.go sql构造
│   │   └── sql_test.go
│   ├── lib 通用类库
│   │   ├── func.go
│   │   ├── http.go
│   │   └── string.go
│   ├── middleware gin中间件
│   │   ├── auth.go 鉴权中间件 IpAuth/AKSK...
│   │   ├── jwt.go
│   │   ├── logger.go http请求日志
│   │   └── prometheus.go http请求打点
│   ├── route 路由处理
│   │   └── route.go
│   └── utils 辅助工具
│       ├── config.go 配置信息 local + app.json + consul
│       ├── consul.go consul client
│       ├── jwt.go
│       ├── logger.go 日志
│       ├── rabbitmq.go mq常用方法
│       └── validator.go 验证器的初始化
├── service 业务服务
```

#### Go的安装 

[查看](https://www.jianshu.com/p/ad57228c6e6a)

### 实践

我们将在一个通用脚手架的封装过程中不断学习GO语言.

#### 路由的配置化注册

我们将在 `conf/routes.go` 中集中定义项目路由, 项目启动时统一注册到`Gin`框架中

- [x] conf/routes.go 基础路由定义
- [x] conf/routes/**.go 多文件定义

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

- [x] 验证器
- [x] 接口参数合法性自动验证

### 权限验证

- [x] JWT
- [ ] AKSK

### 缓存

- [x] redis的基础封装
- [ ] 全量方法完善

### 健康控制

- [ ] 接口频率控制
    - [x] 基于滑动窗口的频率限制
    - [x] ip维度
    - [ ] 用户维度
- [ ] prometheus打点
    - [x] api打点
    - [ ] db数据打点
    - [x] redis打点

### 微服务

- [x] grpc

### 其他

- [x] consul
- [x] 基于consul的配置中心
- [ ] 开关服务
- [ ] 规则引擎
- [x] RabbitMQ
- [ ] kafka
- [ ] nsq
- [x] 事件总线
- [ ] 配置解密
- [ ] swagger

### 参考资料
- [GoProxy解决安装慢](https://goproxy.cn/)
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
- [air热重启-推荐](https://github.com/cosmtrek/air)
- [logrus日志](https://juejin.im/post/6844904061393698823)
- [如何避免用动态语言的思维写Go代码](https://juejin.im/post/6861048173989724173)
- [Gin中间件详解](https://juejin.im/post/6844903833164857358)
- [数据验证](https://segmentfault.com/a/1190000022541905)
- [validator](https://frankhitman.github.io/zh-CN/gin-validator/)
- [Go 语言设计与实现](https://draveness.me/golang/docs/part1-prerequisite/ch02-compile/golang-compile-intro/)
- [Go by Example](https://gobyexample.com/)
- [mac/ab压测](https://xushanxiang.com/2019/10/mac-web-ab.html)
- [prometheus](https://yunlzheng.gitbook.io/prometheus-book/)
- [macRabbitMQ](https://www.jianshu.com/p/60c358235705)

### 常用命令

以下命令均在项目根目录执行

```bash
go run main.go
go test -v ./pkg/db
godoc -http=:8888
air
```

### 如何参与

```bash
go get github.com/daodao97/egin
cd $GOPATH/src/github.com/daodao97/egin
goland . #编辑器打开  vscode .
air #开发模式启动
```

### 规划

当功能逐步完善稳定后 `pkg` 目录下的公用库将独立成为一个 `go package`, 方便在任意项目中复用

### 常见问题

-  dial tcp 127.0.0.1:8080: socket: too many open files

    [永久修改 mac mac files/proc 限制](https://javasgl.github.io/mac-max-limit/)
 
 - db查询中的内存泄露问题
 
   [databases/sql内存泄露](https://gocn.vip/topics/9963)





