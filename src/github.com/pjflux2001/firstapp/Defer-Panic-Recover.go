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

	// defer isn't good for loops

	//
	a := "start"
	defer fmt.Println(a)
	a = "end"

	// panic
	// a, b := 1, 0
	// ans := a / b
	// fmt.Println(ans)
	// panic("error happened")

	// panic in web server
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Hello Go!"))
	// })
	// err = http.ListenAndServe(":8080", nil)
	// if err != nil {
	// 	panic(err.Error())
	// }

	// panics happen after defer
	// normal -> defer -> panic

	// anonyonomous fnc
	fmt.Println("Start")
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error : ", err)
		}
	}()
	panic("something bad")
	fmt.Println("done")
}
