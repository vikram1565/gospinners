package main

import (
	spinner "go-spinners/spinners"
)

func main() {
	si := spinner.New("ArrowSpinner", 5)
	si.StartSpinner()
}
