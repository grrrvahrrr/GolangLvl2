package task2

import (
	"log"
	"runtime"
)

func WorkerPull() {
	var workers = make(chan struct{}, 10)
	var number int

	for i := 0; i < 1000; i++ {
		if i%100 == 0 {
			runtime.Gosched()
		}
		workers <- struct{}{}
		go func() {
			defer func() {
				<-workers
			}()
			number++
		}()
	}
	log.Println("Final result is", number)
}
