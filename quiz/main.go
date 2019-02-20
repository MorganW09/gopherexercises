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
)

type QuestionAndAnswer struct {
	Question string
	Answer   int
}

func RunGame(questionAndAnswers []QuestionAndAnswer) {
	correct := 0
	for _, qa := range questionAndAnswers {
		fmt.Printf("%s: ", qa.Question)
		var input string
		fmt.Scanln(&input)

		guess, err := strconv.Atoi(input)
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
	questionAndAnswers := ReadInCsv(*csvFilename)
	RunGame(questionAndAnswers)
}
