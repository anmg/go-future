package groutine_test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestGroutine(t *testing.T){
	for i:=0; i< 10; i++{
		go func (i int) {
			fmt.Println(i)
		}(i)
	}
	time.Sleep(time.Microsecond * 50)
}

func TestCounter(t *testing.T){
	counter := 0
	for i:=0; i< 50000; i++{
		go func() {
			counter ++
			//fmt.Println(counter)
		}()
	}

	//time.Sleep(1 * time.Second)
	t.Logf("counter = %d", counter)
}

func TestCounterSafe(t *testing.T){
	var mut sync.Mutex
	counter := 0
	for i:=0; i< 50000; i++{
		go func() {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			counter ++
			//fmt.Println(counter)
		}()
	}

	//time.Sleep(1 * time.Second)
	t.Logf("counter = %d", counter)
}

func TestCounterWaitGroup(t *testing.T){
	var mut sync.Mutex
	var wg sync.WaitGroup
	counter := 0
	for i:=0; i< 50000; i++{
		wg.Add(1)
		go func() {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			counter ++
			wg.Done()
			//fmt.Println(counter)
		}()
	}

	wg.Wait()
	//time.Sleep(1 * time.Second)
	t.Logf("counter = %d", counter)
}


func TestCounterRWWaitGroup(t *testing.T){
	var mut sync.RWMutex
	var wg sync.WaitGroup

	counter := 0
	for i:=0; i< 50000; i++{
		wg.Add(1)
		go func() {
			defer func() {
				//mut.Unlock()
				mut.RUnlock()
			}()
			//mut.Lock()
			mut.RLock()
			counter ++
			wg.Done()
			//fmt.Println(counter)
		}()
	}

	wg.Wait()
	//time.Sleep(1 * time.Second)
	t.Logf("counter = %d", counter)
}