# github.com/google/uuid

[`google/uuid`](https://pkg.go.dev/github.com/google/uuid) 生成和解析 UUID。

## 引入

```go
import "github.com/google/uuid"
```

## 安装

```bash
go get github.com/google/uuid
```

## 示例

文件：*[main.go](./main.go)*

```go
package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	id := uuid.New()
	fmt.Printf("id: %v\n", id)

	idString := uuid.NewString() // 相当于 uuid.New().String()
	fmt.Printf("idString: %v\n", idString)

	ramUUID := "d68c863a-28f9-4204-ac67-db96f575bf85"
	parsedUUID := uuid.MustParse(ramUUID).String()

	fmt.Printf("parsedUUID: %v\n", parsedUUID)
}
```

## 运行

```bash
go run main.go
```
