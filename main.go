package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {

	weekday := time.Now().Weekday()
	current_time := time.Now().Format(time.Kitchen)

	var message = flag.String("message", " ", "Todo item")
	flag.Parse()

	fmt.Println("Weekday: ", weekday)
	fmt.Printf("%s @ %s\n", *message, current_time)
}
