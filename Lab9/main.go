package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var wg = sync.WaitGroup{}

func main() {

	fmt.Println("GOOS\t\t", runtime.GOOS)
	fmt.Println("ARCH\t\t", runtime.GOARCH)
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("GoRoutines\t", runtime.NumGoroutine())

	wg.Add(1)
	go foo() //Calls a new go routine
	bar()

	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("GoRoutines\t", runtime.NumGoroutine())
	wg.Wait()

	//----------------------
	fmt.Println("----------------")

	counter := 0
	const gs = 100
	var wg2 sync.WaitGroup
	var wg3 sync.WaitGroup
	wg2.Add(gs)
	wg3.Add(gs)

	var mu sync.Mutex

	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			v := counter
			//time.Sleep(time.Second)
			runtime.Gosched() //Yields processor, allowing other goroutines to run
			v++
			counter = v
			mu.Unlock()
			wg2.Done()
		}()
		fmt.Println("GoRoutines :", runtime.NumGoroutine())
	}
	wg2.Wait()
	fmt.Println("GoRoutines\t", runtime.NumGoroutine())
	fmt.Println("counter\t\t", counter)

	var counter2 int64

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&counter2, 1)
			runtime.Gosched()
			fmt.Println("Counter \t", atomic.LoadInt64(&counter2))
			wg3.Done()
		}()
		fmt.Println("GoRoutines :", runtime.NumGoroutine())
	}
	wg3.Wait()
	fmt.Println("GoRoutines\t", runtime.NumGoroutine())
	fmt.Println("counter\t\t", counter)
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("foo:", i)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("bar:", i)
	}
}
