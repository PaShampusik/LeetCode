package main

import (
	"context"
	"errors"
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

/*
У нас есть система для продажи ограниченного количества товара. Всего доступно N штук. Пользователи отправляют заказы через API.
Каждый заказ — это горутина, вызывающая метод PlaceOrder.

Нужно реализовать потокобезопасную систему, которая:
Принимает заказы через метод PlaceOrder(userID string) error
Возвращает ошибку, если товар закончился
Не допускает oversell даже при высокой нагрузке (1000+ конкурентных заказов)
Поддерживает context.Context — если заказ слишком долго висит, он отменяется

PlaceOrder после того как продал все товары должен вызвать shutdown — после вызова shutdown, новые заказы не принимаются, а текущие обрабатываются

Пример использования

s := NewOrderService(100)

for i := 0; i < 1000; i++ {
    go func(id int) {
        ctx, cancel := context.WithTimeout(context.Background(), time.Second)
        defer cancel()

        err := s.PlaceOrder(ctx, fmt.Sprintf("user-%d", id))
        if err != nil {
            fmt.Println("fail:", err)
        }
    }(i)
}

time.Sleep(3 * time.Second)

============================================
type OrderService struct {
    // ваш код
}

func NewOrderService(stock int) *OrderService

func (s *OrderService) PlaceOrder(ctx context.Context, userID string) error {
    time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

    // ваш код
}
func (s *OrderService) shutdown() { // ваш код }
*/

func main() {

	s := NewOrderService(100)

	for i := 0; i < 1000; i++ {
		go func(id int) {
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			err := s.PlaceOrder(ctx, fmt.Sprintf("user-%d", id))
			if err != nil {
				fmt.Println("fail:", err)
			}
		}(i)
	}

	time.Sleep(7 * time.Second)
}

type OrderService struct {
	stock        int32         // Атомарный счетчик остатка
	shutdownCh   chan struct{} // Канал для shutdown-сигнала
	shutdownOnce sync.Once     // Гарантирует однократный shutdown
}

func NewOrderService(stock int) *OrderService {
	return &OrderService{
		stock:      int32(stock),
		shutdownCh: make(chan struct{}),
	}
}

func (s *OrderService) PlaceOrder(ctx context.Context, userId string) error {
	select {
	case <-s.shutdownCh:
		return errors.New("service shut down")
	default:
	}

	var current int32
	for {
		current = atomic.LoadInt32(&s.stock)
		if current <= 0 {
			return errors.New("out of stock")
		}
		if atomic.CompareAndSwapInt32(&s.stock, current, current-1) {
			break
		}
		runtime.Gosched()
	}

	if ctx.Err() != nil {
		atomic.AddInt32(&s.stock, 1)
		return ctx.Err()
	}

	delay := 2000 * time.Millisecond
	select {
	case <-time.After(delay):
		if current == 1 {
			s.shutdown()
		}
		return nil
	case <-ctx.Done():
		atomic.AddInt32(&s.stock, 1)
		return ctx.Err()
	}
}

func (s *OrderService) shutdown() {
	s.shutdownOnce.Do(func() {
		close(s.shutdownCh)
	})
}
