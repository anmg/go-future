package hello_http

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, world!\n")
}

func TimeServer(w http.ResponseWriter, req *http.Request) {
	t := time.Now()
	timestr := fmt.Sprintf("{\"time\":\"%s\"}",t)
	w.Write([]byte(timestr))
}

//func main() {
//	http.HandleFunc("/hello", HelloServer)
//	http.HandleFunc("/hello", TimeServer)
//	err := http.ListenAndServe(":8080", nil)
//	if err != nil {
//		log.Fatal("ListenAndServe: ", err)
//	}
//}

