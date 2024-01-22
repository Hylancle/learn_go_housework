package main

// 接口，只有方法
type List interface {
	Add(index int, val any)
	Append(val any) error
	Delete(index int) error
}

type LinkedList struct {
	head node
	name string
	Head node //外部可访问
}

func (l *LinkedList) Append(val any) error {
	//TODO implement me
	panic("implement me")
}

func (l *LinkedList) Delete(index int) error {
	//TODO implement me
	panic("implement me")
}

func (l *LinkedList) Add(index int, val any) {

}

type node struct {
}
