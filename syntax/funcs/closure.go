package main

func Closure(name string) func() string {
	// 闭包
	// name变量
	// 是否引用了方法外的变量
	return func() string {
		return "hello, " + name
	}
}
