package interface_any

import (
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	var cache map[any]any
	cache = make(map[any]any, 0)
	fmt.Println(cache[0])

}
