package complexmutexexample

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

type Income struct {
	Source string
	Amount int
}

func Example() {
	//variable for bank balance

	var bankBalance int = 0
	var mutex sync.Mutex

	//print out starting values; 0
	fmt.Printf("initial account balance: %d.00", bankBalance)
	fmt.Println()

	//define weekly revenue
	incomes := []Income{
		{
			Source: "Main Job",
			Amount: 500,
		},
		{
			Source: "Grandma Gift",
			Amount: 10,
		},
		{
			Source: "Dog Walker",
			Amount: 50,
		},
		{
			Source: "Investments",
			Amount: 100,
		},
	}

	wg.Add(len(incomes))

	//loop through 52 weeks and print out how much is made
	for i, income := range incomes {

		go func(i int, income Income) {
			defer wg.Done()

			for week := 1; week <= 52; week++ {
				mutex.Lock()

				temp := bankBalance
				temp += income.Amount
				bankBalance = temp

				mutex.Unlock()
				fmt.Printf("On week %d, you earned $%d.00 from the source: %s\n", week, income.Amount, income.Source)
			}
		}(i, income)

	}

	wg.Wait()

	//print out final balance
	fmt.Printf("Final bank balance: $%d.00", bankBalance)
	fmt.Println()
}
