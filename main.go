// fibonacci-test
package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello World!")

	//	fibValues := []int
	r := mux.NewRouter()
	r.HandleFunc("/next", next)
	r.HandleFunc("/current", current)
	r.HandleFunc("/last", last)
	http.Handle("/", r)

	http.ListenAndServe(":8080", nil)
	fmt.Println("Listening on 8080")

}

func next(w http.ResponseWriter, r *http.Request) {
	fmt.Println("next")
}

func current(w http.ResponseWriter, r *http.Request) {
	fmt.Println("current")
}

func last(w http.ResponseWriter, r *http.Request) {
	fmt.Println("last")
}
