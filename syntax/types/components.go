package main

type Outer struct {
	Inner // 组合，优先使用这种。
	// 不是继承，相当于你再outer实现同名方法，仍然会用父类的方法
}

type OuterV1 struct {
	*Inner
}

type Inner struct {
}

func (i Inner) Hello() {
	println("Hello 我是Inner")
}

func Components() {
	var o Outer
	o.Hello()
	//指针需要初始化
	o1 := OuterV1{
		&Inner{},
	}
	o1.Hello()
}
