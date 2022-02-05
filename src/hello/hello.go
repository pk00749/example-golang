package main

import (
	"fmt"
	// 此处导入的名字和生成go.mod的命名相同，官网是"example.com/greetings"，本文章改成了 “greetings”
	// 官网引入
	//  "example.com/greetings"
	// 本文章引入
	"greetings"
)

func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}
