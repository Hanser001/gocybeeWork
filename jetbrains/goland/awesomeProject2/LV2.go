package main

import (
	"fmt"
	"time"
)

var tFlag1 = make(chan bool)
var tFlag2 = make(chan bool)

func loadRounds() {
	time.Sleep(time.Second) //装弹花一秒
	fmt.Print("装弹->")
	tFlag1 <- true
}

func takeAim() {
	<-tFlag1                    //保证先装弹后瞄准
	time.Sleep(2 * time.Second) //瞄准花两秒
	fmt.Print("瞄准->")
	tFlag2 <- true
}

func fire() {
	<-tFlag2                //保证先瞄准，后开火
	time.Sleep(time.Second) //花一秒发射
	fmt.Print("发射!\n")
}

func main() {
	for {
		go loadRounds()
		go takeAim()
		fire()
	}
}
