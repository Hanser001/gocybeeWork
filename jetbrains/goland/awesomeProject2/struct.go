package main

import (
	"time"
)

type Question struct {
	id           string    //用户的ID
	ipAdress     string    //IP属地
	content      string    //标题
	answer       string    //回答
	createdTime  time.Time //问题创建时间
	deletedTime  time.Time //删除时间
	updatedTime  time.Time //回答时间
	numofagree   int       //赞同数
	numofanswers int       //共多少条回答
}
