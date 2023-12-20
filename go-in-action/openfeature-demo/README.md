# openfeature go demo

1. 运行 flagd：

```bash
docker run -p 8013:8013 -v $(pwd)/:/etc/flagd/ -it ghcr.io/open-feature/flagd:latest start --uri file:/etc/flagd/flags.flagd.json
```

2. 运行 demo：

```bash
go mod tidy
go run main.go
```

3. 访问 <http://localhost:8080/hello>
4. 修改 `flags.flagd.json` 文件 `defaultVariant` 的值，观察 <http://localhost:8080/hello> 的返回值变化
