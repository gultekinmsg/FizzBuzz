package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

const (
	fizz = "fizz"
	buzz = "buzz"
)

func main() {
	http.HandleFunc("/fizzbuzz", quoteHandler)
	err := http.ListenAndServe(":8080", nil);
	if  err != nil {
		log.Fatal(err)
	}
}
func quoteHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/fizzbuzz" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	keys, ok := r.URL.Query()["count"]
	if !ok || len(keys) == 0 {
		http.Error(w, "400 bad request.", http.StatusBadRequest)
		return
	}
	if r.Method == http.MethodGet {
		count, err := strconv.Atoi(keys[0])
		if err != nil {
			log.Fatal(err)
		}
		answer := fizzBuzz(count)
		output, err := json.Marshal(answer)
		if err != nil {
			log.Fatal(err)
		}
		w.Header().Set("content-type", "application/json")
		_, err = w.Write(output)
		if err != nil {
			log.Fatal(err)
		}

	} else {
		http.Error(w, "501 not implemented.", http.StatusNotImplemented)
		return
	}
}
func fizzBuzz(count int) FizzBuzz {
	var slice []string
	for i := 1; i <= count; i++ {
		if i%3 == 0 {
			slice = append(slice, fizz)
		} else if i%5 == 0 {
			slice = append(slice, buzz)
		} else {
			slice = append(slice, strconv.Itoa(i))
		}
	}
	return FizzBuzz{
		Fizzbuzz: slice,
	}

}
