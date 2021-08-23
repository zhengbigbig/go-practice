package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	// 1. 指针从最左侧开始向右移动，初始指针-1，长度0
	// 2. 向右移过程删除第一个字符的map kv
	// 3. 一直向右移动，当右侧值不在map中继续移动，如果在则退出
	// 4. 比较长度并赋值
	m := map[byte]int{}
	n := len(s)
	ptr,ret := -1,0
	for i:=0;i<n;i++ {
		if ptr == n -1 {
			break
		}
		if i != 0 {
			delete(m,s[i-1])
		}
		for ptr+1 < n && m[s[ptr+1]]==0{
			m[s[ptr+1]]++
			ptr++
		}
		ret = func(x,y int)int {
			if x > y {
				return x
			}
			return y
		}(ret,ptr-i+1)
	}
	return ret
}

// 第一轮 abcbbcabcd {a:1,b:1,c:1} -1 -> 2
// 第二轮 bcbbcabcd {b:1,c:1}    2  -> 2
// 第三轮 cbbcabcd {c:1,b:1}     2  -> 3
// ...
//


func main() {
	fmt.Println(lengthOfLongestSubstring("abcbbcabcd"))
	fmt.Println(lengthOfLongestSubstring("bbb"))
}
