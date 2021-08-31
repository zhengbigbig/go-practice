package main

import "fmt"

// 分硬币
//给定数量不限的硬币，币值为25分、10分、5分和1分，编写代码计算n分有几种表示法。
//示例：
//输入 n = 5
//输出 2
//解释：有两种方式可以凑成总金额 5=5 1+1+1+1+1=5


// 硬币[1,5,10,25]
// f(k,n) 只用前k种硬币组成n分的情况数量
// = f(k-1,n) // 只用前k-1种硬币组成n分的情况（使用0个第k个硬币）
// + f(k-1,n-COIN) // 只用前k-1种硬币组成n分的情况（使用0个第k个硬币）
// ...
// + f(k-1,n-i*COIN) // 只用前k-1种硬币组成n分的情况（使用0个第k个硬币）
// 直到n-i*COIN  为零
var coin = []int{0,1,5,10,25} // 避免数组下标0开始

func f(k,n int) int {
	if k == 1 {
		return 1
	}
	currentCoin := coin[k]
	result:=0
	for i := 0; n-i*currentCoin >=0; i++ {
		result += f(k-1,n-i*currentCoin)
	}
	return result
}

func howManyCoins(n int) int {
	return f(4,n)
}

func main(){
	fmt.Println(howManyCoins(10))
}