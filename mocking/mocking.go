package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Sleeper interface {
	Sleep()
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (s *ConfigurableSleeper) Sleep() {
	s.sleep(s.duration)
}

const finalWord = "GO!"
const countdownStart = 3

func Countdown(buffer io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(buffer, i)
		sleeper.Sleep()
	}
	fmt.Fprint(buffer, finalWord)
}

func main() {
	Countdown(os.Stdout, &ConfigurableSleeper{
		duration: 1 * time.Second,
		sleep:    time.Sleep,
	})
}
