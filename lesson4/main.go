package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	err := ctxPractice()
	if err != nil {
		log.Println(err)
	}
}

func workerPull() {
	var workers = make(chan struct{}, 10)
	var number int

	for i := 0; i < 1000; i++ {
		workers <- struct{}{}
		go func() {
			defer func() {
				<-workers
			}()
			number++
			time.Sleep(1 * time.Second)
		}()
	}
	time.Sleep(5 * time.Second)
	log.Println("Final result is", number)
}

func ctxPractice() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	ctxInf, cancelInf := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	defer cancelInf()

	numbers := make(chan int)
	go func(context.Context) {
		var number int
		for {
			number++
			numbers <- number
			time.Sleep(time.Second * 1)
		}
	}(ctxInf)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	for {
		select {
		case val := <-numbers:
			log.Println(val)
		case <-ctxInf.Done():
			return ctxInf.Err()
		case <-sigs:
			func(context.Context) {
				log.Println("Shutting down in a sec")
			}(ctx)
			return ctx.Err()
		}
	}
}
