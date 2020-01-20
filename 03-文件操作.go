package main

import (
	"fmt"
	"io"
	"os"
)

// 1. 创建文件 Create   文件不存在创建，文件存在，将文件内容清空
// 2. 打开文件 Open   以只读方式打开文件, 文件不存在，打开失败
// 3. 打开文件 OpenFile

// 4. 按字符串写
// 5. 按位置写
// 6. 按字节写

// 7.

//func main()  {
//	f,err:=  os.Create("/Users/penglimin/Desktop/aaa.jpg")
//	if err!=nil {
//		fmt.Println("create err:",err)
//		return
//	}
//
//	defer f.Close()
//
//	fmt.Println("sucessful")
//}
//
//func main(){
//	f,err := os.Open("/Users/penglimin/Desktop/aaa.txt")
//
//	if err!=nil {
//		fmt.Println("open err:",err)
//		return
//	}
//
//	defer f.Close()
//
//	_, err = f.WriteString("#####")
//	if err!=nil{
//		fmt.Println("Write String Err:",err)
//		return
//	}
//
//	fmt.Println("sucessful")
//}



//func main(){
//	f,err:=  os.Create("/Users/penglimin/Desktop/aaa.txt")
//	f,err = os.OpenFile("/Users/penglimin/Desktop/aaa.txt",os.O_RDWR,6)
//
//	if err!=nil {
//		fmt.Println("openFile err:",err)
//		return
//	}
//
//	defer f.Close()
//
//	_, err = f.WriteString("#####")
//	if err!=nil{
//		fmt.Println("Write String Err:",err)
//		return
//	}
//
//	fmt.Println("sucessful")
//}


func main(){
	f,err := os.OpenFile("/Users/penglimin/Desktop/aaa.txt",os.O_RDWR,6)

	if err!=nil {
		fmt.Println("openFile err:",err)
		return
	}

	defer f.Close()
	fmt.Println("sucessful")


	//n, err := f.WriteString("#####")

	n, err := f.WriteString("helloworld\r\n")

	if err!=nil{
		fmt.Println("Write String Err:",err)
		return
	}

	fmt.Println("writstrsing n= ",n)

	off,_ :=f.Seek(-5,io.SeekEnd)
	fmt.Println(off)

	n, _ = f.WriteAt([]byte("11111"),off)

	fmt.Println("write at n:",n)

}