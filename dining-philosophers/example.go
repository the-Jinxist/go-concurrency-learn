package diningphilosophers

import (
	"fmt"
	"sync"
	"time"
)

// The Dining Philosophers problem is well known in computer science circles.
// Five philosophers, numbered from 0 through 4, live in a house where the
// table is laid for them; each philosopher has their own place at the table.
// Their only difficulty – besides those of philosophy – is that the dish
// served is a very difficult kind of spaghetti which has to be eaten with
// two forks. There are two forks next to each plate, so that presents no
// difficulty. As a consequence, however, this means that no two neighbours
// may be eating simultaneously, since there are five philosophers and five forks.
//
// This is a simple implementation of Dijkstra's solution to the "Dining
// Philosophers" dilemma.

// Philosopher is a struct that stores info about a philosopher
type Philosopher struct {
	Name      string
	RightFork int
	LeftFork  int
}

// list of philosophers
var philosophers = []Philosopher{
	{Name: "Plato", RightFork: 0, LeftFork: 4},
	{Name: "Socrates", RightFork: 1, LeftFork: 0},
	{Name: "Aristotle", RightFork: 2, LeftFork: 1},
	{Name: "Pascal", RightFork: 3, LeftFork: 2},
	{Name: "Locke", RightFork: 4, LeftFork: 3},
}

var orderFinished []string
var orderMutex = &sync.Mutex{}

// define some variables
var hunger = 3                  // how many times does a person eat
var eatTime = 1 * time.Second   // how long a person eats
var thinkTime = 3 * time.Second // how long a person thinks
var sleepTime = 1 * time.Second // how long a person sleeps

func Example() {

	fmt.Println("Dining philoshoper's problem")
	fmt.Println("-----------------------------")

	fmt.Println("The table is empty")

	dine()

	fmt.Println("The table is empty")
	fmt.Printf("Order finished: %v ", orderFinished)

}

func dine() {

	// eatTime = 0 * time.Second
	// sleepTime = 0 * time.Second
	// thinkTime = 0 * time.Second

	wg := &sync.WaitGroup{}
	wg.Add(len(philosophers))

	seated := &sync.WaitGroup{}
	seated.Add(len(philosophers))

	//forks is the map of all 5 forks
	var forks = make(map[int]*sync.Mutex)
	for i := 0; i < len(philosophers); i++ {
		forks[i] = &sync.Mutex{}
	}

	//start the meal
	for i := 0; i < len(philosophers); i++ {
		go diningProblem(philosophers[i], wg, forks, seated)
	}

	wg.Wait()
}

func diningProblem(philosopher Philosopher, wg *sync.WaitGroup, forks map[int]*sync.Mutex, seated *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("%s is seated at the table \n", philosopher.Name)
	seated.Done()

	seated.Wait()

	for i := hunger; i > 0; i-- {

		if philosopher.LeftFork > philosopher.RightFork {
			//get a lock on both forks
			forks[philosopher.RightFork].Lock()
			fmt.Printf("%s has taken his right fork \n", philosopher.Name)

			forks[philosopher.LeftFork].Lock()
			fmt.Printf("%s has taken his left fork \n", philosopher.Name)

		} else {
			//get a lock on both forks
			forks[philosopher.LeftFork].Lock()
			fmt.Printf("%s has taken his left fork \n", philosopher.Name)

			forks[philosopher.RightFork].Lock()
			fmt.Printf("%s has taken his right fork \n", philosopher.Name)
		}

		fmt.Printf("%s has both forks and is eating \n", philosopher.Name)
		time.Sleep(eatTime)

		fmt.Printf("%s is thinking \n", philosopher.Name)
		time.Sleep(thinkTime)

		forks[philosopher.LeftFork].Unlock()
		forks[philosopher.RightFork].Unlock()

		fmt.Printf("%s put down the forks \n", philosopher.Name)

	}

	fmt.Println(philosopher.Name, " is satisfied. ")
	fmt.Println(philosopher.Name, " has left the table. ")

	orderMutex.Lock()
	orderFinished = append(orderFinished, philosopher.Name)
	orderMutex.Unlock()

}
