package main

import (
	"context"
	"fmt"
)

func main() {
	fmt.Println(deleteKey("user::cp_bell_num::2023-09-25::241445913631"))
	//fmt.Println(deleteKey(""))
}

func createClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "r-2ze1c8t4jhzsgsahzx.redis.rds.aliyuncs.com:6379", // Redis 地址
		Password: "t9gwDg5zL0",                                       // Redis 密码，如果没有设置密码，可以为空字符串 ""
		DB:       1,                                                  // Redis 数据库索引
	})

	return client
}

func deleteKey(key string) error {
	client := createClient()

	// 使用 Del 方法删除键
	result := client.Del(context.Background(), key)

	if err := result.Err(); err != nil {
		return err
	}

	return nil
}
