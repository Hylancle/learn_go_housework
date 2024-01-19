package main

import (
	"math"
	"strconv"
	"strings"
	"unicode/utf8"
)

func main() {
	//Caculate()
	//String()
	//Byte()
	Bool()
	//ExtremeNum()
}

// 计算
func Caculate() {
	var a int = 456
	var b int = 123
	println(a + b)
	println(a * b)
	println(float64(a) / float64(b))
	a++
	println(a)
	b--
	println(b)
	// 类型不同不会自动转化
	//println(float64(a) / b) 报错
	println(math.Abs(-123.4))
}

// 极值
func ExtremeNum() {
	println(math.MinInt64)
	println("float32最小正数", math.SmallestNonzeroFloat32)
}

// 字符串
func String() {
	// 单引号换行。 没法转义
	println("我是双引号")
	println(`我是单引号
	换行啦
	`)

	// int 转字符串itoa
	println("go" + strconv.Itoa(123))
	// 计算长度
	println(len("hello"))
	println(len("hello你好"))                    //11
	println(utf8.RuneCountInString("hello你好")) //7 主要用这个
	// 字符串相关方法： strings.xxx
	println(strings.Compare("123", "123"))
}

func Byte() {
	// 切片
	var str string = "this is a string"
	var bs []byte = []byte(str)
	var str1 = string(bs)
	println(bs)
	println(str1)
	var a byte = 'a' //97
	println(a)
}

func Bool() {
	// !(a&&b) => !a||!b
	// !(a||b) => !a && !b
	var a = true
	var b bool = false
	println(a)
	println(a && b)
}
