package main

import (
	"log"
	"flag"
	"fmt"
	"time"
	"os"
)

	
func check(e error) {
	if e != nil {
			panic(e)
	}
}

func ensure_dir(dir string) {
	_, err := os.Stat(dir)
 
	if os.IsNotExist(err) {
		errDir := os.MkdirAll(dir, 0755)
		if errDir != nil {
			log.Fatal(err)
		}
 
	}
}


func main() {

	weekday := time.Now().Weekday()
	current_time := time.Now().Format(time.Kitchen)
	year, week_number := time.Now().UTC().ISOWeek()

	var message = flag.String("message", "<message goes here>", "Todo item")
	flag.Parse()


	var title = fmt.Sprintf("Week %d %d Notes\n", week_number, year)

	fmt.Printf(title)
	fmt.Println("Weekday: ", weekday)
	fmt.Printf("%s @ %s\n", *message, current_time)

	ensure_dir("files")


	// Write file 
	f, err := os.Create("files/sample.md")
	check(err)

	// ensure Close is performed later in the program’s execution, cleanup
	defer f.Close()

}
