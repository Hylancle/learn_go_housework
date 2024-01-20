package main

import "hello/syntax/variables/demo"

// 块声明(在包下都可用）
var (
	a1 int     = 123
	a2 float64 = 12.3
)

func main() {
	// var name type = value, 驼峰命名
	var str = "hello" // 可推断
	println(str)
	println("=======================")
	println(a1)
	a1 := 1234 // 变为局部变量
	println(a1)
	println("=======================")

	var e int
	println(e) // 0 默认值

	// 大小写可以控制包外访问,尽可能小写
	println(demo.Global)
	//println(demo.internal) error！
	println(demo.Status2)
}
