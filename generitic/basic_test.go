package generitic

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSliceInsert(t *testing.T) {
	slice1 := []int{1, 2, 3, 4}
	slice1 = SliceInsert(slice1, 2, 5)
	fmt.Println(slice1)
}

func TestIntSlices(t *testing.T) {
	ints := IntSlice[int]{1, 2, 3, 4}
	fmt.Println(ints)

	c := 'c'
	fmt.Printf("%v", reflect.TypeOf(c))

	int2s := IntSlice[rune]{'a', 'b', 'c'}
	fmt.Println(int2s)

	// int 转 int64
	fmt.Println(int64(100))
	fmt.Println(int64(100.0))
	fmt.Println(int64(100.))

	// int64 转 int
	fmt.Println(int(100))
}

func TestCompare1(t *testing.T) {
	ints := IntSlice[int]{1, 2, 3, 4}
	for i := 0; i < 1000; i++ {
		ints = SliceInsert(ints, 2, i)
	}
}

func TestCompare2(t *testing.T) {
	ints := IntSlice[int]{1, 2, 3, 4}
	for i := 0; i < 1000; i++ {
		ints = IntSliceInsert(ints, 2, i)
	}

}

func BenchmarkCompare1(b *testing.B) {
	ints := IntSlice[int]{1, 2, 3, 4}
	for i := 0; i < b.N; i++ {
		ints = SliceInsert(ints, 2, i)
	}
}
func BenchmarkCompare2(b *testing.B) {
	ints := IntSlice[int]{1, 2, 3, 4}
	for i := 0; i < b.N; i++ {
		ints = IntSliceInsert(ints, 2, i)
	}
}

func TestUse1(t *testing.T) {
	GenericsUse1(P1{})
}
