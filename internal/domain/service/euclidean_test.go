package service_test

import (
	"reflect"
	"testing"

	"github.com/tigoCaval/recommendation-go/internal/domain/model"
	"github.com/tigoCaval/recommendation-go/internal/domain/service"
)

func TestRecommend_EuclideanWhenMinScoreIsZero(t *testing.T) {
	client := service.NewEuclidean()
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

	tests := []struct {
		name     string
		user     string
		expected map[string]float64
	}{
		{"John", "John", map[string]float64{"C": 0.86}},
		{"Mary", "Mary", map[string]float64{"F": 1, "C": 0.74}},
		{"James", "James", map[string]float64{"F": 1, "G": 1}},
		{"Bob", "Bob", map[string]float64{"F": 1, "G": 1, "B": 0.25}},
		{"Luke", "Luke", map[string]float64{"A": 0.83, "C": 0.8, "F": 1, "G": 1}},
		{"Ryan", "Ryan", map[string]float64{"G": 1, "F": 1}},
		{"Betty", "Betty", map[string]float64{"A": 0.67, "G": 1, "C": 0.5}},
		{"Laura", "Laura", map[string]float64{"A": 0.87, "G": 1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := client.Recommend(data, tt.user, 0)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestRecommend_EuclideanWhenUserHasNoRatings(t *testing.T) {
	client := service.NewEuclidean()

	data := []model.Transaction{
		{ProductID: "A", Score: 1, UserID: "John"},
		{ProductID: "B", Score: 0, UserID: "John"},
		{ProductID: "A", Score: 0, UserID: "Mary"},
		{ProductID: "B", Score: 1, UserID: "Mary"},
		{ProductID: "C", Score: 1, UserID: "James"},
	}

	// User "Steve" does not exist in the dataset
	result := client.Recommend(data, "Steve", 0)
	if len(result) != 0 {
		t.Errorf("expected empty result for user with no ratings, got %v", result)
	}
}

func TestRecommend_EuclideanWhenUserRatedAllProducts(t *testing.T) {
	client := service.NewEuclidean()

	// Create a dataset where user "Alice" has rated all products
	data := []model.Transaction{
		{ProductID: "A", Score: 1, UserID: "Alice"},
		{ProductID: "B", Score: 1, UserID: "Alice"},
		{ProductID: "C", Score: 1, UserID: "Alice"},
		{ProductID: "A", Score: 0, UserID: "John"},
		{ProductID: "B", Score: 1, UserID: "Mary"},
		{ProductID: "C", Score: 1, UserID: "Bob"},
	}

	result := client.Recommend(data, "Alice", 0)
	if len(result) != 0 {
		t.Errorf("expected empty result for user who rated all products, got %v", result)
	}
}

func TestRecommend_EuclideanWhenMinScoreIsHigh(t *testing.T) {
	client := service.NewEuclidean()

	data := []model.Transaction{
		// Target user
		{ProductID: "A", Score: 1, UserID: "John"},
		{ProductID: "B", Score: 1, UserID: "John"},

		// Other users
		{ProductID: "A", Score: 1, UserID: "Mary"},  // similar to John on A
		{ProductID: "B", Score: 0, UserID: "Mary"},  // less similar on B
		{ProductID: "C", Score: 1, UserID: "Mary"},  // new product for John
		{ProductID: "C", Score: 1, UserID: "James"}, // new product for John
		{ProductID: "B", Score: 1, UserID: "James"}, // similar on B
		{ProductID: "D", Score: 0, UserID: "James"}, // low score product
	}

	minScore := 0.1
	expected := map[string]float64{"C": 1}

	result := client.Recommend(data, "John", minScore)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}
