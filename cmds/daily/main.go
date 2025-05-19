package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {

	// 删除文件
	target_file := "./source/index.md"
	{
		os.Remove(target_file)
	}

	// 编辑文件
	AllFiles := make([]string, 0)
	{
		dir := "./source/_posts/think/闪念"
		files, _ := ioutil.ReadDir(dir)
		for _, f := range files {
			if !strings.HasSuffix(f.Name(), ".md") {
				continue
			}
			AllFiles = append(AllFiles, filepath.Join(dir, f.Name()))
		}
	}

	{
		f, err := os.OpenFile(target_file, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		if err != nil {
			panic(err)
		}
		defer f.Close()

		for i := len(AllFiles) - 1; i >= 0; i-- {
			content, err := ioutil.ReadFile(AllFiles[i])
			if err != nil {
				continue
			}
			text := string(content)
			find := "---\n"
			temp := text[len(find):]
			tempLen := strings.Index(temp, find)
			temp = text[tempLen+len(find)*2:]

			if _, err = f.WriteString(temp); err != nil {
				panic(err)
			}
		}
	}

}
