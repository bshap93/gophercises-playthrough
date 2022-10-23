package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {

	var numCorrect int = 0
	var numTotal int = 0

	filename := flag.String("filename", "problems.csv", "the name of the file the quiz is in")

	timeLimit := flag.Int("limit", 3, "the timelimit for the quiz in seconds")

	fmt.Println("Time limit is", *timeLimit, "press enter to begin.")
	fmt.Scan()
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	go func() {
		<-timer.C
		msg := fmt.Sprintf("Ahem,", numCorrect, "correct answers out of", numTotal)
		exitWithMessage(msg)
	}()
	// data, err := os.ReadFile(filename)
	f, err := os.Open(*filename)

	if err != nil {
		panic(err)
	}

	csvReader := csv.NewReader(f)

	data, err := csvReader.ReadAll()
	for range data {
		numTotal++
	}

	if err != nil {
		log.Fatal(err)
	}

	for _, line := range data {
		var userAnswer string
		for j, field := range line {
			if j == 0 {
				println(field)
				fmt.Scan(&userAnswer)
			} else if j == 1 {
				if userAnswer == field {
					numCorrect++
				}
			}
		}
	}
	f.Close()

	msg := fmt.Sprintf("Ahem,", numCorrect, "correct answers out of", numTotal)

	exitWithMessage(msg)

}

func exitWithMessage(message string) {
	fmt.Println(message)
	os.Exit(3)
}
