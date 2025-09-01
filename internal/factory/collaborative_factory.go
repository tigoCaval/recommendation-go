package factory

import (
	"github.com/tigoCaval/recommendation-go/internal/domain/model"
	"github.com/tigoCaval/recommendation-go/internal/domain/ports"
)

type CollaborativeFactory struct{}

func (f *CollaborativeFactory) DoFactory(
	method ports.Collaborative,
	table []model.Transaction,
	user string,
	score float64,
) map[string]float64 {
	return method.Recommend(table, user, score)
}
