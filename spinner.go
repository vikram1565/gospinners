package gospinners

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"
)

// spinner - spinner struct
type spinner struct {
	SpinnerName     string
	SpinnerInterval int
	SpinnerDuration time.Duration
	SpinnerParts    []string
}

// SpinnerInfo - SpinnerInfo struct
type SpinnerInfo struct {
	spinner spinner
	running bool
	done    context.CancelFunc
}

// getSpinner - getSpinner
func getSpinner(spinnerName string, duration time.Duration) (spinner, error) {
	switch spinnerName {
	case "ArrowSpinner":
		return spinner{
			SpinnerName:     "ArrowSpinner", // "ArrowSpinner"
			SpinnerInterval: 50,             // interval
			SpinnerDuration: duration,
			SpinnerParts:    []string{`←`, `↖`, `↑`, `↗`, `→`, `↘`, `↓`, `↙`},
		}, nil
	case "DashSpinner":
		return spinner{
			SpinnerName:     "DashSpinner", // "DashSpinner"
			SpinnerInterval: 50,            // interval
			SpinnerDuration: duration,
			SpinnerParts:    []string{`-`, `\`, `/`},
		}, nil
	}
	return spinner{}, errors.New("Invalid spinner name")
}

// New spinner
func New(spinnerName string, duration time.Duration) *SpinnerInfo {
	if spinnerName == "" {
		log.Fatal("Spinner name is required")
	}
	if duration == 0 {
		duration = 5
	}
	s, err := getSpinner(spinnerName, duration)
	if err != nil {
		log.Fatal(err)
	}
	l := SpinnerInfo{
		spinner: s,
	}
	return &l
}

// StartSpinner - start spinner
func (l *SpinnerInfo) StartSpinner() {
	if !l.running {
		// ctx, done := context.WithCancel(context.Background()) // WithCancel
		ctx, done := context.WithTimeout(context.Background(), time.Second*l.spinner.SpinnerDuration) // withtimeout
		l.done = done
		l.running = true
		// call to print spinner
		l.printSpinner(ctx)
	}
}

// printSpinner - print spinner
func (l *SpinnerInfo) printSpinner(ctx context.Context) {
	t := time.NewTicker(time.Duration(l.spinner.SpinnerInterval) * time.Millisecond)
	n := 0
	for {
		select {
		case <-ctx.Done():
			t.Stop() // stop : to release associated resources
			if l.done != nil {
				l.done()
			}
			l.running = false
			return
		case <-t.C:
			fmt.Printf("\r%s", l.spinner.SpinnerParts[n%len(l.spinner.SpinnerParts)]+" loading")
			n++
		}
	}
}
