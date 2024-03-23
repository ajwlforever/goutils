package leetcode

import "testing"

func groupAnagrams(strs []string) [][]string {
	// bitmap
	res := make([][]string, 0)
	idxs := make(map[int]int, 0)
	for _, str := range strs {
		cnt := 0
		for _, c := range str {
			cnt <<= int(c - 'a')
		}
		if val, ok := idxs[cnt]; !ok {
			res = append(res, []string{str})
			idxs[cnt] = len(res)
		} else {
			res[val] = append(res[val], str)
		}

	}
	return res
}

type Int interface {
	~int | ~int32 | ~int64 | ~uint | ~uint64 | ~uint32
}

func Test1233(t *testing.T) {

}
