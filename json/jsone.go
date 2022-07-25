package jsone

import (
	"fmt"
	"strings"
)

//VtoJ 将键值对转换为JSON格式
func VtoJ(v interface{})([]byte){
	//首先将写入字符串
	str := fmt.Sprintf("%#v",v)
	//将字符串中的有用字段提取
	a := strings.SplitN(str,"",-1)
	var i int
	for m,v := range a {
		if v == "{"{
			i = m
			break
		}
	}
	var rs string
	for ;i<len(a);i++{
		rs =rs+a[i]
	}
	//对新字符串进行修饰，"变为\",字母转化为小写
	rs1 := strings.ToLower(rs)
	rs2 := strings.ReplaceAll(rs1,"\"","\\\"")
	//将最后的字符串转化为byte类型
	return []byte(rs2)
}