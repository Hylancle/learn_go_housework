package main

import "fmt"

type Fish struct {
}

func (f Fish) swim() {

}

// 衍生类型
type FakeFish Fish

func (f FakeFish) FakeSwim() {

}

type Yu = Fish //别名
func UserFish() {
	f1 := Fish{}
	f1.swim()

	f2 := FakeFish{}
	f2.FakeSwim()

	f3 := FakeFish(f1)
	fmt.Println(f1, f3)
}
