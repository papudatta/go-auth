package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Person struct {
	First string
}

func main() {
	//	p1 := Person{
	//		First: "Julie",
	//	}
	//
	//	p2 := Person{
	//		First: "Rilang",
	//	}
	//
	//	// Marshal
	//	xp := []Person{p1, p2}
	//	bs, err := json.Marshal(xp)
	//	if err != nil {
	//		log.Panic(err)
	//	}
	//	fmt.Println("Print json: ", string(bs))
	//
	//	// Now unmarshalling
	//	xp2 := []Person{}
	//	err = json.Unmarshal(bs, &xp2)
	//	if err != nil {
	//		log.Panic(err)
	//	}
	//
	//	fmt.Println("back to Go data structure", xp2)

	http.HandleFunc("/encode", foo)
	http.HandleFunc("/decode", bar)
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	p1 := Person{
		First: "Julie",
	}

	// For sending a slice
	// people := []Person{p1, p2}
	// json.NewEncoder(w).Encode(people)

	err := json.NewEncoder(w).Encode(p1)
	if err != nil {
		log.Println("Encode bad data ", err)
	}
}

func bar(w http.ResponseWriter, r *http.Request) {
	var p2 Person
	// for a slice, do
	// people := []Person{}
	// err := json.NewDecoder(r.Body).Decode(&people)
	err := json.NewDecoder(r.Body).Decode(&p2)
	if err != nil {
		log.Println("Decoded bad data ", err)
	}
	log.Println("Person: ", p2)
}
