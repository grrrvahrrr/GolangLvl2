package main

import (
	"context"
	"log"
	"testing"
	"time"
)

func Test_workerPull(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	expected := 2000
	recieved, err := workerPull(ctx, 2000)
	if err != nil {
		log.Println(err)
	}
	if int(recieved) != expected {
		t.Errorf("Expected %d, but recieved %d", expected, recieved)
	}
}

func Test_workerPullMutex(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	expected := 2000
	recieved, err := workerPullMutex(ctx, 2000)
	if err != nil {
		log.Println(err)
	}
	if int(recieved.num) != expected {
		t.Errorf("Expected %d, but recieved %d", expected, recieved.num)
	}
}
