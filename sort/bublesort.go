package sort

func bubbleSort(nums []int) {
	n := len(nums)

	for i := 0; i < n; i++ {
		flag := false
		for j := 0; j < n-i-1; j++ {
			// 把大的放后面
			if nums[j] > nums[j+1] {
				swap(nums, j, j+1)
				flag = true
			}
		}
		if flag == false {
			return
		}
	}
}
func bubbleSortWithoutFlag(nums []int) {
	n := len(nums)

	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			// 把大的放后面
			if nums[j] > nums[j+1] {
				swap(nums, j, j+1)

			}
		}
	}
}
