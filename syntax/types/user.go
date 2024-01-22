package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func ChageUser() {
	u1 := User{"Tom", 18}
	fmt.Printf("%+v \n", u1)
	fmt.Printf("u1 address %p ]n", &u1)
	// 实际上，u1.ChangeName的时候，u1发生了复制，所以修改不成功
	u1.ChangeName("Jerry")
	u1.ChangeAge(35)
	fmt.Printf("%+v \n", u1)
	fmt.Printf("u1 address %p ]n", &u1)
}
func (u User) ChangeName(name string) {
	u.Name = name
}

func (u *User) ChangeAge(age int) {
	u.Age = age
}
