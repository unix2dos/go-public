package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type Content struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Info struct {
			Id           int    `json:"id"`
			BasicType    int    `json:"basic_type"`
			Title        string `json:"title"`
			Type         int    `json:"type"`
			Price        string `json:"price"`
			Status       int    `json:"status"`
			ChapterId    int    `json:"chapter_id"`
			ChapterName  string `json:"chapter_name"`
			TopName      string `json:"top_name"`
			SequenceType int    `json:"sequence_type"`
		} `json:"info"`
		Lists []struct {
			Id       int    `json:"id"`
			Name     string `json:"name"`
			ImgVideo string `json:"img_video"`
			IvType   int    `json:"iv_type"`
			Type     int    `json:"type"`
			Option   []struct {
				Name  string `json:"name"`
				Value string `json:"value"`
				Text  string `json:"text"`
			} `json:"option"`
			Answer         string      `json:"answer"`
			Tof            interface{} `json:"tof"`
			RefAnswer      interface{} `json:"ref_answer"`
			Analyze        string      `json:"analyze"`
			AnalyzeFile    string      `json:"analyze_file"`
			FileType       int         `json:"file_type"`
			KnowledgePoint interface{} `json:"knowledge_point"`
			Star           int         `json:"star"`
			Similar        int         `json:"similar"`
			Tag            int         `json:"tag"`
			CollectStatus  int         `json:"collect_status"`
		} `json:"lists"`
	} `json:"data"`
	Timestamp int `json:"timestamp"`
}

func main() {
	// 打开JSON文件
	dir := "dir/"

	// 使用Walk函数遍历文件夹
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		// 错误处理
		if err != nil {
			fmt.Printf("Error accessing path %q: %v\n", path, err)
			return err
		}
		// 忽略文件夹本身
		if path == dir {
			return nil
		}
		// 打印文件名
		RW(dir, info.Name())
		return nil
	})

}

func RW(dir string, name string) {
	filename := dir + name

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// 读取文件内容
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// 解析JSON数据
	var content Content
	err = json.Unmarshal(data, &content)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	// 输出解析结果
	index := 0
	var totalText string
	for _, v := range content.Data.Lists {
		index++
		title := fmt.Sprintf("%d: %s", index, v.Name)

		var ans string
		for _, an := range v.Option {
			ans += fmt.Sprintf("%s %s\n", an.Name, an.Text)
		}

		okText := fmt.Sprintf("答：%s，%s", v.Answer, v.Analyze)
		allText := title + "\n" + ans + "\n" + okText + "\n\n"
		totalText += allText
	}

	{
		// 创建一个新文件，如果文件已存在则覆盖
		file, err := os.Create(strings.Replace(name, ".json", ".txt", -1))
		if err != nil {
			fmt.Println("Error creating file:", err)
			return
		}
		defer file.Close()

		// 写入内容到文件中
		_, err = file.WriteString(totalText)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
	}
}
