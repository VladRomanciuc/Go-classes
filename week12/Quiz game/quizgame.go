/*

Part 1
Create a program that will read in a quiz provided via a CSV file (more details below) and will then give the quiz to a user
keeping track of how many questions they get right and how many they get incorrect.
Regardless of whether the answer is correct or wrong, the next question should be asked immediately afterward.

The CSV file should default to problems.csv (example shown below),
but the user should be able to customize the filename via a flag.

The CSV file will be in a format like the below, where the first column is a question and the second column is in the same row
is the answer to that question.

You can assume that quizzes will be relatively short (< 100 questions) and will have single word/number answers.

At the end of the quiz, the program should output the total number of questions correct and how many questions there were in total.
Questions given invalid answers are considered incorrect.

NOTE: CSV files may have questions with commas in them. Eg: "what 2+2, sir?",4 is a valid row in a CSV.
I suggest you look into the CSV package in Go and don't try to write your own CSV parser.


Part 2

Adapt your program from part 1 to add a timer. The default time limit should be 30 seconds,
but should also be customizable via a flag.

Your quiz should stop as soon as the time limit has been exceeded.
That is, you shouldn't wait for the user to answer one final question but should ideally stop the quiz entirely even
if you are currently waiting on an answer from the end-user.

Users should be asked to press enter (or some other key) before the timer starts,
and then the questions should be printed out to the screen one at a time until the user provides an answer.
Regardless of whether the answer is correct or wrong, the next question should be asked.

At the end of the quiz, the program should still output the total number of questions correct and how many questions
there were in total.
Questions given invalid answers or unanswered are considered incorrect.
*/

package main

import (
	"fmt"
	"flag"
	"os"
	"strings"
	"encoding/csv"
	"time"
)

//Declare a struct for a question
type question struct {
	problem string
	result string
}

//A line parser function that return a slice of questions
func parseLines(lines [][]string) []question {
	
	questions := make([]question, len(lines))
	
	for i, line := range lines {
		questions[i] = question {
			problem:line[0],
			//Empty space trimmer
			result: strings.TrimSpace(line[1]),
		}
	}
	return questions
}

//Simple function to exit the program
func exit(message string) {
	fmt.Println(message)
	os.Exit(1)
}

//The function will open the file, read it and parse each line and return a slice
func fileReader (csvFilename string) []question {
	//File opener
	file, err := os.Open(csvFilename)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s \n", csvFilename))
	}
	//File reader
	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	if err != nil {
		exit("Failed to parse the provided CSV file.")
	}

	//Uses the parseLines function to return a slice of questions
	questions := parseLines(lines)

	return questions
}

//The main logic of the program
func main() {
	//Declare 2 flags(for the csv file and default time limit)
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of `question, answer`")
	timeLimit := flag.Int("limit", 20, "the time limit for the quiz in seconds")
	flag.Parse()
	
	//Initialise variables to be used and a channel to recieve the answers
	questions := fileReader(*csvFilename)
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	correct := 0
	answerCh := make(chan string)

	//Print a message to start the program
	fmt.Printf("The quiz has %v questions and you have 20 seconds to solve it.\n", len(questions))

	//Starting the loop
	for index, question := range questions {
		fmt.Printf("%d. Can you solve this: %s = ", index+1, question.problem)
		
		//Go routine to scan the inputs and send the result to channel
		go func(){
			var entry string
			fmt.Scanf("%s\n", &entry)
			answerCh <- entry
		}()

		//Options to end the program
		select {
			//Case run out of time
		case <- timer.C:
			fmt.Printf("\nYour score is %c out of %d.\n", correct, len(questions))
			return
			//Case right answer to a question
		case answer := <- answerCh:
			if answer == question.result {
			fmt.Printf("Your are right!\n")
			correct++
			}
		}
	}
}