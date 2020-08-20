## Go学习实践

本项目作为`Go`入门的实践项目, 将基于`Gin`框架完成一些封装. 最终目标实现一个`配置化的`,`能够快速上手的`, `规范目录结构的`脚手架.

1. Go编程基础(视频) [这里](https://study.163.com/course/courseMain.htm?courseId=306002)
2. GOWeb基础(视频) [这里](https://study.163.com/course/courseMain.htm?courseId=328001)
3. 入门指南(文档) [这里](https://github.com/unknwon/the-way-to-go_ZH_CN)
4. Go语言圣经(文档) [这里](https://github.com/golang-china/gopl-zh)

### 目录结构定义
```bash
├── controler 控制器
├── config 项目配置 
├── module 数据模型
├── pkg  基础类库
├── main.go
```

#### Go的安装 

[查看](https://www.jianshu.com/p/ad57228c6e6a)

#### 基础参考
- Gin框架 [Gin](https://github.com/gin-gonic/gin) [GinDoc](https://learnku.com/docs/gin-gonic/2019)

```bash
go get -u github.com/gin-gonic/gin
```

- 更详细的变量打印 [Spew](https://github.com/davecgh/go-spew)

```bash
go get -u github.com/davecgh/go-spew/spew
```

- 实时编译, 当代码变更时自动重启服务 [fresh](https://github.com/gravityblast/fresh)

```bash
go get -u github.com/pilu/fresh
```

### 实践

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
- [ ] pkg/db/model.go crud 方法封装 

### 数据验证

- [ ] 验证器
- [ ] 接口参数合法性自动验证

### 权限验证
- [ ] JWT
- [ ] AKSK

### 参考资料
[database/sql资料](https://segmentfault.com/a/1190000003036452)
