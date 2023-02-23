package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	// 获取用户选择的语言
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("请输入要生成 .gitignore 文件的语言：")

	scanner.Scan()
	selectedLang := scanner.Text()

	// 大小写不敏感
	selectedLang = strings.ToLower(selectedLang)
	// 去除空格
	selectedLang = strings.TrimSpace(selectedLang)
	// 首字母大写
	selectedLang = strings.Title(selectedLang)

	// 获取选定语言的 .gitignore 文件内容
	resp, err := http.Get(fmt.Sprintf("https://raw.githubusercontent.com/panw3i/gitignore/main/%s.gitignore", selectedLang))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// 将内容写入到 .gitignore 文件中
	err = ioutil.WriteFile(".gitignore", body, 0644)
	if err != nil {
		panic(err)
	}

	fmt.Println("生成 .gitignore 文件成功！")
}
