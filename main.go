package main

import (
	"fmt"
	"os"

	"github.com/tirasundara/go-learn/depinject"
	"github.com/tirasundara/go-learn/hello"
	"github.com/tirasundara/go-learn/mocks"
)

func main() {
	name := "Tira"

	fmt.Println(hello.Hello(name, "Spanish"))
	depinject.Greet(os.Stdout, name)

	fmt.Println("")

	sleeper := &mocks.DefaultSleeper{}
	mocks.Countdown(os.Stdout, sleeper)

	// The Internet
	// log.Fatal(http.ListenAndServe(":5000", http.HandlerFunc(depinject.MyGreetHandler)))
}
