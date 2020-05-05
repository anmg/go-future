package singleton

import (
	"fmt"
	"sync"
	"testing"
)

type Singleton struct {
	Name string
	age int
}

var singleInstance *Singleton
var once sync.Once

func GetSingletonObj() *Singleton{
	once.Do(func() {
		fmt.Println("Creat obj")
		singleInstance = new(Singleton)
		singleInstance.Name = "anmg"
		singleInstance.age = 30
	})
	return  singleInstance
}

func TestGetSingletonObj(t *testing.T){
	wg := sync.WaitGroup{}
	for i:=0; i < 10; i++ {
		wg.Add(1)
		go func() {
			obj := GetSingletonObj()
			fmt.Println(obj.Name, obj.age)
			wg.Done()
		}()
	}
	wg.Wait()
}


type Config struct {

}

var cfg *Config

func getInstance() *Config {
	if cfg == nil {
		cfg = new(Config)
	}
	return cfg
}

