package main

func Map() {
	m1 := map[string]string{
		"key1": "value1",
		"key3": "value3",
		"key4": "value4",
		"key5": "value5",
	}
	println(m1)
	m2 := make(map[string]string, 4) // 尽可能容量。不然扩容很慢
	m2["key2"] = "value2"
	val1, ok := m1["key1"]
	println(val1, ok)
	// val2是值， ok是是否存在, 若不存在返回默认值
	val2, ok := m2["key2"]
	println(val2, ok)
	println("第1次遍历")

	for k, v := range m1 {
		println(k, v)

	}
	// map是随机的

	// 删除操作
	delete(m1, "key1")
	for k, v := range m1 {
		println(k, v)
	}
}
