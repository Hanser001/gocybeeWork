package dao

import (
	"context"
	"fmt"
)

var ctx = context.Background()

// 写入后可以获取用户名和ID
func String(username string, id int) {
	err := Rdb.Set(ctx, username, id, 0)
	if err != nil {
		fmt.Println(err)
	}

	val := Rdb.Get(ctx, username)
	if val.Err() != nil {
		fmt.Println(val.Err())
	}
	fmt.Println("key", val)
}

// 用set实现点赞,也可以取消点赞
func SetLikes(userId int) {
	err := Rdb.SAdd(ctx, "BeLikedId", userId)
	if err != nil {
		fmt.Println(err)
	}
}

func DeleteLikes(userId int) {
	err := Rdb.SRem(ctx, "BeLikedId", userId)
	if err != nil {
		fmt.Println(err)
	}
}
