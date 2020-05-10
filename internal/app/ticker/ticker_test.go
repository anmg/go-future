package ticker

import (
    "log"
    "testing"
    "time"
)

// TickerDemo 用于演示ticker基础用法
func TestTickerDemo(t *testing.T) {
    ticker := time.NewTicker(1 * time.Second)
    defer ticker.Stop()

    for range ticker.C {
        log.Println("Ticker tick.")
    }
}

// TickerLaunch用于演示ticker聚合任务用法
func TestTickerLaunch(t *testing.T) {
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