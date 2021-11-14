package lesson5task3

import (
	"log"
	"sync"
)

func writeRead(writeNum int, readNum int) {
	var number int
	var mu sync.Mutex

	for i := 0; i < writeNum; i++ {
		go func() {
			mu.Lock()
			number++
			mu.Unlock()
		}()
	}

	for i := 0; i < readNum; i++ {
		go func() {
			mu.Lock()
			log.Println(number)
			mu.Unlock()
		}()
	}
}

func writeReadRW(writeNum int, readNum int) {
	var number int
	var mu sync.RWMutex

	for i := 0; i < writeNum; i++ {
		go func() {
			mu.Lock()
			number++
			mu.Unlock()
		}()
	}

	for i := 0; i < readNum; i++ {
		go func() {
			mu.RLock()
			log.Println(number)
			mu.RUnlock()
		}()
	}
}
