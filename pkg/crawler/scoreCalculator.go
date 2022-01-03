package crawler

import "github.com/lessbutter/alloff-api/internal/core/domain"

func GetProductScore(pd *domain.ProductDAO) *domain.ProductScoreInfoDAO {
	// TODO: Dummy data
	return &domain.ProductScoreInfoDAO{
		IsNewlyCrawled: false,
		IsUpdated:      false,
		ManualScore:    10,
		AutoScore:      10,
		TotalScore:     20,
	}
}
