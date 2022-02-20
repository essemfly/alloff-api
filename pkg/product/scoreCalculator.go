package product

import (
	"math/rand"
	"time"

	"github.com/lessbutter/alloff-api/internal/core/domain"
)

func GetProductScore(pd *domain.ProductDAO) *domain.ProductScoreInfoDAO {
	newlyCrawledCriterion := time.Now().Add(-45 * time.Hour)

	isNew := true
	if pd.Created.Before(newlyCrawledCriterion) {
		isNew = false
	}
	totalScore := rand.Intn(100)
	return &domain.ProductScoreInfoDAO{
		IsNewlyCrawled: isNew,
		ManualScore:    10,
		AutoScore:      10,
		TotalScore:     totalScore,
	}
}
