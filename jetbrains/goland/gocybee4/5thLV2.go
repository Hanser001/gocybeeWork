package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

var ch = make(chan bool)

// 实现单次提醒功能
func oneRemind() {
	var inputTime string
	var schedule string
	fmt.Println("请设置提醒的时间(格式为：2006-01-02 15:04:05)：")
	scanner := bufio.NewScanner(os.Stdin) //这样才能接收字符串中的空格
	if scanner.Scan() {
		s1 := scanner.Text()
		inputTime = s1
	}
	fmt.Println("请设置提醒的内容：")
	fmt.Scanln(&schedule)
	now := time.Now()
	//获取时区
	local, _ := time.LoadLocation("Asia/Shanghai")
	//解析输入字符串得到时间
	scheduledTime, _ := time.ParseInLocation("2006-01-02 15:04:05", inputTime, local)
	//判断是否已经到达预定时间
	if now.Equal(scheduledTime) {
		fmt.Printf("您预定的日程\"%s\"将在10分钟开始\n", schedule)
	}
}

// 实现每日提醒功能
func dailyRemind() {
	var inputTime string
	var schedule string
	fmt.Println("请设置每日提醒的时间(格式为:2006-01-02 15:04:05)：")
	scanner := bufio.NewScanner(os.Stdin) //这样才能接收字符串中的空格
	if scanner.Scan() {
		s1 := scanner.Text()
		inputTime = s1
	}
	//获取时区
	local, _ := time.LoadLocation("Asia/Shanghai")
	//解析输入字符串得到时间
	scheduledTime, _ := time.ParseInLocation("2006-01-02 15:04:05", inputTime, local)
	fmt.Println("请设置提醒的内容：")
	fmt.Scanln(&schedule)
	//输入完毕后解锁weeklyRemind
	ch <- true
	for {
		now := time.Now()
		// next将接受当前的年月日和我们预定时间的时分秒与当地时区
		next := time.Date(now.Year(), now.Month(), now.Day(), scheduledTime.Hour(), scheduledTime.Minute(), scheduledTime.Second(), 0, local)
		// 用if来判断目前时间是不是已经超过了当日的预定时间,若超过换到第二天的预定时间
		if next.Sub(now) < 0 {
			next = next.Add(24 * time.Hour)
			next = time.Date(next.Year(), next.Month(), next.Day(), scheduledTime.Hour(), scheduledTime.Minute(), scheduledTime.Second(), 0, local)
		}
		t := time.NewTimer(next.Sub(now))
		<-t.C //阻塞程序直到到达预定时间
		fmt.Printf("您预定的日程\"%s\"将在10分钟开始\n", schedule)
	}
}

// 实现每周提醒功能
func weeklyRemind() {
	//先dailyRemind输入预定时间之前先堵塞住
	<-ch
	var inputTime string
	var schedule string
	fmt.Println("请设置每周提醒的时间(格式为:2006-01-02 15:04:05)：")
	scanner := bufio.NewScanner(os.Stdin) //这样才能接收字符串中的空格
	if scanner.Scan() {
		s1 := scanner.Text()
		inputTime = s1
	}
	//获取时区
	local, _ := time.LoadLocation("Asia/Shanghai")
	//解析输入字符串得到时间
	scheduledTime, _ := time.ParseInLocation("2006-01-02 15:04:05", inputTime, local)
	fmt.Println("请设置提醒的内容：")
	fmt.Scanln(&schedule)
	for {
		now := time.Now()
		// next将接受当前的年月日和我们预定时间的时分秒与当地时区
		next1 := time.Date(now.Year(), now.Month(), now.Day(), scheduledTime.Hour(), scheduledTime.Minute(), scheduledTime.Second(), 0, local)
		// next2是next1七天后的时间
		next2 := next1.Add(168 * time.Hour)
		// 用if来判断目前时间是不是已经超过了预定时间,若超过换到下一周的预定时间
		if next2.Sub(now) < 0 {
			next2 = next2.Add(24 * time.Hour)
			next2 = time.Date(next2.Year(), next2.Month(), next2.Day(), scheduledTime.Hour(), scheduledTime.Minute(), scheduledTime.Second(), 0, local)
		}
		t := time.NewTimer(next2.Sub(now))
		<-t.C //阻塞程序直到到达预定时间
		fmt.Printf("您预定的日程\"%s\"将在10分钟开始\n", schedule)
	}
}

func main() {
	var n int //n代表序号
	fmt.Println("提醒功能如下：")
	fmt.Println("1.单次日程提醒功能")
	fmt.Println("2.重复日程提醒功能")
	fmt.Println("请输入序号：")
	fmt.Scanln(&n)
	switch n {
	case 1:
		oneRemind()
	case 2:
		go weeklyRemind()
		go dailyRemind()
	default:
		os.Exit(0)
	}
}
