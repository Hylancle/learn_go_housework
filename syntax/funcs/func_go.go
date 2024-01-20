package main

import (
	"strconv"
)

// func 方法名（参数）返回值
func Func1() {
	println("hello")
}

func Func2(name string) {
	println("hello" + name)
}

func Func3(a, b int) (str string, err error) {
	return "a + b = " + strconv.Itoa(a+b), nil
}

func Func4() (string, error) {
	return "hello world", nil
}
func main() {
	//Func5()
	//Func6()
	//Func7Invoke()
	//Func8()
	//YourNameInvoke()
	//Defer()
	//DeferClosure()
	//DeferClosureV1()
	//fmt.Println("deferReturn的返回值", DeferReturn())
	//	//fmt.Println("deferReturnV1的返回值", DeferReturnV1())
	//	//structReturn := DeferReturnV2()
	//	//println(structReturn.name)
	//DeferClosureLoopV1()
	//DeferClosureLoopV2()
	DeferClosureLoopV3()

}
