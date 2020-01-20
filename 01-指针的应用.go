package main

import "fmt"

func main1(){
	var a int  = 10
	var p *int = &a

	a = 100
	fmt.Println("a = ", a) // 100

	fmt.Println("*p =", *p) // 100

	*p = 200
	fmt.Println("a =", a) // 200
}

