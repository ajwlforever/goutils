package leetcode

func maximumOddBinaryNumber(s string) string {
	n := len(s)
	s1 := []byte(s)
	cnt := 0
	for _, v := range s1 {
		if v == '1' {
			cnt++
		}
	}
	s1[n-1] = '1'
	if cnt == 1 {

		for i := 0; i < n-1; i++ {
			s1[i] = '0'
		}
		return string(s1)
	} else {
		for i := 0; i < cnt-1; i++ {
			s1[i] = '1'
		}
		for i := cnt - 1; i < n-1; i++ {
			s1[i] = '0'
		}
		return string(s1)
	}
}
