package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Go())
	var s = quote.Glass()
	fmt.Println(s + " hurrdurr")
}
