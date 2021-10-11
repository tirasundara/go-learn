package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/tirasundara/go-learn/depinject"
	"github.com/tirasundara/go-learn/hello"
)

func main() {
	name := "Tira"

	fmt.Println(hello.Hello(name, "Spanish"))
	depinject.Greet(os.Stdout, name)

	// The Internet
	log.Fatal(http.ListenAndServe(":5000", http.HandlerFunc(depinject.MyGreetHandler)))
}
