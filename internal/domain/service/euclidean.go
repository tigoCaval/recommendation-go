package service

import (
	"math"

	"github.com/tigoCaval/recommendation-go/internal/domain/model"
	"github.com/tigoCaval/recommendation-go/internal/domain/ports"
)

type Euclidean struct{}

func NewEuclidean() ports.Collaborative {
	return &Euclidean{}
}

// Recommend generates recommendations for a user using optimized Euclidean distance
func (e *Euclidean) Recommend(table []model.Transaction, user string, score float64) map[string]float64 {
	// 1. Build a quick lookup map: user -> product -> score
	userScores := map[string]map[string]float64{}
	for _, t := range table {
		if userScores[t.UserID] == nil {
			userScores[t.UserID] = map[string]float64{}
		}
		userScores[t.UserID][t.ProductID] = t.Score
	}

	myProducts := userScores[user]

	// 2. Calculate similarity with other users
	similarity := map[string]float64{}
	for otherUser, otherScores := range userScores {
		if otherUser == user {
			continue
		}
		sum := 0.0
		matched := false
		for pid, myScore := range myProducts {
			if otherScore, ok := otherScores[pid]; ok && myScore >= score && otherScore >= score {
				diff := myScore - otherScore
				sum += diff * diff
				matched = true
			}
		}
		if matched {
			similarity[otherUser] = math.Round((1/(1+math.Sqrt(sum)))*100) / 100
		}
	}

	// 3. Compute weighted average for each product
	result := map[string]float64{}
	weightSum := map[string]float64{}
	for otherUser, sim := range similarity {
		for pid, otherScore := range userScores[otherUser] {
			if _, ok := myProducts[pid]; ok {
				continue // skip already rated products
			}
			result[pid] += sim * otherScore
			weightSum[pid] += sim
		}
	}

	// 4. Normalize the results and apply minScore filter
	finalResult := map[string]float64{}
	for pid, val := range result {
		if weightSum[pid] > 0 {
			normalized := math.Round((val/weightSum[pid])*100) / 100
			if normalized >= score {
				finalResult[pid] = normalized
			}
		}
	}

	return finalResult
}
