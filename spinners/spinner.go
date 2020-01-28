package main

import (
	"context"
	"fmt"
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
	// out     io.Writer
	running bool
	done    context.CancelFunc
}

func main() {
	// getSpinner - get spinner
	// s := getSpinner()
	// o := os.Stdout
	// // assign values to SpinnerInfo
	// l := &SpinnerInfo{
	// 	spinner: s,
	// 	out:     o,
	// }
	// // start the spinner
	// l.startSpinner()
}

// getSpinner - getSpinner
func getSpinner(spinnerName string, duration time.Duration) spinner {
	switch spinnerName {
	case "ArrowSpinner":
		return spinner{
			SpinnerName:     "ArrowSpinner", // "ArrowSpinner"
			SpinnerInterval: 50,             // add interval
			SpinnerDuration: duration,
			SpinnerParts:    []string{`←`, `↖`, `↑`, `↗`, `→`, `↘`, `↓`, `↙`},
		}
	}
	return spinner{}
}

// New spinner
func New(spinnerName string, duration time.Duration) *SpinnerInfo {
	s := getSpinner(spinnerName, duration)
	// o := os.Stdout
	l := SpinnerInfo{
		spinner: s,
		// out:     o,
	}
	return &l
}

// startSpinner - start spinner
func (l *SpinnerInfo) startSpinner() {
	if !l.running {
		// ctx, done := context.WithCancel(context.Background()) // WithCancel
		ctx, done := context.WithTimeout(context.Background(), time.Second*5) // withtimeout
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
			return
		case <-t.C:
			// part := "\r" + l.spinner.SpinnerParts[n%len(l.spinner.SpinnerParts)] + " loading"
			// fmt.Fprint(l.out, part)
			fmt.Printf("\r%s", l.spinner.SpinnerParts[n%len(l.spinner.SpinnerParts)]+" loading")
			n++
		}
	}
}

// Stop - stop is ticker struct method
func (l *SpinnerInfo) Stop() {
	if l.done != nil {
		l.done()
	}
	l.running = false
}
