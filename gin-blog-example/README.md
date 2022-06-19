##  一个基于gin框架的demo项目，参考于[Gin搭建Blog Api's](https://eddycjy.com/posts/go/gin/2018-02-11-api-01/)

### 1. 目录结构解析

```shell
go-gin-example/
├── conf
├── middleware
├── models
├── pkg
├── routers
└── runtime
```

- conf : 存储配置文件
- middleware : 应用中间件
- models : 应用数据库模型-orm映射模型
- pkg : 第三方包
- routers : 路由逻辑处理
- runtime : 应用运行时数据