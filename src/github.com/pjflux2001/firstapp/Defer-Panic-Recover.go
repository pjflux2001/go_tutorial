package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Start")
	defer fmt.Println("Middle")
	fmt.Println("end")

	// last defer is executed first after synchronous program is done
	// FILO and LIFO
	defer fmt.Println("Middle1")
	defer fmt.Println("Middle2")
	defer fmt.Println("Middle3")

	// defers are executed towards the end

	res, err := http.Get("http://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	// closing resource before using it in body but with defer
	defer res.Body.Close()
	robots, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", robots)
}
