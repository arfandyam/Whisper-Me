package main

import (
	"flag"
	"fmt"
	"strconv"
	"github.com/arfandyam/Whisper-Me/initializers"
	"github.com/arfandyam/Whisper-Me/seeders/tables"
	"github.com/google/uuid"
)


func init() {
	initializers.LoadEnvVariables()
	initializers.ConnDB()
}

func main() {
	tableFlag := flag.String("table", "", "specify table to seed")
	userIdFlag := flag.String("userId", "", "specify userId related to question")
	questionIdFlag := flag.String("questionId", "", "specify questionId related to responses")
	amountFlag := flag.String("amount", "1", "specify amount of data to seed")

	flag.Parse()

	amount, err := strconv.Atoi(*amountFlag)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// options := tables.Options{
	// 	Amount: amount,
	// 	UserId: uuid.MustParse(*userIdFlag),
	// }
	switch {
	case *tableFlag == "questions":
		tables.AddQuestions(initializers.DB, tables.Options{
			Amount: amount,
			UserId: uuid.MustParse(*userIdFlag),
		})
	case *tableFlag == "responses":
		tables.AddResponses(initializers.DB, tables.Options{
			Amount: amount,
			QuestionId: uuid.MustParse(*questionIdFlag),
		})
	default:
		fmt.Println("your argument not found")
	}
}
