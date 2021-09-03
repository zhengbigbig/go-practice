package main

import "fmt"

func Kmp(str, tg string) int {
	next := getNext(tg)
	j := 0
	for i := 0; i < len(str); i++ {
		for j >0 && str[i] != tg[j] {
			j = next[j-1]
		}
		if str[i] == tg[j] {
			j++
		}
		if j == len(tg) {
			return i - j + 1
		}
	}
	return -1
}

func getNext(tg string) []int {
	next := make([]int, len(tg)+1)
	next[0] = 0
	i, j := 1, 0
	for i <= len(tg) {
		if j == 0 || tg[i-1] == tg[j-1] {
			j++
			next[i] = j
			i++
		} else {
			j = next[j-1]
		}
	}
	return next
}

func main() {
	a := "abcdabvdaca"
	b := "aca"
	fmt.Println(Kmp(a,b))
}
