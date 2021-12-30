package ioc

import (
	"github.com/lessbutter/alloff-api/internal/core/repository"
)

type iocRepo struct {
	Brands   repository.BrandsRepository
	Products repository.ProductsRepository
	Sources  repository.CrawlSourcesRepository
}

var Repo iocRepo
