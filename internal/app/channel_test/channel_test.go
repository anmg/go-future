package channel_test

import (
	"fmt"
	"testing"
	"time"
)

func service() string {
	time.Sleep(time.Microsecond * 50)
	return "-2-share str"
}

func otherTask()  {
	fmt.Println("working on otherTask")
	time.Sleep(time.Microsecond* 100)
	fmt.Println("otherTask is Done")
}

func TestService(t *testing.T){
	fmt.Println(service())
	otherTask()
}

func AsyncService() chan string{
	retCh := make(chan string)
	//retCh := make(chan string, 1)
	go func() {
		fmt.Println("-1-service start")
		ret := service()
		fmt.Println("-3-service end")
		retCh <- ret
		fmt.Println("-4-channel set share")
	}()
	return retCh
}
func TestAsyncService(t *testing.T){
	retCh := AsyncService()
	otherTask()
	fmt.Println(<-retCh)
}

func addNumberToChan(chanName chan int) {
	for {
		chanName <- 1
		time.Sleep(1 * time.Second)
	}
}

func TestAddNumber(t *testing.T) {
	var chan1 = make(chan int, 10)
	var chan2 = make(chan int, 10)

	go addNumberToChan(chan1)
	go addNumberToChan(chan2)

	for {
		select {
		case e := <- chan1 :
			fmt.Printf("Get element from chan1: %d\n", e)
		case e := <- chan2 :
			fmt.Printf("Get element from chan2: %d\n", e)
		default:
			fmt.Printf("No element in chan1 and chan2.\n")
			time.Sleep(1 * time.Second)
		}
	}
}
