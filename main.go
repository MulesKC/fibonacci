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
			fmt.Println("next route hit")
			curPos++
			if curPos >= len(fibonacciValues) {
				fmt.Println("making new num")
				fibonacciValues = newFibNum(fibonacciValues)
			}
			w.Write([]byte(strconv.Itoa(fibonacciValues[curPos])))

		case "current":
			fmt.Println("current route hit")
			w.Write([]byte(strconv.Itoa(fibonacciValues[curPos])))

		case "last":
			fmt.Println("last route hit")
			curPos--
			w.Write([]byte(strconv.Itoa(fibonacciValues[curPos])))

		default:
			fmt.Println("error, no matching endpoint")
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("error, no matching endpoint found"))
		}
		return
	}
}

func newFibNum(fibonacciValues []int) []int {

	newNumToAdd := fibonacciValues[len(fibonacciValues)-1] + fibonacciValues[len(fibonacciValues)-2]
	fibonacciValues = append(fibonacciValues, newNumToAdd)

	return fibonacciValues
}
