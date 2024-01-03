package bingfa

import "fmt"

func sum(args ...int) int {
	var sum int
	for _, v := range args {
		sum += v
	}
	return sum
}

func myPrintf(args ...interface{}) {
	for _, v := range args {
		switch v.(type) {
		case int:
			fmt.Printf("int : %v", v)
		case string:
			fmt.Printf("string: %v", v)
		default:
			fmt.Printf("unkonwn type: %v", v)
		}
	}
}
