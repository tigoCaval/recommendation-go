package service

import (
	"math"

	"github.com/tigoCaval/recommendation-go/internal/domain/model"
	"github.com/tigoCaval/recommendation-go/internal/domain/ports"
)

type SlopeOne struct {
	diff   map[string]map[string]float64 // sum of differences between item pairs
	weight map[string]map[string]float64 // number of observations per item pair
}

func NewSlopeOne() ports.Collaborative {
	return &SlopeOne{
		diff:   map[string]map[string]float64{},
		weight: map[string]map[string]float64{},
	}
}

// Recommend generates recommendations for a given user based on the transaction table
func (s *SlopeOne) Recommend(table []model.Transaction, user string, minScore float64) map[string]float64 {
	userScores := s.buildUserScores(table)
	s.computeDiffs(userScores)
	return s.generateRecommendations(userScores[user], minScore)
}

// buildUserScores organizes data: user -> product -> score
func (s *SlopeOne) buildUserScores(table []model.Transaction) map[string]map[string]float64 {
	userScores := make(map[string]map[string]float64)
	for _, t := range table {
		if userScores[t.UserID] == nil {
			userScores[t.UserID] = make(map[string]float64)
		}
		userScores[t.UserID][t.ProductID] = t.Score
	}
	return userScores
}

// computeDiffs calculates average differences and weights for each item pair
func (s *SlopeOne) computeDiffs(userScores map[string]map[string]float64) {
	for _, scores := range userScores {
		items := make([]string, 0, len(scores))
		for item := range scores {
			items = append(items, item)
		}

		for i, baseItem := range items {
			baseScore := scores[baseItem]
			for _, otherItem := range items[i+1:] { // slice from i+1 to end
				otherScore := scores[otherItem]

				s.updateDiff(baseItem, otherItem, baseScore-otherScore)
				s.updateDiff(otherItem, baseItem, otherScore-baseScore)
			}
		}
	}
}

// updateDiff updates the sum of differences and the weight, initializing maps if necessary
func (s *SlopeOne) updateDiff(baseItem, otherItem string, diff float64) {
	if s.diff[baseItem] == nil {
		s.diff[baseItem] = make(map[string]float64)
		s.weight[baseItem] = make(map[string]float64)
	}
	s.diff[baseItem][otherItem] += diff
	s.weight[baseItem][otherItem]++
}

// generateRecommendations generates predictions for the target user
func (s *SlopeOne) generateRecommendations(myProducts map[string]float64, minScore float64) map[string]float64 {
	if myProducts == nil {
		return nil
	}

	predictions := make(map[string]float64)
	weightSum := make(map[string]float64)

	// calculate predictions
	for prodOther, itemDiffs := range s.diff {
		for prodRef, diffSum := range itemDiffs {
			myScore, ok := myProducts[prodRef]
			if !ok {
				continue
			}
			w := s.weight[prodOther][prodRef]
			predictions[prodOther] += (diffSum/w + myScore) * w
			weightSum[prodOther] += w
		}
	}

	// normalize and filter by minScore
	result := make(map[string]float64)
	for pid, val := range predictions {
		if _, rated := myProducts[pid]; rated {
			continue
		}
		if w := weightSum[pid]; w > 0 {
			avg := val / w
			if avg >= minScore {
				result[pid] = math.Round(avg*100) / 100
			}
		}
	}
	return result
}
