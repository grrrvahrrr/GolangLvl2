package main

import (
	"context"
	"log"
	"sync"
	"sync/atomic"
	"time"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	ctx0, cancel0 := context.WithTimeout(context.Background(), time.Second)
	defer cancel0()
	test0, err := workerPull(ctx0, 1000)
	if err != nil {
		log.Println(err)
	}
	spew.Dump(test0)

	ctx1, cancel1 := context.WithTimeout(context.Background(), time.Second)
	defer cancel1()
	test1, err := workerPullMutex(ctx1, 1000)
	if err != nil {
		log.Println(err)
	}
	spew.Dump(test1)
}

func workerPull(ctx context.Context, operations int32) (int32, error) {
	var workers = make(chan struct{}, 10)
	var number int32
	var wg sync.WaitGroup

	wg.Add(int(operations))
	for i := 0; i < int(operations); i++ {
		workers <- struct{}{}
		go func() {
			defer func() {
				<-workers
				wg.Done()
			}()
			atomic.AddInt32(&number, 1)
		}()
	}
	wg.Wait()
	log.Println("Final result is", number)
	return number, ctx.Err()
}

func workerPullMutex(ctx context.Context, operations int32) (*number, error) {
	var workers = make(chan struct{}, 10)
	var number number
	var wg sync.WaitGroup

	wg.Add(int(operations))
	for i := 0; i < int(operations); i++ {
		workers <- struct{}{}
		go func() {
			defer func() {
				<-workers
				wg.Done()
			}()
			number.addOne()
		}()
	}
	wg.Wait()
	log.Println("Final result is", number.num)
	return &number, ctx.Err()
}

type number struct {
	num int
	mu  sync.Mutex
}

func (n *number) addOne() {
	n.mu.Lock()
	defer n.mu.Unlock()
	n.num++
}
