package __channel_csp

import (
	"fmt"
	"testing"
	"time"
)

func service1() string {
	time.Sleep(time.Millisecond * 50)
	return "Done"
}

func otherTask() {
	fmt.Println("working on something else")
	time.Sleep(time.Millisecond * 100)
	fmt.Println("Task is done.")
}

func TestService(t *testing.T) {
	fmt.Println(service())
	otherTask()
}

func AsyncService1() chan string {
	retCh := make(chan string, 1)
	//retCh := make(chan string, 1)
	go func() {
		ret := service1()
		fmt.Println("returned result.")
		retCh <- ret
		fmt.Println("service exited.")
	}()
	return retCh
}

//
func TestAsynService(t *testing.T) {
	retCh := AsyncService1()
	otherTask()
	fmt.Println(<-retCh)
	time.Sleep(time.Second * 1)
}
