package main

import (
	"fmt"
	"time"
)

func main() {

	//fmt.Println("----------------")

	c := make(chan int)

	//c <- 42 //fatal error: all goroutines are asleep - deadlock!

	go func() { //That's the solution
		c <- 42
	}()
	fmt.Println(<-c)

	//Or

	fmt.Println("----------------")
	c2 := make(chan int, 1) //Buffered channel
	c2 <- 10
	//c2 <- 20 //fatal error: all goroutines are asleep - deadlock! //Buffered channel can only hold one value
	fmt.Println(<-c2)
	fmt.Println("----------------")
	c3 := make(chan int, 2) //Buffered channel
	c3 <- 10
	c3 <- 52
	fmt.Println(<-c3)
	fmt.Println(<-c3)
	fmt.Println("----------------")

	/*c4 := make(chan<- int, 2) //send //Does not work
	c4 <- 18
	c4 <- 21
	fmt.Println(<-c4)	//cannot receive from send-only channel c4 (variable of type chan<- int)
	fmt.Println(<-c4)	//cannot receive from send-only channel c4 (variable of type chan<- int)*/

	/*c5 := make(<-chan int, 2) //receive //Does not work
	c5 <- 18
	c5 <- 21
	fmt.Println(<-c5)
	fmt.Println(<-c5)*/

	fmt.Println("----------------")

	c4 := make(chan int)

	//send
	go foo(c4)
	//receive
	for v := range c4 {
		fmt.Println(v)
	}
	fmt.Println("About to exit")

	fmt.Println("----------------")

	even := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	//send
	go send(even, odd, quit)

	//receive
	receive(even, odd, quit)
	fmt.Println("About to exit")

	fmt.Println("----------------")

	fmt.Println("Main started")
	go start()
	fmt.Println("Now sleeping main routine")
	time.Sleep(3 * time.Second)
	fmt.Println("Main ended")

}

func receive(e, o, q <-chan int) {
	for {
		select {
		case v := <-e:
			fmt.Println("From the even channel:", v)
		case v := <-o:
			fmt.Println("From the odd channel:", v)
		case v := <-q:
			fmt.Println("From the quit channel:", v)
			return
		}
	}
}

func send(e, o, q chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	close(e)
	close(o)
	q <- 0

}

// send
func foo(c chan<- int) {
	//c <- 42
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
}

// receive
func bar(c <-chan int) {
	fmt.Println(<-c)
}

func start() {
	fmt.Println("Goroutine started")
	time.Sleep(1 * time.Second)
	fmt.Println("Goroutine ended")
}
