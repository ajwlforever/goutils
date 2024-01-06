package syncs

import (
	"fmt"
	"sync"
)

func Mutex1(mutex *sync.Mutex, count *int64) {

	for i := 0; i < 100; i++ {
		mutex.Lock()
		*count = 1 + *count
		mutex.Unlock()
		fmt.Println("Mutex1: ", *count)
	}
}

func Mutex2(mutex *sync.Mutex, count *int64) {
	for i := 0; i < 100; i++ {
		mutex.Lock()
		*count = *count - 1
		mutex.Unlock()
		fmt.Println("Mutex2: ", *count)
	}
}

func NoMutex1(count *int64) {
	for i := 0; i < 100; i++ {
		*count = 1 + *count
		fmt.Println("NOMutex1: ", *count)
	}
}

func NoMutex2(count *int64) {
	for i := 0; i < 100; i++ {
		*count = *count - 1
		fmt.Println("NOMutex2: ", *count)
	}
}
