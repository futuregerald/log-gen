package main

import (
	"log"
	"os"
	"strconv"
)

func main() {

	count, _ := strconv.Atoi(os.Getenv("COUNT"))

	if count == 0 {
		count = 500
	}

	for i := 0; i < count; i++ {
		log.Print(i)
	}
}
