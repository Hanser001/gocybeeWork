package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int //n作为map的容量与我们输入自然数的个数
	fmt.Scanln(&n)
	inputNum := make([]int, n)
	numGroup := make(map[int]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanln(&inputNum[i]) //将输入的自然数用数组装好
	}
	sort.Ints(inputNum) //将数组里的元素从小到大排序

	for _, v := range inputNum { //将数组元素作为map的key填入，将value作为key出现的次数
		if numGroup[v] == 0 {
			numGroup[v] = 1 //如果一个key第一次出现，将它对应的value初始化为1
		} else {
			numGroup[v]++
		}
	}

	for k, v := range numGroup {
		fmt.Printf("%d %d\n", k, v)
	}
}
