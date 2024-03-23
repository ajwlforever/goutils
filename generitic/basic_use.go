package generitic

// SliceInsert insert val into T in pos
func SliceInsert[T any](slice []T, pos int, val T) (res []T) {
	res = append(slice[:pos], append([]T{val}, slice[pos:]...)...)
	return
}

// IntSlice insert val into T in pos
func IntSliceInsert(slice []int, pos int, val int) (res []int) {
	res = append(slice[:pos], append([]int{val}, slice[pos:]...)...)
	return
}

type IntSlice[T int | int32 | int64 |
	uint | uint8 | uint16 | uint32 | uint64] []T
