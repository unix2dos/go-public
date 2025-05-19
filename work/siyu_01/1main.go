package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"sync/atomic"
)

func downloadImage(url string, folderName string) error {
	// 发送 HTTP GET 请求获取图片
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// 检查 HTTP 响应状态码
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to download image: status code %d", resp.StatusCode)
	}

	err = os.MkdirAll(folderName, os.ModePerm)
	if err != nil {
		log.Fatalf("Failed to create directory: %v", err)
	}

	file, err := os.Create(fmt.Sprintf("%s/%s.jpg", folderName, filepath.Base(url)))
	if err != nil {
		fmt.Println("aaa", err)
		return err
	}
	defer file.Close()

	// 将 HTTP 响应的主体（图片数据）复制到文件中
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		fmt.Println("bbb", err)
		return err
	}

	return nil
}

func main() {
	// 打开 CSV 文件
	file, err := os.Open("20240816_2.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// 创建一个 CSV 阅读器
	reader := csv.NewReader(file)

	var counter int32 = 0

	// 逐行读取 CSV 文件
	ch := make(chan struct{}, 20)
	var wg sync.WaitGroup
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break // 文件读取完毕
		}
		if err != nil {
			log.Fatal(err)
		}
		if len(record[2]) < 20 {
			continue
		}
		wg.Add(1)
		go func(record []string) {
			ch <- struct{}{}
			defer func() {
				<-ch
				wg.Done()
			}()
			downloadImage(record[2], record[1])
			atomic.AddInt32(&counter, 1)
			fmt.Println("Counter after increment:", counter)
		}(record)
	}

	wg.Wait()

}
