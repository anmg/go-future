package select_test

import (
    "fmt"
    "testing"
    "time"
)
func Chann(ch chan int, stopCh chan bool) {
    var i int
    i = 10
    for j := 0; j < 10; j++ {
        ch <- i
        time.Sleep(time.Second)
    }
    stopCh <- true
}

func TestSelect(t *testing.T)  {
    ch := make(chan int)
    c := 0
    stopCh := make(chan bool)

    go Chann(ch, stopCh)

    for {
        select {
        case c = <-ch:
            fmt.Println("Recvice", c)
            fmt.Println("channel")
        case s := <-ch:
            fmt.Println("Receive", s)
        case _ = <-stopCh:
            goto end
        }
    }
    end:
}


/*
type GetNewPassenger struct{
}

// TickerLaunch用于演示ticker聚合任务用法
func TickerLaunch() {
    ticker := time.NewTicker(5 * time.Minute)
    maxPassenger := 30                   // 每车最大装载人数
    passengers := make([]string, 0, maxPassenger)

    for {
        passenger := GetNewPassenger() // 获取一个新乘客
        if passenger != "" {
            passengers = append(passengers, passenger)
        } else {
            time.Sleep(1 * time.Second)
        }

        select {
        case <- ticker.C:               // 时间到，发车
            Launch(passengers)
            passengers = []string{}
        default:
            if len(passengers) >= maxPassenger {  // 时间没到，车已座满，发车
                Launch(passengers)
                passengers = []string{}
            }
        }
    }
}
*/
