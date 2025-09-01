package service

import (
	"github.com/tigoCaval/recommendation-go/internal/domain/model"
	"github.com/tigoCaval/recommendation-go/internal/domain/ports"
)

type RankingCollaborative struct {
	product []model.Transaction
	other   []model.Transaction
}

func NewRanking() ports.Collaborative {
	return &RankingCollaborative{}
}

// Recommend generates recommendations for a target user with a minimum score
func (r *RankingCollaborative) Recommend(table []model.Transaction, user string, score float64) map[string]float64 {
	data := r.addRating(table, user, score)
	return r.filterRating(data)
}

// split separates the transactions of the target user from other users
func (r *RankingCollaborative) split(table []model.Transaction, user string) {
	r.product = nil
	r.other = nil
	for _, item := range table {
		if item.UserID == user {
			r.product = append(r.product, item)
		} else {
			r.other = append(r.other, item)
		}
	}
}

// similarUser counts how many times other users have rated the same products with the same score
func (r *RankingCollaborative) similarUser(table []model.Transaction, user string) map[string]int {
	r.split(table, user)
	similar := make(map[string]int)
	for _, myProduct := range r.product {
		for _, other := range r.other {
			if myProduct.ProductID == other.ProductID && myProduct.Score == other.Score {
				similar[other.UserID]++
			}
		}
	}
	return similar
}

// addRating calculates a ranking score for products based on similar users
func (r *RankingCollaborative) addRating(table []model.Transaction, user string, score float64) map[string]float64 {
	similar := r.similarUser(table, user)
	rank := make(map[string]float64)
	for _, other := range r.other {
		for userID, value := range similar {
			// only consider other users who are similar and whose score exceeds the threshold
			if other.UserID == userID && other.Score > score {
				rank[other.ProductID] += float64(value)
			}
		}
	}
	return rank
}

// filterRating removes products already rated by the target user
func (r *RankingCollaborative) filterRating(rank map[string]float64) map[string]float64 {
	filtered := make(map[string]float64)
	// map of products already rated by the target user
	myProducts := make(map[string]bool)
	for _, item := range r.product {
		myProducts[item.ProductID] = true
	}

	for productID, value := range rank {
		if !myProducts[productID] { // only add if not already rated
			filtered[productID] = value
		}
	}

	return filtered
}
