package main

import (
	"fmt"
	"os"

	"github.com/tirasundara/go-learn/depinject"
	"github.com/tirasundara/go-learn/hello"
)

func main() {
	name := "Tira"

	fmt.Println(hello.Hello(name, "Spanish"))
	depinject.Greet(os.Stdout, name)
}
