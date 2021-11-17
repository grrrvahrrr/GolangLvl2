package task3

import (
	"log"
)

func NumCount() {
	var number int
	for i := 0; i < 1000; i++ {
		go func() {
			number++
		}()
	}
	log.Println("Final result is", number)
}
