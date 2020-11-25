package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {

	interval, err := strconv.ParseInt(os.Args[1], 10, 0)

	if err != nil {
		fmt.Println("First parameter must be number")
		os.Exit(1)
	}

	notes := []string{"a", "b", "c", "d", "e", "f", "g"}

	fmt.Printf("Starting note training, with interval of %v second\n", interval)

	for {
		fmt.Printf("Play: %v\n", notes[rand.Intn(7)])
		time.Sleep(time.Duration(interval) * time.Second)
	}

}
