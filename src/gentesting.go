package main

import (
	"log"
	"os"
	"strconv"
)

func main() {

	count, _ := strconv.Atoi(os.Getenv("COUNT"))
	showlogs := os.Getenv("SHOW_LOGS")

	if count == 0 {
		count = 500
	}
	log.Print(showlogs)

	if showlogs == "false" || showlogs == "FALSE" {
		log.Print("logging is disabled")
		return
	}
	for i := 0; i < count; i++ {
		log.Print(i)
	}
}
