package ports

import "github.com/tigoCaval/recommendation-go/internal/domain/model"

type Collaborative interface {
	Recommend(table []model.Transaction, user string, score float64) map[string]float64
}
