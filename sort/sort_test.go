package sort

import (
	"fmt"
	"testing"
)

func TestQuikSort(t *testing.T) {
	nums := []int{4, 1, 3, 6, 7, 5}
	quiksort(nums, 0, len(nums)-1)
	fmt.Println(nums)
	nums = []int{4, 1, 3, 6, 7, 8}
	quiksort(nums, 0, len(nums)-1)
	fmt.Println(nums)

}
func TestQuickSort2(t *testing.T) {
	nums := []int{4, 1, 3, 6, 7, 8, 8, 8}
	QuickSort(nums)
	fmt.Println(nums)
	nums = []int{4, 1, 3, 6, 7, 8, 8}
	QuickSortDesc(nums)
	fmt.Println(nums)
}
func TestBubbleSort(t *testing.T) {
	nums := []int{4, 1, 3, 6, 7, 8, 8, 8}
	bubbleSort(nums)
	fmt.Println(nums)
}

func BenchmarkBubbleSort(b *testing.B) {
	nums := genNums10w()
	b.ResetTimer()
	bubbleSort(nums)
	b.StopTimer()
}
func BenchmarkBubbleSortWithoutFlag(b *testing.B) {
	nums := genNums10w()
	b.ResetTimer()
	bubbleSortWithoutFlag(nums)
	b.StopTimer()
}
func BenchmarkQuickSort(b *testing.B) {

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		nums := genNums10w()
		QuickSort(nums)
	}
	b.StopTimer()
}
