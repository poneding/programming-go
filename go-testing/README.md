# go-testing

## 规则

- 测试文件必须以 `_test.go` 结尾
- 单元测试函数必须以 `Test` 开头
- 基准测试函数必须以 `Benchmark` 开头
- 模糊测试函数必须以 `Fuzz` 开头
- 使用标准库 `testing` 包
- 文件夹以 `_test` 结尾，会被便衣成分离包
- 使用 `go help test` 或 `go help testflag` 查看测试命令帮助
- 测试包名可以保持和被测试包一致，也可以用被测试包名 + `_test` 结尾

## go test 模式

### 本地模式

`go test` 或 `go test -v`

### 列表模式

编译并运行命令行上列出的每个包中的测试

`go test util`、`go test ./...`
