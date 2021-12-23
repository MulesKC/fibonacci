package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {
	fibonacciValues := []int{0, 1}
	curPos := 0
	r := mux.NewRouter()
	r.HandleFunc("/{input}", fibHandler(curPos, fibonacciValues))
	http.Handle("/", r)

	fmt.Println("Listening on 8080")
	http.ListenAndServe(":8080", nil)

}

func fibHandler(curPos int, fibonacciValues []int) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		param := mux.Vars(r)
		switch param["input"] {
		case "next":
			fmt.Println("next")
			if curPos >= len(fibonacciValues) {
				fmt.Println("making new num")
				fibonacciValues = newFibNum(fibonacciValues)
			}

			fmt.Println("curpos = " + strconv.Itoa(curPos))
			fmt.Println("fibNum = " + strconv.Itoa(fibonacciValues[curPos]))
			w.Write([]byte(strconv.Itoa(fibonacciValues[curPos])))
			curPos++
			return
		case "current":
			fmt.Println("current")
			fmt.Println("curpos = " + strconv.Itoa(curPos))
			fmt.Println("fibNum = " + strconv.Itoa(fibonacciValues[curPos]))

			w.Write([]byte(strconv.Itoa(fibonacciValues[curPos])))
			return

		case "last":
			fmt.Println("last")
			curPos--
			fmt.Println("curpos = " + strconv.Itoa(curPos))
			fmt.Println("fibNum = " + strconv.Itoa(fibonacciValues[curPos]))

			w.Write([]byte(strconv.Itoa(fibonacciValues[curPos])))
			return

		default:
		}
	}
}

func newFibNum(fibonacciValues []int) []int {
	var firstValue int
	var secondValue int
	if len(fibonacciValues) < 1 {
		firstValue = 1
		secondValue = 0
	} else if len(fibonacciValues) < 2 {
		firstValue = fibonacciValues[len(fibonacciValues)-1]
		secondValue = 0
	} else {
		firstValue = fibonacciValues[len(fibonacciValues)-1]
		secondValue = fibonacciValues[len(fibonacciValues)-2]
	}
	newNumToAdd := firstValue + secondValue

	fibonacciValues = append(fibonacciValues, newNumToAdd)
	return fibonacciValues
}
