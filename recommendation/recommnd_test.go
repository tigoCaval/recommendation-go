package recommendation_test

import (
	"reflect"
	"testing"

	"github.com/tigoCaval/recommendation-go/recommendation"
)

func TestRanking(t *testing.T) {
	data := []recommendation.Transaction{
		{ProductID: "A", Score: 1, UserID: "John"},
		{ProductID: "B", Score: 1, UserID: "John"},
		{ProductID: "A", Score: 1, UserID: "Mary"},
		{ProductID: "B", Score: 0, UserID: "Mary"},
		{ProductID: "C", Score: 1, UserID: "Mary"},
	}

	expected := map[string]float64{
		"C": 1,
	}

	client := recommendation.NewRecommend()
	result := client.Ranking(data, "John", 0)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestEuclidean(t *testing.T) {
	data := []recommendation.Transaction{
		{ProductID: "A", Score: 1, UserID: "John"},
		{ProductID: "B", Score: 1, UserID: "John"},
		{ProductID: "A", Score: 1, UserID: "Mary"},
		{ProductID: "B", Score: 0, UserID: "Mary"},
		{ProductID: "C", Score: 1, UserID: "Mary"},
	}

	expected := map[string]float64{
		"C": 1,
	}

	client := recommendation.NewRecommend()
	result := client.Euclidean(data, "John", 0)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestSlopeOne(t *testing.T) {
	data := []recommendation.Transaction{
		{ProductID: "A", Score: 1, UserID: "John"},
		{ProductID: "B", Score: 1, UserID: "John"},
		{ProductID: "A", Score: 1, UserID: "Mary"},
		{ProductID: "B", Score: 0, UserID: "Mary"},
		{ProductID: "C", Score: 1, UserID: "Mary"},
		{ProductID: "C", Score: 1, UserID: "James"},
	}

	expected := map[string]float64{
		"C": 1.5,
	}

	client := recommendation.NewRecommend()
	result := client.SlopeOne(data, "John", 0)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}
