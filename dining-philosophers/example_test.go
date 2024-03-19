package diningphilosophers

import (
	"testing"
	"time"
)

func Test_dine(t *testing.T) {

	eatTime = 0 * time.Second
	sleepTime = 0 * time.Second
	thinkTime = 0 * time.Second

	for i := 0; i < 10; i++ {
		orderFinished = []string{}
		dine()

		if len(orderFinished) > 5 {
			t.Errorf("Incorrect length of slice; expected 5 got %d", len(orderFinished))
		}
	}

}

func Test_dineWithVaryingDelays(t *testing.T) {
	var tests = []struct {
		Name  string
		Delay time.Duration
	}{
		{
			Name:  "zero delay",
			Delay: time.Second * 0,
		},

		{
			Name:  "quarter second delay",
			Delay: time.Millisecond * 250,
		},

		{
			Name:  "Half second day",
			Delay: time.Millisecond * 500,
		},
	}

	for _, tt := range tests {
		orderFinished = []string{}

		eatTime = tt.Delay
		sleepTime = tt.Delay
		thinkTime = tt.Delay

		dine()

		if len(orderFinished) > 5 {
			t.Errorf("%s; Incorrect length of slice; expected 5 got %d", tt.Name, len(orderFinished))
		}

	}
}
