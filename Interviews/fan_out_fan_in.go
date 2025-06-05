package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var numWorkers = 10

func generate() chan int {
	in := make(chan int)

	go func() {
		for i := range 100 {
			in <- i
		}
	}()

	return in
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	now := time.Now()
	for v := range fanIn(ctx, fanOut(generate(), numWorkers, f)) {
		fmt.Println(v)
	}
	fmt.Println(time.Since(now))
}

func fanIn(ctx context.Context, chans []chan int) chan int {
	out := make(chan int)

	go func() {
		wg := &sync.WaitGroup{}
		for _, ch := range chans {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for {
					select {
					case val, ok := <-ch:
						if !ok {
							return
						}
						select {
						case out <- val:
						case <-ctx.Done():
							return
						}
					case <-ctx.Done():
						return
					}
				}
			}()
		}

		wg.Wait()
		close(out)
	}()

	return out
}

func fanOut(in chan int, numChans int, f func(int) int) []chan int {
	chans := make([]chan int, numChans)

	for i := range chans {
		chans[i] = pipeline(in, f)
	}

	return chans
}

func pipeline(in chan int, f func(int) int) chan int {
	out := make(chan int)

	go func() {
		for v := range in {
			out <- f(v)
		}
		close(out)
	}()

	return out
}

func f(v int) int {
	return v * 2
}
