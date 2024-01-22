package main

import "fmt"

func Slice() {
	s1 := []int{1, 2, 3, 4}
	fmt.Printf("s1: %v, len: %d, cap: %d \n", s1, len(s1), cap(s1))
	s2 := make([]int, 3, 4) // 初始化3个元素。容量为4 的切片
	fmt.Printf("s2: %v, len: %d, cap: %d \n", s2, len(s2), cap(s2))
	s2 = append(s2, 5)
	fmt.Printf("s2: %v, len: %d, cap: %d \n", s2, len(s2), cap(s2))

	s2 = append(s2, 6) //扩容
	fmt.Printf("s2: %v, len: %d, cap: %d \n", s2, len(s2), cap(s2))
	s3 := make([]int, 4) //长度容量都是0
	fmt.Printf("s3: %v, len: %d, cap: %d \n", s3, len(s3), cap(s3))
	fmt.Printf("s3[2]: %d", s3[2]) // 按照下标索引
	//fmt.Printf("s3[2]： %d", s3[99]) //返回panic

}

func SubSlice() {
	s1 := []int{2, 4, 6, 8, 10}
	s2 := s1[1:3] //不包括下标3
	fmt.Printf("s2: %v, len: %d, cap: %d \n", s2, len(s2), cap(s2))
	s3 := s1[1:]
	fmt.Printf("s3: %v, len: %d, cap: %d \n", s3, len(s3), cap(s3))

	s4 := s1[:3]
	fmt.Printf("s4: %v, len: %d, cap: %d \n", s4, len(s4), cap(s4))

}

func ShareSlice() {
	// 只要任何一个扩容， 就不再共享
	s1 := []int{2, 4, 6, 8, 10}
	s2 := s1[1:3] //不包括下标3
	fmt.Printf("s2: %v, len: %d, cap: %d \n", s2, len(s2), cap(s2))
	s2[0] = 199 // 此时数组1的s1也会被修改
	fmt.Printf("s2: %v, len: %d, cap: %d \n", s2, len(s2), cap(s2))
	fmt.Printf("s1: %v, len: %d, cap: %d \n", s1, len(s1), cap(s1))
	s1 = append(s1, 1)
	// 已经扩容
	fmt.Printf("s1: %v, len: %d, cap: %d \n", s1, len(s1), cap(s1))
	s2[1] = 2000 // s1不再修改
	fmt.Printf("s1: %v, len: %d, cap: %d \n", s1, len(s1), cap(s1))
	fmt.Printf("s2: %v, len: %d, cap: %d \n", s2, len(s2), cap(s2))

}
