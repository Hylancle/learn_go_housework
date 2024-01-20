package demo

var (
	Global   = "包外可访问"
	internal = "包外不可访问"
)

const ( // 常量不可改变
	External   = "外部可访问"
	internalV1 = "外部不可访问"
)

// iota
const (
	Status0 = iota<<1 + 8 // 可以进行批量操作
	Status1
	Status2
	Stxxxx
)
