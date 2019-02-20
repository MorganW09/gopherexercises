package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

type QuestionAndAnswer struct {
	Question string
	Answer   int
}

func RunGame(questionAndAnswers []QuestionAndAnswer, timeLimit int) {
	fmt.Println(timeLimit)
	timer := time.NewTimer(time.Duration(timeLimit) * time.Second)
	correct := 0
problemloop:
	for _, qa := range questionAndAnswers {
		fmt.Printf("%s: ", qa.Question)
		answerCh := make(chan string)

		go func() {
			var input string
			fmt.Scanln(&input)
			answerCh <- input
		}()

		select {
		case <-timer.C:
			fmt.Println()
			break problemloop
		case answer := <-answerCh:
			guess, err := strconv.Atoi(answer)
			if err != nil {
				// handle error
				fmt.Println(err)
			} else {
				if guess == qa.Answer {
					correct++
				}
				fmt.Println(guess)
			}
		}
	}

	fmt.Printf("You scored %d out of %d.\n", correct, len(questionAndAnswers))
}

func ReadInCsv(csvName string) []QuestionAndAnswer {
	csvFile, _ := os.Open(csvName)
	reader := csv.NewReader(bufio.NewReader(csvFile))
	var questionAndAnswers []QuestionAndAnswer
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		answer, _ := strconv.Atoi(line[1])
		questionAndAnswers = append(questionAndAnswers, QuestionAndAnswer{
			Question: line[0],
			Answer:   answer,
		})
	}
	return questionAndAnswers
}

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	timeLimit := flag.Int("limit", 30, "the time limit for the quiz in seconds")
	flag.Parse()

	fmt.Println(*csvFilename)
	fmt.Println(*timeLimit)
	questionAndAnswers := ReadInCsv(*csvFilename)
	RunGame(questionAndAnswers, *timeLimit)
}
