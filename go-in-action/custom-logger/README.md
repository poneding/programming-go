# custom-logger

自定义日志

## 日志级别

- TRACE
- DEBUG
- INFO
- WARN
- ERROR
- FATAL

## 日志格式

`[level] yyyy/mm/dd hh:mm:ss [file:line] message`

例如：

```txt
[INFO] 2023/07/04 09:28:13 cmd/main.go: 6  [Hello Jay Chou and Miachael Jackson]
```

## 写日志文件

```go
cuzlog.SetFile("test.log")

cuzlog.Infoln("Hello World")
```

## 设置日志级别

```go
cuzlog.SetLevel(cuzlog.WARN)

cuzlog.Infoln("Hello Info")
cuzlog.Warnln("Hello Warn")
```