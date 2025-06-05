package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var numWorkers = 10

func generate(ctx context.Context) <-chan int {
	in := make(chan int)
	go func() {
		defer close(in)
		for i := 0; i < 100; i++ {
			select {
			case <-ctx.Done():
				return
			case in <- i:
			}
		}
	}()
	return in
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	in := generate(ctx)
	chans := fanOut(ctx, in, numWorkers, f)
	out := fanIn(ctx, chans)

	now := time.Now()
	for v := range out {
		fmt.Println(v)
	}
	fmt.Println(time.Since(now))
}

func fanIn(ctx context.Context, chans []<-chan int) <-chan int {
	out := make(chan int)

	go func() {
		wg := &sync.WaitGroup{}
		wg.Add(len(chans))

		for _, ch := range chans {
			go func(c <-chan int) {
				defer wg.Done()
				for {
					select {
					case val, ok := <-c:
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
			}(ch)
		}

		go func() {
			wg.Wait()
			close(out)
		}()
	}()

	return out
}

func fanOut(ctx context.Context, in <-chan int, numChans int, f func(int) int) []<-chan int {
	chans := make([]<-chan int, numChans)
	for i := range chans {
		chans[i] = pipeline(ctx, in, f)
	}
	return chans
}

func pipeline(ctx context.Context, in <-chan int, f func(int) int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for {
			select {
			case v, ok := <-in:
				if !ok {
					return
				}
				select {
				case out <- f(v):
				case <-ctx.Done():
					return
				}
			case <-ctx.Done():
				return
			}
		}
	}()
	return out
}

func f(v int) int {
	return v * 2
}
