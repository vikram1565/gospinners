package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"time"
)

// spinner - spinner struct
type spinner struct {
	SpinnerName     string
	SpinnerInterval int
	SpinnerParts    []string
}

// spinnerInfo - spinnerInfo struct
type spinnerInfo struct {
	spinner spinner
	out     io.Writer
	running bool
	done    context.CancelFunc
}

func main() {
	// getSpinner - get spinner
	s := getSpinner()
	o := os.Stdout
	// assign values to spinnerInfo
	l := &spinnerInfo{
		spinner: s,
		out:     o,
	}
	// start the spinner
	l.startSpinner()
}

// getSpinner - getSpinner
func getSpinner() spinner {
	return spinner{
		SpinnerName:     "ArrowSpinner",
		SpinnerInterval: 50, // add interval
		SpinnerParts:    []string{`←`, `↖`, `↑`, `↗`, `→`, `↘`, `↓`, `↙`},
	}
}

// startSpinner - start Spinner
func (l *spinnerInfo) startSpinner() {
	if !l.running {
		ctx, done := context.WithCancel(context.Background()) // WithCancel
		// ctx, done := context.WithTimeout(context.Background(), time.Second*10) // withtimeout
		l.done = done
		l.running = true
		// call to print spinner
		l.printSpinner(ctx)
	}
}

// printSpinner - print spinner
func (l *spinnerInfo) printSpinner(ctx context.Context) {
	t := time.NewTicker(time.Duration(l.spinner.SpinnerInterval) * time.Millisecond)
	n := 0
	for {
		select {
		case <-ctx.Done():
			t.Stop()       // stop : to release associated resources
			return
		case <-t.C:
			part := "\r" + l.spinner.SpinnerParts[n%len(l.spinner.SpinnerParts)] + " loading"
			fmt.Fprint(l.out, part)
			n++
		}
	}
}

// Stop - stop is ticker struct method
func (l *spinnerInfo) Stop() {
	if l.done != nil {
		l.done()
	}
	l.running = false
}
