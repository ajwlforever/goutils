package father

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
	"testing"
)

func TestA1(t *testing.T) {

	file, _ := os.Open("664.txt")

	var bytes []byte
	bytes, _ = io.ReadAll(file)
	str := string(bytes)
	str = replaceDigitsAndRemoveEmptyLines(str)
	fmt.Println(str)

}

func replaceDigitsAndRemoveEmptyLines(input string) string {
	// 创建正则表达式
	regex, err := regexp.Compile(`\d+\.`)
	if err != nil {
		fmt.Println("Invalid regex pattern")
		return ""
	}

	// 替换所有的"数字加点"为一个空字符串
	replaced := regex.ReplaceAllString(input, "")

	// 将字符串分割成行
	lines := strings.Split(replaced, "\n")

	// 删除所有空行
	var nonEmptyLines []string
	for _, line := range lines {
		if strings.TrimSpace(line) != "" {
			nonEmptyLines = append(nonEmptyLines, line)
		}
	}

	// 重新组合为一个字符串
	result := strings.Join(nonEmptyLines, "\n")
	return result
}

func TestA2(t *testing.T) {

	file, _ := os.Open("2.txt")

	var bytes []byte
	bytes, _ = io.ReadAll(file)
	str := string(bytes)
	s := extractChineseChars(str)
	for _, v := range s {
		fmt.Println(v)
	}

}

func extractChineseChars(input string) []string {
	// 使用正则表达式匹配所有汉字
	regex := regexp.MustCompile(`[\p{Han}（）]+`)
	matches := regex.FindAllString(input, -1)

	return matches
}

func splitByDigitsAndRemoveSpaces(input string) string {
	// 第一步：去除所有空格
	spaceRemoved := strings.ReplaceAll(input, " ", "")

	// 第二步：按照数字分行
	// 使用正则表达式找到所有数字，并在数字前添加一个换行符，除了字符串开始的数字
	regex, err := regexp.Compile(`(\d+)`)
	if err != nil {
		fmt.Println("Invalid regex pattern")
		return ""
	}
	split := regex.ReplaceAllString(spaceRemoved, "\n$1")

	// 去除可能出现在字符串开始的换行符
	split = strings.TrimPrefix(split, "\n")

	return split
}
