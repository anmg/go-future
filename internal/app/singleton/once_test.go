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
		//fmt.Println("Creat obj")
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
	Name string
	age int
}
var cfg2 *Config

func init(){
	cfg2 = new(Config)
}

func NewConfig() *Config {
	return cfg2
}

var cfg3 *Config = new(Config)

var cfg *Config
var once2 sync.Once
func getInstance() *Config {
	once2.Do(func() {
		cfg = new(Config)
		cfg.Name = "anmg"
		fmt.Println("Creat Config",cfg)
	})
	return cfg
}

func TestGetSingletonConfig(t *testing.T){
	wg := sync.WaitGroup{}
	for i:=0; i < 100; i++ {
		wg.Add(1)
		go func() {
			obj := getInstance()
			fmt.Println(obj)
			wg.Done()
		}()
	}
	wg.Wait()
}
