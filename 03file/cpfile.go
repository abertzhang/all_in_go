package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	//os Open 打开文件
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	//文件读取
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("Error Reading file", err)
		return
	}
	fmt.Println(string(data))
	//写入文件
	//Create(),文件存在就清理不存在就创建
	file2, err := os.Create("output.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer func(file2 *os.File) {
		err := file2.Close()
		if err != nil {
			fmt.Println("Close file erroe", err)
		}
	}(file2)
	_, err = file2.WriteString("hello world in file\n")
	if err != nil {
		fmt.Println(err)
		return
	}

}
