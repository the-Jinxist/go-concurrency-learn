package sleepingbarber

import (
	"fmt"
	"log"
	"strings"
)

func SimpleExamples() {

	ping := make(chan string)
	pong := make(chan string)

	go shout(ping, pong)

	fmt.Println("Type something and press ENTER (enter Q to quit)")

	for {
		// print a prompt
		fmt.Print(" -> ")

		//get user input
		var userInput string
		_, _ = fmt.Scanln(&userInput)

		if strings.ToLower(userInput) == "q" {
			break
		}

		ping <- userInput
		//wait for a response

		response := <-pong
		fmt.Println("Response: ", response)
	}

	fmt.Println("Closing channels")
	close(ping)
	close(pong)

}

// In the params <-chan means the channel is a recieve-only channel
// chan<- means a send-only channel
func shout(ping <-chan string, pong chan<- string) {
	for {
		s, ok := <-ping
		// ok tells us whether is was a recieved value send on the channel; true
		// or a zero value sent when the channel was closed or empty
		if !ok {
			log.Println("channel is closed")
		}

		pong <- fmt.Sprintf("%s!!!!", strings.ToUpper(s))

	}
}
