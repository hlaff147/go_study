package main

import (
	"fmt"
	"time"
)

func PrintNumbers(name string, max int) {
	for i := 1; i <= max; i++ {
		fmt.Printf("%s: %d\n", name, i)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	go PrintNumbers("Thread 1", 1)
	go PrintNumbers("Thread 2", 1)

	time.Sleep(1 * time.Second)
	fmt.Println("Todas as threads terminaram.")
}
