package leetcode

import (
	"fmt"
	"testing"
)

func Test1702(t *testing.T) {
	fmt.Println(maximumBinaryString1("000110"))
	fmt.Println(maximumBinaryString1("000110") == "111011")
	fmt.Println(maximumBinaryString1("1100") == "1110")
}
func maximumBinaryString1(binary string) string {
	cnt, n := 0, len(binary)
	cnt1 := 0
	flag := false
	if n == 1 {
		return binary
	}

	for _, c := range binary {
		if c == '1' {
			cnt++

			if flag {
				cnt1++
			}
		} else {
			flag = true
		}
	}
	cnt0 := n - cnt
	s := []rune(binary)
	for i := 0; i < n; i++ {
		s[i] = '1'
	}
	//fmt.Println((cnt - cnt1))

	if cnt0-1+(cnt-cnt1) >= 0 {
		s[cnt0-1+(cnt-cnt1)] = '0'
	}

	return string(s)
}
func maximumBinaryString(binary string) string {
	i := 0
	n := len(binary)
	s := []rune(binary)
	if n == 1 {
		return binary
	}
	for i < n-1 {
		if s[i] == '0' && s[i+1] == '0' {
			// "00"
			s[i] = '1'
			i++

		} else if s[i] == '1' && s[i+1] == '0' {
			// "10"
			r := i + 1
			for r+1 < n && s[r+1] == '0' {
				r++
			}
			s[r] = '1'
			s[i] = '0'

		} else {
			// "01" or "11"
			i++
		}
	}
	return string(s)
}
