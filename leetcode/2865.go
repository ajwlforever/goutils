package leetcode

func maximumSumOfHeights(maxHeights []int) int64 {
	var cnt int64 = 0
	s1 := make([]int64, len(maxHeights))
	s2 := make([]int64, len(maxHeights))
	prefix := make([]int64, len(maxHeights))
	suffix := make([]int64, len(maxHeights))

	for i := 0; i < len(maxHeights); i++ {
		// decrement stack
		for {
			if len(s1) != 0 || s1[len(s1)-1] < int64(maxHeights[i]) {
				// pop
				s1 = s1[:len(s1)-1]
			} else {
				break
			}
		}
		// calcu prefix
		if len(s1) == 0 {
			prefix[i] = int64(maxHeights[i])

		} else {
			prefix[i] = s1[len(s1)-1]
		}
		if i-1 >= 0 {
			prefix[i] += prefix[i-1]
		}
	}

	print(prefix)
	for i := len(maxHeights) - 1; i >= 0; i-- {
		// increment stack
		for {
			if len(s2) != 0 || s2[len(s2)-1] < int64(maxHeights[i]) {
				// pop
				s2 = s2[:len(s2)-1]
			} else {
				break
			}
		}
		// calcu suffix
		if len(s2) == 0 {
			suffix[i] = int64(maxHeights[i])

		} else {
			suffix[i] = s2[len(s2)-1]
		}
		if i-1 >= 0 {
			suffix[i] += suffix[i-1]
		}

	}
	print(suffix)
	return cnt
}

// func main() {
// 	maximumSumOfHeights([]int{5, 3, 4, 1, 1})
// }
