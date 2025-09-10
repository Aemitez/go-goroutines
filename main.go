package main

import (
	"fmt"
	"time"
)

func main() {

	// TODO example :No Goroutine (Sequential) vs  Goroutine
	// exampleCompare("Normal")

	// go exampleCompare("Power GO")

	// time.Sleep(time.Millisecond)
	// fmt.Println("done")

	// *--------------------------------------------------------

	// TODO example :2 Goroutine with chanel
	// ch := make(chan int)

	// go producer(ch) //  1 Goroutine
	// go consumer(ch) //  2 Goroutine

	// time.Sleep(time.Second)
	// fmt.Println("done")

	// *--------------------------------------------------------
}

func exampleCompare(state string) {
	start := time.Now()

	// ทำงานปกติ (รอ 2 วิ)
	fmt.Printf("%s: Start working...\n", state)
	time.Sleep(2 * time.Nanosecond)
	fmt.Printf("%s: Work done!\n", state)

	elapsed := time.Since(start)
	fmt.Printf("%s:Finished in %s\n", state, elapsed)
	fmt.Println("------------------------")
}

func producer(ch chan<- int) {
	for i := 1; i <= 5; i++ {
		fmt.Printf("Producing: %d\n", i)
		ch <- i
		time.Sleep(100 * time.Millisecond)
	}
	close(ch) // Close the channel when done
}

func consumer(ch <-chan int) {
	for num := range ch {
		fmt.Printf("Consuming: %d\n", num)
	}
}
