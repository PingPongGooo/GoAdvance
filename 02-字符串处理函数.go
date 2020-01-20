package main

import (
	"fmt"
	"strings"
)

// https://studygolang.com/pkgdoc

func main2()  {
	str := "I love my work"
	// 字符串指定分隔符 拆分
	ret:= strings.Split(str, " ")
	for _, s:= range  ret{
		fmt.Println(s)
	}


	// 字符串 按空格拆分
	strings.Fields(str)
	fmt.Println(ret)


	// 判断字符串结束标记
	flg := strings.HasSuffix("test.abc","abc")
	fmt.Println(flg)

	// 判断字符串其实标记
	flg2 := strings.HasPrefix("test.abc","tes")
	fmt.Println(flg2)
}