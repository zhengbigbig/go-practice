package main

import (
	"fmt"
	"strings"
)

var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

func main() {
	left := dispatchCoin()
	fmt.Println("剩下：", left)
}

func dispatchCoin() interface{} {
	leftCoins := coins
	rules := map[int][]string{
		1: {"e", "E"},
		2: {"i", "I"},
		3: {"o", "O"},
		4: {"u", "U"},
	}
	countGetCoin := func(u string) (coin int) {
		coin = 0
		for k, v := range rules {
			if strings.Contains(u, v[0]) || strings.Contains(u, v[1]) {
				coin += k
			}
		}
		return
	}
	for _, u := range users {
		coin := countGetCoin(u)
		distribution[u] = coin
		leftCoins -= coin
	}
	return leftCoins
}
