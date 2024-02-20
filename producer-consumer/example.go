package producerconsumer

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

const NoOfPizzas = 10

var pizzasMade, pizzasFailed, total int

type Producer struct {
	data chan PizzaOrder
	quit chan chan error
}

type PizzaOrder struct {
	pizzaNumber int
	message     string
	success     bool
}

func (p *Producer) Close() error {
	ch := make(chan error)
	p.quit <- ch

	return <-ch
}

func Start() {
	color.Cyan("The Pizzerial is open for business")
	color.Cyan("----------------------------------")

	pizzaJob := &Producer{
		data: make(chan PizzaOrder),
		quit: make(chan chan error),
	}

	go pizzeria(pizzaJob)

	for i := range pizzaJob.data {
		if i.pizzaNumber <= NoOfPizzas {
			if i.success {
				color.Green(i.message)
				color.Green("Order %d is out for delivery!", i.pizzaNumber)
			} else {
				color.Red(i.message)
				color.Red("Order %d failed!", i.pizzaNumber)
			}

		} else {
			color.Cyan("Done making pizzas")
			if err := pizzaJob.Close(); err != nil {
				color.Red("error closing channel", err)
			}

		}
	}

	color.Cyan("----------------------------")
	color.Cyan("Done for the day")

	color.Cyan("We made %d pizzas but failed to make %d, with %d attemps in total", pizzasMade, pizzasFailed, total)
	switch {
	case pizzasFailed > 9:
		color.Red("it was an awful day")
	case pizzasFailed > 6:
		color.Red("it was not a very good day..")
	case pizzasFailed > 4:
		color.Yellow("it was an okay day..")
	case pizzasFailed >= 2:
		color.Yellow("it was a pretty good day")
	default:
		color.Green("it was a great day")
	}

}

func makePizza(pizzaNumber int) *PizzaOrder {
	pizzaNumber++
	if pizzaNumber <= NoOfPizzas {
		delay := rand.Intn(5) + 1
		fmt.Printf("recieved an order with number: %d!\n", pizzaNumber)

		rnd := rand.Intn(12) + 1
		var msg = ""
		success := false

		if rnd < 5 {
			pizzasFailed++
		} else {
			pizzasMade++
		}
		total++

		fmt.Printf("making pizza number :%d. It will take %d seconds\n", pizzaNumber, delay)

		time.Sleep(time.Duration(delay) * time.Second)

		if rnd <= 2 {
			msg = fmt.Sprintf("*** We ran out of ingredients for pizza %d", pizzaNumber)
		} else if rnd <= 4 {
			msg = fmt.Sprintf("*** The cook quit while making pizza %d", pizzaNumber)

		} else {
			success = true
			msg = fmt.Sprintf("*** Pizza Order %d is ready", pizzaNumber)
		}

		p := PizzaOrder{
			pizzaNumber: pizzaNumber,
			message:     msg,
			success:     success,
		}

		return &p
	}

	return &PizzaOrder{
		pizzaNumber: pizzaNumber,
	}
}

func pizzeria(pizzamaker *Producer) {

	var i = 0

	for {
		currentPizza := makePizza(i)
		if currentPizza != nil {
			i = currentPizza.pizzaNumber
			select {
			case pizzamaker.data <- *currentPizza:

			case quitChan := <-pizzamaker.quit:
				{
					//close channels
					close(pizzamaker.data)
					close(quitChan)
					return
				}

			}
		}

	}

}

//[1, 2, 3, 4] , (n) =>
// [2, 4, 6, 8]
