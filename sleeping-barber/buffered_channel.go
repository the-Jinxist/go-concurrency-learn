package sleepingbarber

import (
	"fmt"
	"time"
)

func BufferedChannelExample() {
	ch := make(chan int, 10)

	go listenToChan(ch)

	for i := 0; i < 100; i++ {
		fmt.Printf("Sending i %d to channel \n", i)
		ch <- i

		fmt.Printf("Sent i %d to channel \n", i)
	}

	fmt.Println("Done")
	close(ch)

}

func listenToChan(ch chan int) {
	for {
		i := <-ch
		fmt.Printf("Got i %d from channel \n", i)

		// simulate doing a lot of work
		time.Sleep(1 * time.Second)
	}
}
