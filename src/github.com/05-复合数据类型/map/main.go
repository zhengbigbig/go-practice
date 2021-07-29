package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	var m1 map[string]int
	fmt.Println(m1 == nil)
	m1 = make(map[string]int, 10)
	m1["age"] = 18
	m1["number"] = 10000
	fmt.Println(m1)
	fmt.Println(m1["age"])
	// 约定成俗用ok接收返回的布尔值
	fmt.Println(m1["a"]) // 如果不存在这个key，拿到对应值类型的零值，不推荐此写法
	value, ok := m1["a"]
	if !ok {
		fmt.Println("no key")
	} else {
		fmt.Println(value)
	}

	// map遍历
	for k, v := range m1 {
		fmt.Println(k, v)
	}
	for k := range m1 {
		fmt.Println(k)
	}
	for _, v := range m1 {
		fmt.Println(v)
	}
	// delete
	delete(m1, "age")

	rand.Seed(time.Now().UnixNano()) // 初始化随机数种子 拿到当前的纳秒数
	// 如果不使用随机数种子，每次拿到的结果都是一样的
	var scoreMap = make(map[string]int, 200)
	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%02d", i) // 生成stu开头的字符串
		value := rand.Intn(100)          // 生成0~99的随机整数
		scoreMap[key] = value
	}
	// 取出map中所有key存入切片keys
	var keys = make([]string, 0, 200)
	for key := range scoreMap {
		keys = append(keys, key)
	}
	// 对切片进行排序
	sort.Strings(keys)
	// 按照排序后的key遍历map
	for _, key := range keys {
		fmt.Println(key, scoreMap[key])
	}

	// 元素类型为map的切片
	var s1 = make([]map[int]string,10,10)
	//s1[0][0] = "A" 没有对slice和map进行初始化都会报错
	s1[0] = make(map[int]string,1)
	s1[0][10] = "xxx"
	fmt.Println(s1)

	// 值为切片类型的map
	var s2 = make(map[string][]int,10)
	s2["x"] = []int{1,2,3,4}
	fmt.Println(s2)
}
