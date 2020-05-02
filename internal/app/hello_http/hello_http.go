package go_work

import (
	"fmt"
	"net/http"
	"time"
)

func main(){
	http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "hello word")
	})
	http.HandleFunc("/time",func(w http.ResponseWriter, r *http.Request){
		t := time.Now()
		timestr := fmt.Sprintf("{\"time\":\"%s\"}",t)
		w.Write([]byte(timestr))
		fmt.Fprintf(w, "hello word")
	})
	http.ListenAndServe(":8080", nil)
}

