package main

import (
	"fmt"
	"os"
	"time"

	"github.com/tirasundara/go-learn/concurrency"
	"github.com/tirasundara/go-learn/depinject"
	"github.com/tirasundara/go-learn/hello"
	"github.com/tirasundara/go-learn/mocks"
)

func main() {
	name := "Tira"

	fmt.Println(hello.Hello(name, "Spanish"))
	depinject.Greet(os.Stdout, name)

	fmt.Println("")

	sleeper := &mocks.ConfigurableSleeper{
		Duration:  1 * time.Second,
		SleepFunc: time.Sleep,
	}
	mocks.Countdown(os.Stdout, sleeper)

	// The Internet
	// log.Fatal(http.ListenAndServe(":5000", http.HandlerFunc(depinject.MyGreetHandler)))

	fmt.Println()

	// CheckWebsites
	urls := []string{
		"http://www.google.com",
		"http://www.tirasundara.com",
	}
	CheckWebResults := concurrency.CheckWebsites(concurrency.CheckWebsite, urls)
	webStatus := map[bool]string{
		true:  "up",
		false: "down",
	}

	for url, status := range CheckWebResults {
		fmt.Printf("%s is %s\n", url, webStatus[status])
	}

}
