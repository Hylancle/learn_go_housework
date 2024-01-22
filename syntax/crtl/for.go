package main

func Loop2() {
	for i := 0; i < 10; i++ {

	}
	i := 0
	for i < 10 {
		println(i)
	}
}

func ForArray() {
	println("遍历数组")
	arr := [3]string{"A", "B", "C"}
	for idx, val := range arr {
		println(idx, val)
	}
}

// 切片
func ForSlice() {
	println("遍历切片")
	arr := []string{"A", "B", "C"}
	for idx, val := range arr {
		println(idx, val)
	}
}

func ForMap() {
	println("遍历map")
	m := map[string]string{
		"1": "A",
		"2": "B",
	}
	for key, value := range m {
		println(key, value)
	}
}

func LoopBug() {
	users := []User{
		{
			name: "Tom",
		}, {
			name: "Jerry",
		},
	}
	// 指针
	m := make(map[string]*User)
	for _, u := range users {
		m[u.name] = &u // 拿走了一个桶， 指向的永远是第一个地址
	}
	for _, u := range m {
		println(u.name) // Jerry、 jerry
	}
}

type User struct {
	name string
}

func Switch(status []int) string {
	switch status {
	case nil:
		return "初始化"
	}
	return "未知"
}

func SwitchV1(age int) string {
	switch {
	case age > 17:
		return "成年"
	case age > 35:
		return "失业"
	}
	return "未知"
}

func main() {
	println(SwitchV1(30))
}
