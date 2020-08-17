##Go学习实践

本项目作为`Go`入门的实践项目, 将基于`Gin`框架完成一些封装. 最终目标实现一个`配置化的`,`能够快速上手的`, `规范目录结构的`脚手架.

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

### 控制器的封装

基于单一数据模型的`CRUD`通用控制器方法

### 数据模型的封装

基于单表的`CRUD`方法
