package main

// 先定义的后调用 defer
func Defer() {
	defer func() {
		println("The 1st defer")
	}()

	defer func() {
		println("The 2nd defer")
	}()
	//The 2nd defer
	//The 1st defer
}

func DeferClosure() { // 1
	i := 0
	defer func() {
		println(i)
	}()
	i++
}

func DeferClosureV1() { //0
	i := 0
	defer func(val int) {
		println(val)
	}(i)
	i++
}

// 在defer里面想要篡改返回值，唯一的办法就是把数定义在参数里
func DeferReturn() int {
	a := 0
	defer func() {
		a = 1
	}()
	return a //0
}

func DeferReturnV1() (a int) {
	a = 0
	defer func() {
		a = 1
	}()
	return a //1
}

func DeferReturnV2() *MyStruct {
	res := &MyStruct{
		name: "Tom",
	}
	defer func() {
		// 能改成功是因为只改了指针指向的值
		res.name = "Jerry"
	}()
	return res
}

type MyStruct struct {
	name string
}

func DeferClosureLoopV1() {
	for i := 0; i < 10; i++ {
		defer func() {
			print(i)
		}()
	}
}
func DeferClosureLoopV2() {
	for i := 0; i < 10; i++ {
		defer func(val int) {
			print(val)
		}(i)
	}
}

func DeferClosureLoopV3() {
	for i := 0; i < 10; i++ {
		defer func() {
			print(i)
		}()
	}
}
