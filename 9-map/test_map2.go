package main

func printMap(m map[string]string) {
	// 传递的是引用类型
	for key, value := range m {
		println(key, value)
	}
}

func main() {
	cityMap := map[string]string{
		"China": "Beijing",
		"Japan": "Tokyo",
		"USA":   "Washington",
		"UK":    "London",
	}

	// 遍历
	printMap(cityMap)

	// 删除
	delete(cityMap, "China")

	// 修改
	cityMap["Japan"] = "Osaka"

	println("----")

	printMap(cityMap)
}
