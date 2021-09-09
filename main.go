package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {

	weekday := time.Now().Weekday()
	current_time := time.Now().Format(time.Kitchen)
	year, week_number := time.Now().UTC().ISOWeek()

	var message = flag.String("message", " ", "Todo item")
	flag.Parse()


	var title = fmt.Sprintf("Week %d %d Notes\n", week_number, year)

	fmt.Printf(title)
	fmt.Println("Weekday: ", weekday)
	fmt.Printf("%s @ %s\n", *message, current_time)
}
