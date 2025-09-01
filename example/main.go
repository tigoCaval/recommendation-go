package main

import (
	"fmt"

	"github.com/tigoCaval/recommendation-go/recommendation"
)

func main() {
	table := []recommendation.Transaction{
		{ProductID: "A", Score: 1, UserID: "Pedro"},
		{ProductID: "B", Score: 0, UserID: "Pedro"},
		{ProductID: "A", Score: 0, UserID: "Maria"},
		{ProductID: "B", Score: 1, UserID: "Maria"},
	}

	client := recommendation.NewRecommend()
	result := client.Ranking(table, "Pedro", 0)
	fmt.Println(result)
}
