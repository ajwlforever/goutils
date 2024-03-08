package sort

func quiksort(nums []int, i int, j int) {
	//fmt.Println(i, ",", j)
	if i >= j {
		return
	}
	left, right := i, j
	for left < right {
		// 从右边找比nums[left]小的，swap
		for left < right {
			if nums[right] > nums[left] {
				right--
			} else {
				// swap
				swap(nums, left, right)
				break
			}
		}

		// 从左边找比nums[right]大的,swap
		for left < right {
			if nums[left] < nums[right] {
				left++
			} else {
				//swap
				swap(nums, left, right)
				break
			}
		}
	}

	quiksort(nums, i, left-1)
	quiksort(nums, right+1, j)
}
func quicksort2(nums []int, i int, j int) {
	if i >= j {
		return //boundary
	}
	//fmt.Println(i, ",", j)
	left, right := i, j
	for left < right {
		// 从右边找比nums[left]小的，swap
		for left < right && nums[left] <= nums[right] {
			right--
		}
		//fmt.Println(left, ",", right, ",", nums)
		swap(nums, left, right)
		// 从左边找比nums[right]大的，swap
		for left < right && nums[left] <= nums[right] {
			left++
		}
		swap(nums, left, right)
	}

	quicksort2(nums, i, left-1)
	quicksort2(nums, right+1, j)
}
func quicksort2Desc(nums []int, i int, j int) {
	if i >= j {
		return //boundary
	}
	//fmt.Println(i, ",", j)
	left, right := i, j
	for left < right {
		// 从右边找比nums[left]小的，swap
		for left < right && nums[left] <= nums[right] {
			right--
		}
		//fmt.Println(left, ",", right, ",", nums)
		swap(nums, left, right)
		// 从左边找比nums[right]大的，swap
		for left < right && nums[left] <= nums[right] {
			left++
		}
		swap(nums, left, right)
	}

	quicksort2Desc(nums, i, left-1)
	quicksort2Desc(nums, right+1, j)
}
func QuickSort(nums []int) {
	quicksort2(nums, 0, len(nums)-1)
}
func QuickSortDesc(nums []int) {
	quicksort2(nums, 0, len(nums)-1)
}
func swap(nums []int, left int, right int) {
	nums[right], nums[left] = nums[left], nums[right]
}
