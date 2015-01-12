package main

import "fmt"
import "recursion/eliminate"

func main() {

	f := fanIn(getIntChannelFromRecursion(3, -6), getIntChannelFromLoop(3, -6))

	fmt.Println(<-f)
	fmt.Println(<-f)
}

func getIntChannelFromRecursion(m, n int) <-chan int {
	c := make(chan int)
	go func() {
		c <- eliminate.TailRecursion(m, n)
	}()
	return c
}

func getIntChannelFromLoop(m, n int) <-chan int {
	c := make(chan int)
	go func() {
		c <- eliminate.TailRecEliminated(m, n)
	}()
	return c
}

func fanIn(input1, input2 <-chan int) <-chan string {
	c := make(chan string)
	go func() { c <- fmt.Sprintf("Recusion %d", <-input1) }()
	go func() { c <- fmt.Sprintf("Loop %d", <-input2) }()
	return c
}
