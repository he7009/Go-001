package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type SlideCounter struct {
	win      map[int64]int
	interval int64
	mu       *sync.Mutex
}

func (s *SlideCounter) Incr() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	t := time.Now().Unix()
	if _, ok := s.win[t]; ok {
		s.win[t] += 1
		s.RemoveOldWid(t)
		return s.win[t]
	}
	s.win[t] = 1
	s.RemoveOldWid(t)
	return s.win[t]
}

func (s *SlideCounter) Sum() int {
	sub := time.Now().Unix() - s.interval
	s.mu.Lock()
	defer s.mu.Unlock()
	var sum int
	for timestamp, r := range s.win {
		if timestamp > sub {
			sum += r
		}
	}
	return sum
}

func (s *SlideCounter) RemoveOldWid(t int64)  {
	sub := t - s.interval
	for i, _ := range s.win {
		if i <= sub {
			delete(s.win, i)
		}
	}
}

func NewSlide(i int64) *SlideCounter {
	return &SlideCounter{
		win:      make(map[int64]int),
		interval: i,
		mu:       &sync.Mutex{},
	}
}

func main() {
	s := NewSlide(5)

	for i := 0; i < 30; i++ {
		go func() {
			for {
				s.Incr()
				rand.Seed(time.Now().UnixNano())
				c := rand.Intn(100)
				time.Sleep(time.Duration(c) * time.Millisecond)
			}
		}()
	}

	tick := time.Tick(time.Second)
	for range tick {
		fmt.Printf("SUM:%v \n",s.Sum())
		fmt.Println(s.win)
	}
}
