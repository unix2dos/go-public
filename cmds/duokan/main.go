package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	files, _ := ioutil.ReadDir("./")
	for _, f := range files {
		if !strings.HasPrefix(f.Name(), "duokan") && strings.HasSuffix(f.Name(), ".md") {
			Write(f.Name())
		}
	}
}

func Write(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	newFileName := "duokan_" + filename
	newFile, err := os.OpenFile(newFileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Println("文件打开失败", err)
		return
	}
	defer newFile.Close()

	var Title string
	var Date string
	var output string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		trimLine := strings.TrimSpace(line)

		{
			// 标题
			{
				pre1 := "- ## "
				if strings.HasPrefix(line, pre1) {
					Title = line[len(pre1):]
					//line = strings.Replace(line, pre1, "# ", -1)
					line = ""
					trimLine = ""
				}
				pre1 = "## "
				if strings.HasPrefix(line, pre1) {
					Title = line[len(pre1):]
					//line = strings.Replace(line, pre1, "# ", -1)
					line = ""
					trimLine = ""
				}
			}
			// 作者
			{
				pre2 := "- ##### "
				if strings.HasPrefix(line, pre2) {
					line = ""
					trimLine = ""
				}
				pre2 = "##### "
				if strings.HasPrefix(line, pre2) {
					line = ""
					trimLine = ""
				}
			}

			if _, _err := time.Parse("2006-01-02 15:04:05", trimLine); _err == nil {
				Date = trimLine
			}
			if strings.HasPrefix(trimLine, "**") && strings.HasSuffix(trimLine, "**") {
				line = strings.Replace(line, "**", "", -1)
				line = "\n\n" + "### " + line
			}

			output += line + "\n"
		}
	}

	header := GetHeader(Title, Date)
	output = header + output
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	//fmt.Println(output)
	newWrite := bufio.NewWriter(newFile)
	newWrite.WriteString(output)
	newWrite.Flush()
}

func GetHeader(title string, date string) string {
	text := fmt.Sprintf(`---
title: 《%s》
date: %s`, title, date)
	text += "\n---\n"
	return text
}
