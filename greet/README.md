### type用法和go一致，service用来定义get/post/head/delete等api请求，解释如下：

service bookstore-api 这一行定义了service名字
@handler定义了服务端handler名字
get /add(addReq) returns(addResp)定义了get方法的路由、请求参数、返回参数等
使用goctl生成API Gateway代码


```
goctl api go -api bookstore.api -dir .
```
生成的文件结构如下：

```
api
├── bookstore.api                  // api定义
├── bookstore.go                   // main入口定义
├── etc
│   └── bookstore-api.yaml         // 配置文件
└── internal
    ├── config
    │   └── config.go              // 定义配置
    ├── handler
    │   ├── addhandler.go          // 实现addHandler
    │   ├── checkhandler.go        // 实现checkHandler
    │   └── routes.go              // 定义路由处理
    ├── logic
    │   ├── addlogic.go            // 实现AddLogic
    │   └── checklogic.go          // 实现CheckLogic
    ├── svc
    │   └── servicecontext.go      // 定义ServiceContext
    └── types
        └── types.go               // 定义请求、返回结构体
```
在 api 目录下启动API Gateway服务，默认侦听在8888端口
```
go run bookstore.go -f etc/bookstore-api.yaml
```
测试API Gateway服务
```
curl -i "http://localhost:8888/check?book=go-zero"
```
返回如下：
```
HTTP/1.1 200 OK
Content-Type: application/json
Date: Thu, 03 Sep 2020 06:46:18 GMT
Content-Length: 25

{"found":false,"price":0}
```

### grpc 自动生成

```
goctl rpc protoc add.proto --go_out=./pb --go-grpc_out=./pb --zrpc_out=.
```

#### 直接运行
```
$ go run greet.go -f etc/greet-api.yaml
  Starting rpc server at 127.0.0.1:8080...
  ```

### api 自动生成

```
goctl api go -api bookstore.api -dir .
```

### docker 生成
```
goctl docker -go greet.go
```


### gen 自动生成
### 注意修改里面的表名
```
cd cmd
go run gen.go
```
