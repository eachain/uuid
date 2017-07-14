# uuid

一个简单粗暴，模拟uuid格式，并且不暴露物理机器信息的uuid生成库。

## Usage:

```go
package main

import (
	"github.com/eachain/uuid"
)

func main() {
	println(uuid.New())
}
```

