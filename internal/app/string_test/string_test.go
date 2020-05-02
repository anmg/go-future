package string_test

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

func TestString(t *testing.T)  {
	str1 := "中"
	t.Log(len(str1))

	c := []rune(str1)
	t.Log(c[0])
	t.Log(len(c))
}

func TestStringRune(t *testing.T)  {
	str1 := "中华人民共和国"
	for _,c := range str1{
		t.Logf("%[1]c, %[1]d", c)
	}
}

func TestStringFn(t *testing.T)  {
	s := "A,B,C"
	parts := strings.Split(s, ",")

	t.Log(parts[0])

	s1 := strings.Join(parts, "_")

	t.Log(s1)

	strlist := []string {"eew","4rer"}
	reader3 := strings.Join(strlist, "_")
	t.Log(reader3)
}

func TestStringConv(t *testing.T)  {
	s := strconv.Itoa(1111)
	t.Log("str"+s)

	if i,err := strconv.Atoi("10");err ==nil {
		t.Log(10 + i)
	}

	t.Logf("%T", s)

	t.Log(reflect.TypeOf(s))
}
func TestStringToUpper(t *testing.T) {
	reader2 := strings.ToUpper(`{"branch":"beta","change_log":"add the rows{10}","channel":"fros","create_time":"2017-06-13 16:39:08","firmware_list":"","md5":"80dee2bf7305bcf179582088e29fd7b9","note":{"CoreServices":{"md5":"d26975c0a8c7369f70ed699f2855cc2e","package_name":"CoreServices","version_code":"76","version_name":"1.0.76"},"FrDaemon":{"md5":"6b1f0626673200bc2157422cd2103f5d","package_name":"FrDaemon","version_code":"390","version_name":"1.0.390"},"FrGallery":{"md5":"90d767f0f31bcd3c1d27281ec979ba65","package_name":"FrGallery","version_code":"349","version_name":"1.0.349"},"FrLocal":{"md5":"f15a215b2c070a80a01f07bde4f219eb","package_name":"FrLocal","version_code":"791","version_name":"1.0.791"}},"pack_region_urls":{"CN":"https://s3.cn-north-1.amazonaws.com.cn/xxx-os/ttt_xxx_android_1.5.3.344.393.zip","default":"http://192.168.8.78/ttt_xxx_android_1.5.3.344.393.zip","local":"http://192.168.8.78/ttt_xxx_android_1.5.3.344.393.zip"},"pack_version":"1.5.3.344.393","pack_version_code":393,"region":"all","release_flag":0,"revision":62,"size":38966875,"status":3}`)
	fmt.Println(reader2)
}