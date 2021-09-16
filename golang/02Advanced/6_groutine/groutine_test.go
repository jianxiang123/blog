package __groutine

import (
	"fmt"
	"sync"
	"testing"
)

func TestGroutine(t *testing.T)  {
	var wg sync.WaitGroup
	for i:=0;i<8;i++{
		wg.Add(1)
		go func(i int) {
			fmt.Println(i)
			wg.Done()
		}(i)
	}
	wg.Wait()  //使用WaitGroup 完成后自动结束
}