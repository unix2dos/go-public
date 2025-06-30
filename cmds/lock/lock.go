package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var rwMu sync.RWMutex

	// 获取写锁
	rwMu.Lock()

	// 尝试获取读锁
	go func() {
		if rwMu.TryRLock() {
			fmt.Println("成功获取读锁")
			rwMu.RUnlock()
		} else {
			fmt.Println("获取读锁失败")
		}
	}()

	// 尝试获取写锁
	go func() {
		time.Sleep(50 * time.Millisecond)
		if rwMu.TryLock() {
			fmt.Println("成功获取写锁")
			rwMu.Unlock()
		} else {
			fmt.Println("获取写锁失败")
		}
	}()

	time.Sleep(200 * time.Millisecond)
	rwMu.Unlock()
}
