package main

import (
	"fmt"
	"time"
)

type play func() string

func caculate(ftype play) (result string) {
	result = ftype()
	return result
}

func 欢迎来我家玩() string {
	// 花费 5s 前往杰哥家
	time.Sleep(5 * time.Second)
	return "登dua郎"
}

func 打电动() string {
	return "输了啦，都是你害的."
}

func main() {
	go fmt.Println(caculate(打电动)) // 阿伟想在去杰哥家的路上打电动
	fmt.Println(caculate(欢迎来我家玩))
}
