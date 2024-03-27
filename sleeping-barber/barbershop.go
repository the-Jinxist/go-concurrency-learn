package sleepingbarber

import (
	"time"

	"github.com/fatih/color"
)

type Barbershop struct {
	ShopCapacity    int
	HairCutDuration time.Duration
	NumberOfBarbers int
	BarbersDoneChan chan bool
	ClientsChan     chan string
	Open            bool
}

func (shop *Barbershop) addBarber(barber string) {
	shop.NumberOfBarbers++

	go func() {
		isSleeping := false
		color.Yellow("%s goes to the waiting room to check for clients", barber)

		for {
			if len(shop.ClientsChan) == 0 {
				isSleeping = true

				color.Yellow("%s takes a nap as there's nothing to do", barber)
			}

			client, shopOpen := <-shop.ClientsChan
			//shopOpen is true if the channel is still open
			// it becomes false after close(shop.ClientsChan) has been called
			if shopOpen {
				if isSleeping {
					color.Yellow("%s wakes %s up", client, barber)
					isSleeping = false

				}

				shop.cutHair(barber, client)

			} else {
				shop.sendBarberHome(barber)
				return

			}

		}

	}()
}

func (shop *Barbershop) cutHair(barber, client string) {
	color.Green("%s is cutting %s's hair", barber, client)
	time.Sleep(shop.HairCutDuration)
	color.Green("%s is finished cutting %s's hair", barber, client)

}

func (shop *Barbershop) sendBarberHome(barber string) {
	color.Cyan("%s is going home", barber)
	shop.BarbersDoneChan <- true
}

func (shop *Barbershop) closeShopForDay() {
	color.Cyan("Close shop for the day")
	close(shop.ClientsChan)
	shop.Open = false

	for a := 1; a <= shop.NumberOfBarbers; a++ {
		// this blocks until we get a value specifying that all the barbers are done
		<-shop.BarbersDoneChan
		close(shop.BarbersDoneChan)

		color.Green("--------------------------------------------------s")
		color.Green("Shop is closed for the day. Everyone has gone home")

	}

}

func (shop *Barbershop) addClient(client string) {
	color.Green("*** %s arrives", client)

	if shop.Open {
		select {
		// this actually sends to shop.ClientsChan
		case shop.ClientsChan <- client:
			color.Blue(" %s takes a seat in the waiting room", client)
		// this executes when shop.ClientsChan is full
		default:
			color.Red("The waiting room is filled so %s leaves", client)
		}
	} else {
		color.Red("The shop is already closed, so %s leaves", client)
	}

}
