package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
	"sync/atomic"
)

func downloadImage(url1 string, url2 string, index int32) error {
	{
		resp, err := http.Get(url1)
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("failed to download image: status code %d", resp.StatusCode)
		}
		file, err := os.Create(fmt.Sprintf("%d_saver.jpg", index))
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
	}

	{
		resp, err := http.Get(url2)
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("failed to download image: status code %d", resp.StatusCode)
		}
		file, err := os.Create(fmt.Sprintf("%d_origin.jpg", index))
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
	}
	return nil
}

func main() {
	// 打开 CSV 文件
	file, err := os.Open("0904.csv")
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
		fmt.Println(record)
		if len(record[1]) < 20 {
			continue
		}
		wg.Add(1)
		go func(record []string) {
			ch <- struct{}{}
			defer func() {
				<-ch
				wg.Done()
			}()
			atomic.AddInt32(&counter, 1)
			downloadImage(record[0], record[1], counter)
			//fmt.Println(record[0], record[1])
			//fmt.Println("Counter after increment:", counter)
		}(record)
	}

	wg.Wait()

}
