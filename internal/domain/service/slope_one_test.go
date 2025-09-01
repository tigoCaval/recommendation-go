package service_test

import (
	"reflect"
	"testing"

	"github.com/tigoCaval/recommendation-go/internal/domain/model"
	"github.com/tigoCaval/recommendation-go/internal/domain/service"
)

func TestRecommend_SlopeOneWhenMinScoreIsZero(t *testing.T) {
	client := service.NewSlopeOne()

	data := []model.Transaction{
		{ProductID: "A", Score: 1, UserID: "John"},
		{ProductID: "B", Score: 0, UserID: "John"},
		{ProductID: "A", Score: 0, UserID: "Mary"},
		{ProductID: "B", Score: 1, UserID: "Mary"},
		{ProductID: "C", Score: 1, UserID: "James"},
		{ProductID: "A", Score: 1, UserID: "James"},
		{ProductID: "A", Score: 1, UserID: "Bob"},
		{ProductID: "B", Score: 0, UserID: "Luke"},
		{ProductID: "C", Score: 1, UserID: "Bob"},
		{ProductID: "G", Score: 1, UserID: "John"},
		{ProductID: "A", Score: 1, UserID: "Ryan"},
		{ProductID: "B", Score: 1, UserID: "Betty"},
		{ProductID: "C", Score: 0, UserID: "Ryan"},
		{ProductID: "G", Score: 1, UserID: "Mary"},
		{ProductID: "F", Score: 1, UserID: "Betty"},
		{ProductID: "B", Score: 0, UserID: "James"},
		{ProductID: "F", Score: 1, UserID: "John"},
		{ProductID: "C", Score: 1, UserID: "Laura"},
		{ProductID: "F", Score: 1, UserID: "Laura"},
		{ProductID: "B", Score: 0, UserID: "Laura"},
		{ProductID: "B", Score: 1, UserID: "Ryan"},
	}

	expected := map[string]float64{"C": 0.57}

	result := client.Recommend(data, "John", 0)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}

}

func TestRecommend_SlopeOneWhenUserHasNoRatings(t *testing.T) {
	client := service.NewSlopeOne()

	data := []model.Transaction{
		{UserID: "John", ProductID: "A", Score: 5},
		{UserID: "John", ProductID: "B", Score: 3},
	}

	result := client.Recommend(data, "Steve", 0)
	if len(result) != 0 {
		t.Errorf("expected empty result for user with no ratings, got %v", result)
	}
}

func TestRecommend_SlopeOneWithMinScore(t *testing.T) {
	client := service.NewSlopeOne()

	data := []model.Transaction{
		// Target user
		{ProductID: "A", Score: 1, UserID: "John"},
		{ProductID: "B", Score: 1, UserID: "John"},

		// Other users
		{ProductID: "A", Score: 1, UserID: "Mary"},
		{ProductID: "B", Score: 0, UserID: "Mary"},
		{ProductID: "C", Score: 1, UserID: "Mary"},
		{ProductID: "C", Score: 1, UserID: "James"},
		{ProductID: "B", Score: 1, UserID: "James"},
		{ProductID: "D", Score: 0, UserID: "James"},
	}

	minScore := 0.1
	expected := map[string]float64{"C": 1.33}

	result := client.Recommend(data, "John", minScore)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}
