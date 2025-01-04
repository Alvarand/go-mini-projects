package models

import (
	"log"
)

// CheckZeroArgs - checking number of args
func CheckZeroArgs(args []string) {
	if len(args) < 1 {
		log.Fatal(ErrorZeroLenArgs)
	}
}
