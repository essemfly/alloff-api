package hometab

import (
	"log"
	"time"

	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/pkg/product"
)

type HomeTabItemRequest struct {
	Title        string
	Description  string
	Tags         []string
	BackImageUrl string
	StartedAt    time.Time
	EndedAt      time.Time
	Requester    ItemRequest
}

type ItemRequest interface {
	fillItemContents(*domain.HomeTabItemDAO) *domain.HomeTabItemDAO
}

type BrandsItemRequest struct {
	BrandKeynames []string
}

// 브랜드들만 있는 것을 보여주는 것
func (req *BrandsItemRequest) fillItemContents(item *domain.HomeTabItemDAO) *domain.HomeTabItemDAO {
	brandDaos := []*domain.BrandDAO{}
	for _, keyname := range req.BrandKeynames {
		brandDao, err := ioc.Repo.Brands.GetByKeyname(keyname)
		if err != nil {
			log.Println("brand not found: "+keyname, err)
			continue
		}
		brandDaos = append(brandDaos, brandDao)
	}

	item.Type = domain.HOMETAB_ITEM_BRANDS
	item.Brands = brandDaos
	item.Reference = &domain.ReferenceTarget{
		Path:   "brands",
		Params: "",
	}

	return item
}

type BrandExhibitionItemRequest struct {
	BrandKeyname string
	ExhibitionID string
}

// 하나의 브랜드와 3,4개의 상품들이 포함된 기획전 보여주는 것
func (req *BrandExhibitionItemRequest) fillItemContents(item *domain.HomeTabItemDAO) *domain.HomeTabItemDAO {
	exhibitionDao, err := ioc.Repo.Exhibitions.Get(req.ExhibitionID)
	if err != nil {
		log.Println("err in brand exhibition item req", err)
	}

	brandDao, err := ioc.Repo.Brands.GetByKeyname(req.BrandKeyname)
	if err != nil {
		log.Println("err in brand exhibition item req", err)
	}

	item.Type = domain.HOMETAB_ITEM_BRAND_EXHIBITION
	item.Brands = []*domain.BrandDAO{
		brandDao,
	}
	item.Exhibitions = []*domain.ExhibitionDAO{
		exhibitionDao,
	}
	item.Products = exhibitionDao.ListCheifProducts()
	item.Reference = &domain.ReferenceTarget{
		Path:   "exhibition",
		Params: exhibitionDao.ID.Hex(),
	}

	return item
}

// 기획전들 여러개 보여주는 것
type ExhibitionsItemRequest struct {
	ExhibitionIDs []string
}

func (req *ExhibitionsItemRequest) fillItemContents(item *domain.HomeTabItemDAO) *domain.HomeTabItemDAO {
	exhibitionDaos := []*domain.ExhibitionDAO{}
	for _, exhibitionID := range req.ExhibitionIDs {
		exhibitionDao, err := ioc.Repo.Exhibitions.Get(exhibitionID)
		if err != nil {
			log.Println("exhibition id not found: "+exhibitionID, err)
			continue
		}
		exhibitionDaos = append(exhibitionDaos, exhibitionDao)
	}
	item.Type = domain.HOMETAB_ITEM_EXHIBITIONS
	item.Exhibitions = exhibitionDaos
	item.Reference = &domain.ReferenceTarget{
		Path:   "exhibitions",
		Params: "",
	}

	return item
}

// 기획전인데 기획전에 속한 상품 몇개 보여주는 것
type ExhibitionItemRequest struct {
	ExhibitionID string
}

func (req *ExhibitionItemRequest) fillItemContents(item *domain.HomeTabItemDAO) *domain.HomeTabItemDAO {
	exhibitionDao, err := ioc.Repo.Exhibitions.Get(req.ExhibitionID)
	if err != nil {
		log.Println("err in brand exhibition item req", err)
	}

	item.Type = domain.HOMETAB_ITEM_BRAND_EXHIBITION
	item.Exhibitions = []*domain.ExhibitionDAO{
		exhibitionDao,
	}

	item.Products = exhibitionDao.ListCheifProducts()
	item.Reference = &domain.ReferenceTarget{
		Path:   "exhibition",
		Params: exhibitionDao.ID.Hex(),
	}

	return item
}

// 기존 Curation: AlloffCategory와 Options로 Sorting된 것 보여주는 것
type AlloffCategoryItemRequest struct {
	AlloffCategoryID string
	SortingOptions   []string
}

func (req *AlloffCategoryItemRequest) fillItemContents(item *domain.HomeTabItemDAO) *domain.HomeTabItemDAO {
	numProductsToShow := 10
	products, _, err := product.AlloffCategoryProductsListing(0, numProductsToShow, nil, req.AlloffCategoryID, "", req.SortingOptions)
	if err != nil {
		log.Println("alloffcat id not found: " + req.AlloffCategoryID)
	}

	item.Type = domain.HOMETAB_ITEM_PRODUCTS
	item.Products = products
	item.Reference = &domain.ReferenceTarget{
		Path:    "products",
		Params:  req.AlloffCategoryID,
		Options: req.SortingOptions,
	}

	return item
}
