## 返回值
- Go 中 main 函数不不⽀支持任何返回值
- 通过 os.Exit 来返回状态
## 获取命令⾏行行参数
- main 函数不不⽀支持传⼊入参数
- main 函数不不⽀支持传⼊入参数 func main(arg []string)
- 在程序中直接通过 os.Args 获取命令⾏行行参数

举例
```go
package main

import (
	"fmt"
	"os"
)

func main()  {
	if len(os.Args) > 0{
		fmt.Println("hello world",os.Args[1])
	}
	os.Exit(2)
}
```
执行结果
```shell
go run  main.go jianxiang
hello world jianxiang
exit status 2

```
