package main

import (
	"fmt"
	"sync"
)

func printSomething(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(s)
}

var msg string
var wg sync.WaitGroup

func updateMessage(s string) {
	defer wg.Done()

	msg = msg + s

}

func main() {
	msg = "Hello World"

	// var mutex sync.Mutex

	wg.Add(2)
	go updateMessage("Hello, universe")
	go updateMessage("Hello, cosmos")

	wg.Wait()

	fmt.Println(msg)

}

// func waitgroupTest() {
// 	words := []string{"alpha", "beta", "delta", "gamma", "pi", "zeta", "eta", "theta", "epsilon"}

// 	wg.Add(len(words))
// 	for i, x := range words {
// 		go printSomething(fmt.Sprintf("%d: %s", i, x), &wg)
// 	}

// 	wg.Wait()

// 	wg.Add(1)
// 	printSomething("second", &wg)
// }

// func updateMessage(s string, mutex *sync.Mutex) {
// 	defer wg.Done()

// 	mutex.Lock()
// 	msg = msg + s
// 	mutex.Unlock()

// }

// func main() {
// 	msg = "Hello World"

// 	var mutex sync.Mutex

// 	wg.Add(2)
// 	go updateMessage("Hello, universe", &mutex)
// 	go updateMessage("Hello, cosmos", &mutex)

// 	wg.Wait()

// 	fmt.Println(msg)

// }
