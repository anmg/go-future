package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
func main() {
	http.HandleFunc("/hello", hello_http.HelloServer)
	http.HandleFunc("/time", hello_http.TimeServer)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
*/

//RESTful路由get函数、
func helloworldGet(c *gin.Context){
	c.String(http.StatusOK, "Hello world in Get")
}

func main() {
	router := gin.Default()

	//RESTful路由
	router.GET("/helloworld", helloworldGet)

	//服务启动
	router.Run("127.0.0.1:8082")
}