package main

import (
	"log"
	"os"
	"reflect"
	"strconv"
)

func main() {
	if arg() {
		log.Println(checkEvenOdd(os.Args[1]))
	}
}
func checkEvenOdd(param string) string {
	switch reflect.TypeOf(param).String() {
	case "string":
		if inputNumber, _ := strconv.Atoi(os.Args[1]); inputNumber > 0 {
			if Even(inputNumber) {
				return "positive even"
			}
			return "positive odd"
		} else if inputNumber < 0 {
			if Even(inputNumber) {
				return "positive even"
			}
			return "positive odd"
		} else if inputNumber == 0 {
			return "zero"
		}
		break
	}
	return "NaN"
}
func arg() bool {
	if len(os.Args) < 2 {
		log.Println("Missing argument")
		os.Exit(1)
	}
	return true
}
func Even(n int) bool {
	return n%2 == 0
}
