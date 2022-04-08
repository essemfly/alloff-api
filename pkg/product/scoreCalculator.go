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
	totalScore := rand.Intn(150) // base 100 + average discount rate 50

	newManuelScore := totalScore
	if pd.Score != nil {
		newManuelScore = pd.Score.TotalScore
	}

	return &domain.ProductScoreInfoDAO{
		IsNewlyCrawled: isNew,
		ManualScore:    0, // To be developed
		AutoScore:      0, // To be developed
		TotalScore:     newManuelScore,
	}
}
