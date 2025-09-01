package recommendation

import (
	"github.com/tigoCaval/recommendation-go/internal/domain/model"
	"github.com/tigoCaval/recommendation-go/internal/domain/service"
	"github.com/tigoCaval/recommendation-go/internal/factory"
)

type Transaction = model.Transaction

type Recommend struct {
	factory *factory.CollaborativeFactory
}

func NewRecommend() *Recommend {
	return &Recommend{factory: &factory.CollaborativeFactory{}}
}

func (r *Recommend) Ranking(table []Transaction, user string, score float64) map[string]float64 {
	return r.factory.DoFactory(service.NewRanking(), table, user, score)
}

func (r *Recommend) Euclidean(table []Transaction, user string, score float64) map[string]float64 {
	return r.factory.DoFactory(service.NewEuclidean(), table, user, score)
}

func (r *Recommend) SlopeOne(table []Transaction, user string, score float64) map[string]float64 {
	return r.factory.DoFactory(service.NewSlopeOne(), table, user, score)
}
