# im-system

即时通信系统：Instant Messaging System

![202306291126761](https://pding.oss-cn-hangzhou.aliyuncs.com/images/202306291126761.png)


## 启动服务端

当前目录打开一个终端窗口，运行服务端：

```bash
go run cmd/server/main.go
```

nc 测试

再打开一个终端窗口，运行客户端 1：

```bash
nc localhost 8081
```

再打开一个终端窗口，运行客户端 2：

```bash
nc localhost 8081
```

客户端命令：

```bash
who
rename dp

# chat
@jay hello
```

## 客户端

```bash
go run cmd/client/main.go
```