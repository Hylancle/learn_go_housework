package main

func main() {
	//// 初始化结构体
	//u1 := &User{}
	//println(u1)
	//
	//u2 := User{}
	//u2.Name = "John"
	//fmt.Printf("%+v \n", u2)
	//
	//u2.Name = "Herry"
	//
	//var u3 User
	//
	//println(u3.Age)
	//
	//u5 := User{Name: "Jerry", Age: 18}
	//println(u5.Name)
	//
	//ChageUser()
	Components()
}

func UseList() {
	l1 := LinkedList{}
	// 指针
	l1Ptr := &l1
	//解引用， l2不是指针
	var l2 LinkedList = *l1Ptr
	println(l2.name)
}
