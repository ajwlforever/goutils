package generitic

import "fmt"

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

type MyMap[Key int | string, Value float32 | float64] map[Key]Value

type MyStruct[T int | string] struct {
	Key  string
	Data T
}

type IPrintData[T int | float32 | string] interface {
	Print(data T)
}

// 一个泛型通道，可用类型实参 int 或 string 实例化
type MyChan[T int | string] chan T

type C interface {
	P1 | P2
	Print() // 必须加这里
}

type P1 struct {
}
type P2 struct {
}

func (p P1) Print() {
	fmt.Println("p1")
}

func (p P2) Print() {
	fmt.Println("p2")
}

func GenericsUse1[T C](t T) {
	t.Print()
}

type C1 struct {
	*int
}
type F struct{}

func (F) M1[T any](t T){} // syntax error: method must have no type parameters

func main() {
    var f F[string]
    f.M1("hello")
	[4]byte(x)
}

*(*[4]byte)(x)