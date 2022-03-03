package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(quoteHandler)
	request, err := http.NewRequest(http.MethodGet, "/fizzbuzz?count=20", nil)
	if err != nil {
		t.Fatal(err)
	}
	handler.ServeHTTP(recorder, request)
	status := recorder.Code
	if status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var slice []string
	slice = append(slice, "1")
	slice = append(slice, "2")
	slice = append(slice, fizz)
	slice = append(slice, "4")
	slice = append(slice, buzz)
	slice = append(slice, fizz)
	slice = append(slice, "7")
	slice = append(slice, "8")
	slice = append(slice, fizz)
	slice = append(slice, buzz)
	slice = append(slice, "11")
	slice = append(slice, fizz)
	slice = append(slice, "13")
	slice = append(slice, "14")
	slice = append(slice, fizz)
	slice = append(slice, "16")
	slice = append(slice, "17")
	slice = append(slice, fizz)
	slice = append(slice, "19")
	slice = append(slice, buzz)
	var expected FizzBuzz
	expected.Fizzbuzz = slice
	var result FizzBuzz
	err = json.NewDecoder(recorder.Body).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("we got %s result should be %s", result, expected)
	}
}
func Test400(t *testing.T) {
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(quoteHandler)
	request, err := http.NewRequest(http.MethodGet, "/fizzbuzz", nil)
	if err != nil {
		t.Fatal(err)
	}
	handler.ServeHTTP(recorder, request)
	status := recorder.Code
	if status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}
}
func Test404(t *testing.T) {
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(quoteHandler)
	request, err := http.NewRequest(http.MethodGet, "/fizzbuzzXX?count=20", nil)
	if err != nil {
		t.Fatal(err)
	}
	handler.ServeHTTP(recorder, request)
	status := recorder.Code
	if status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusNotFound)
	}
}
func Test501(t *testing.T) {
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(quoteHandler)
	request, err := http.NewRequest(http.MethodPost, "/fizzbuzz?count=20", nil)
	if err != nil {
		t.Fatal(err)
	}
	handler.ServeHTTP(recorder, request)
	status := recorder.Code
	if status != http.StatusNotImplemented {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusNotImplemented)
	}
}
