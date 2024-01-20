package main

func Func5() {
	// 方法赋值到变量
	myFunc3 := Func3
	s, err := myFunc3(1, 2)
	println(s, err)
}

func Func6() {
	// 局部方法
	fn := func(name string) string {
		return "hello, " + name
	}
	str := fn("大明")
	println(str)
}

func Func7() func(name string) string {
	return func(name string) string {
		return "hello, " + name + ", 我是方法作为返回值"
	}
}

// 调用返回方法
func Func7Invoke() {
	sayHello := Func7()
	str := sayHello("蛋蛋")
	println(str)
}

func Func8() {
	// 直接调用方法
	fn := func(name string) string {
		return "hello" + name
	}("小明")
	println(fn)
}
