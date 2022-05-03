package mapper

import (
	"sort"
	"time"

	"github.com/lessbutter/alloff-api/api/apiServer/model"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/pkg/exhibition"
)

func MapProductGroupDao(pgDao *domain.ProductGroupDAO) *model.ProductGroup {
	var pds []*model.Product
	var soldouts []*model.Product

	sort.Slice(pgDao.Products, func(i, j int) bool {
		return pgDao.Products[i].Priority < pgDao.Products[j].Priority
	})

	for _, productInPg := range pgDao.Products {
		if productInPg.Product.Removed {
			continue
		}

		pd := MapProductDaoToProduct(productInPg.Product)
		isSoldOut := true
		for _, inv := range productInPg.Product.Inventory {
			if inv.Quantity > 0 {
				isSoldOut = false
			}
		}
		if isSoldOut || pd.Soldout {
			soldouts = append(soldouts, pd)
		} else {
			pds = append(pds, pd)
		}
	}

	pds = append(pds, soldouts...)

	brand := &model.Brand{}
	if pgDao.Brand != nil {
		brand = MapBrandDaoToBrand(pgDao.Brand, false)
	}

	pg := &model.ProductGroup{
		ID:            pgDao.ID.Hex(),
		Title:         pgDao.Title,
		ShortTitle:    pgDao.ShortTitle,
		Instruction:   pgDao.Instruction,
		ImgURL:        pgDao.ImgUrl,
		NumAlarms:     pgDao.NumAlarms,
		Products:      pds,
		StartTime:     pgDao.StartTime.Add(9 * time.Hour).String(),
		FinishTime:    pgDao.FinishTime.Add(9 * time.Hour).String(),
		SetAlarm:      false,
		Brand:         brand,
		TotalProducts: len(pgDao.Products),
	}
	return pg
}

func MapExhibition(exDao *domain.ExhibitionDAO, brief bool) *model.Exhibition {
	pgs := []*model.ProductGroup{}

	if !brief {
		for _, pg := range exDao.ProductGroups {
			pgs = append(pgs, MapProductGroupDao(pg))
		}
	}
	sales := 0
	if exDao.ExhibitionType == domain.EXHIBITION_GROUPDEAL {
		sales = exhibition.GetCurrentSales(exDao)
	}

	numPgs := len(exDao.ProductGroups)
	numProducts := 0
	for _, pg := range exDao.ProductGroups {
		numProducts += len(pg.Products)
	}

	return &model.Exhibition{
		ID:                 exDao.ID.Hex(),
		BannerImage:        exDao.BannerImage,
		ThumbnailImage:     exDao.ThumbnailImage,
		Title:              exDao.Title,
		SubTitle:           exDao.SubTitle,
		Description:        exDao.Description,
		ProductGroups:      pgs,
		StartTime:          exDao.StartTime.Add(9 * time.Hour).String(),
		FinishTime:         exDao.FinishTime.Add(9 * time.Hour).String(),
		ExhibitionType:     MapExhibitionType(exDao.ExhibitionType),
		TargetSales:        exDao.TargetSales,
		CurrentSales:       sales,
		Banners:            mapBanners(exDao.Banners),
		TotalProducts:      numProducts,
		TotalProductGroups: numPgs,
		NumUsersRequired:   exDao.NumUsersRequired,
		TotalParticipants:  exDao.TotalParticipants,
		TotalUserGroups:    exDao.TotalGroups,
		UserGroup:          &model.UserGroup{},
	}
}

// MapExhibitionType : TODO 이거 ExhibitionDAO 에 메서드로 넣고싶다..
func MapExhibitionType(enum domain.ExhibitionType) model.ExhibitionType {
	switch enum {
	case domain.EXHIBITION_GROUPDEAL:
		return model.ExhibitionTypeGroupdeal
	case domain.EXHIBITION_NORMAL:
		return model.ExhibitionTypeNormal
	case domain.EXHIBITION_TIMEDEAL:
		return model.ExhibitionTypeTimedeal
	}
	return model.ExhibitionTypeNormal
}

// mapBanners : TODO 이거 ExhibitionDAO 에 메서드로 넣고싶다..
func mapBanners(bannersDaos []domain.ExhibitionBanner) []*model.ExhibitionBanner {
	var res []*model.ExhibitionBanner
	for _, banner := range bannersDaos {
		bannerDto := model.ExhibitionBanner{
			ImgURL:         banner.ImgUrl,
			ProductGroupID: banner.ProductGroupId,
			Title:          banner.Title,
			Subtitle:       banner.Subtitle,
		}
		res = append(res, &bannerDto)
	}
	return res
}
