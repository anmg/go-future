package json_test

import (
	"encoding/json"
	"fmt"
	"github.com/json-iterator/go"
	"strings"
	"testing"
)

func TestJson(t *testing.T) {
	var jsonBlob = []byte(`[
        {"Name": "Platypus", "Order": "Monotremata"},
        {"Name": "Quoll",    "Order": "Dasyuromorphia"}
    ]`)
	type Animal struct {
		Name  string
		Order string
		Age int
	}
	var animals []Animal
	err := json.Unmarshal(jsonBlob, &animals)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%+v\n", animals)

	var animals2 []Animal
	var json_iterator = jsoniter.ConfigCompatibleWithStandardLibrary
	json_iterator.Unmarshal(jsonBlob, &animals2)
	fmt.Printf("%+v\n", animals2)

	for i, item := range animals2{
		t.Log(i, item.Name, item.Order)
	}
}

func TestJsonGet(t *testing.T) {
	val := []byte(`{"ID":1,"Name":"Reds","Colors":["Crimson","Red","Ruby","Maroon"]}`)
	//str := jsoniter.Get(val, "Colors", 1).ToString()
	str2 := jsoniter.Get(val, "Colors").ToString()
	fmt.Println(str2)
}

func TestJsonDecode(t *testing.T) {
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	reader := strings.NewReader(`{"branch":"beta","change_log":"add the rows{10}","channel":"fros","create_time":"2017-06-13 16:39:08","firmware_list":"","md5":"80dee2bf7305bcf179582088e29fd7b9","note":{"CoreServices":{"md5":"d26975c0a8c7369f70ed699f2855cc2e","package_name":"CoreServices","version_code":"76","version_name":"1.0.76"},"FrDaemon":{"md5":"6b1f0626673200bc2157422cd2103f5d","package_name":"FrDaemon","version_code":"390","version_name":"1.0.390"},"FrGallery":{"md5":"90d767f0f31bcd3c1d27281ec979ba65","package_name":"FrGallery","version_code":"349","version_name":"1.0.349"},"FrLocal":{"md5":"f15a215b2c070a80a01f07bde4f219eb","package_name":"FrLocal","version_code":"791","version_name":"1.0.791"}},"pack_region_urls":{"CN":"https://s3.cn-north-1.amazonaws.com.cn/xxx-os/ttt_xxx_android_1.5.3.344.393.zip","default":"http://192.168.8.78/ttt_xxx_android_1.5.3.344.393.zip","local":"http://192.168.8.78/ttt_xxx_android_1.5.3.344.393.zip"},"pack_version":"1.5.3.344.393","pack_version_code":393,"region":"all","release_flag":0,"revision":62,"size":38966875,"status":3}`)
	fmt.Println(reader.Len(), reader.Size(), *reader)
	decoder := json.NewDecoder(reader)
	params := make(map[string]interface{})
	err := decoder.Decode(&params)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%+v\n", params)
	}
}