package ios

import (
	"fmt"
	"io"
	"testing"
)

func TestIo1(t *testing.T) {
	var reader io.Reader // 这里Reader只是一个空接口读不出东西，具体的实例可以是文件、网络连接或其他实现了 io.Reader 接口的类型。
	var p []byte         // 声明切片但是没有赋值，
	p = make([]byte, 1024)
	n, err := reader.Read(p)
	if err != nil {
		return
	}
	fmt.Printf("n %d", n)
}
