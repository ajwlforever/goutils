package leetcode

import (
	"fmt"
	"testing"
)

type MyHashMap struct {
	cache map[int]int
}

func Constructor() MyHashMap {
	return MyHashMap{
		cache: make(map[int]int, 0),
	}
}

func (this *MyHashMap) Put(key int, value int) {
	if _, ok := this.cache[key]; !ok {
		this.cache[key] = value
	}
}

func (this *MyHashMap) Get(key int) int {
	if v, ok := this.cache[key]; ok {
		return v
	}
	return -1
}

func (this *MyHashMap) Remove(key int) {
	if _, ok := this.cache[key]; ok {
		delete(this.cache, key)
	}
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */

func Test706(t *testing.T) {

	obj := Constructor()
	obj.Put(1, 1)
	fmt.Println(obj.Get(1))
	obj.Remove(1)
	fmt.Println(obj.Get(1))

}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
