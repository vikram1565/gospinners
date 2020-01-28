package main

import (
	spinner "github.com/vikram1565/gospinners"
)

func main() {
	// spinnerName is required. for spinner names please check README.md file.
	// default duration is 5 sec

	si := spinner.New("ArrowSpinner", 0)
	si.StartSpinner()
}
