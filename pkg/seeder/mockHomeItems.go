package seeder

import (
	"log"

	"github.com/lessbutter/alloff-api/api/front/model"
	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/pkg/product"
)

func AddHomeItems() {

	sandro, _ := ioc.Repo.Brands.GetByKeyname("SANDRO")
	brandsList := map[string]*domain.BrandDAO{
		"LUCKYCHOUETTE":      nil,
		"SYSTEM":             nil,
		"STUDIOTOMBOY":       nil,
		"O2ND":               nil,
		"DEWL":               nil,
		"TIME":               nil,
		"KUHO":               nil,
		"JILLSTUARTNY":       nil,
		"MINE":               nil,
		"MICHAA":             nil,
		"SHESMISS":           nil,
		"TOMMYHILFIGERW":     nil,
		"LATT":               nil,
		"HAZZYSL":            nil,
		"IZZATBABA":          nil,
		"CLAUDIEPIERLOT":     nil,
		"ISABELMARANT":       nil,
		"DKNY":               nil,
		"MOJOSPHINE":         nil,
		"VINCE":              nil,
		"SJSJ":               nil,
		"LANVINCOLLECTION":   nil,
		"OBZEE":              nil,
		"VANESSABRUNO":       nil,
		"THEIZZATCOLLECTION": nil,
		"VOV":                nil,
		"JIGOTT":             nil,
		"SANDRO":             nil,
		"LACOSTE":            nil,
		"NICECLAUP":          nil,
		"RENEEVON":           nil,
		"BENETTON":           nil,
		"SISLEY":             nil,
		"EGOIST":             nil,
		"KENNETHLADY":        nil,
		"LAP":                nil,
		"LINE":               nil,
		"LYNN":               nil,
		"MOE":                nil,
		"OLIVEDESOLIVE":      nil,
		"ONANDON":            nil,
		"PLASTICISLAND":      nil,
	}

	for brandname := range brandsList {
		brand, err := ioc.Repo.Brands.GetByKeyname(brandname)
		if err != nil {
			brandsList[brandname] = sandro
		} else {
			brandsList[brandname] = brand
		}
	}

	onepieceCategory, _ := ioc.Repo.AlloffCategories.GetByKeyname("1_ONEPIECE")
	onepieceProducts, _, _ := product.AlloffCategoryProductsListing(0, 10, nil, onepieceCategory.ID.Hex(), "", nil)
	trenchCategory, _ := ioc.Repo.AlloffCategories.GetByKeyname("2_TRENCH")
	trenchProducts, _, _ := product.AlloffCategoryProductsListing(0, 10, nil, trenchCategory.ID.Hex(), "", nil)
	coatCategory, _ := ioc.Repo.AlloffCategories.GetByKeyname("2_COAT")
	coatProducts, _, _ := product.AlloffCategoryProductsListing(0, 10, nil, coatCategory.ID.Hex(), "", nil)

	pgs, _ := ioc.Repo.ProductGroups.List()

	homeItems := []*domain.HomeItemDAO{
		{
			Priority:       100,
			Title:          "프리미엄 브랜드, 한정 수량 특가!",
			ItemType:       model.HomeItemTypeTimedeal,
			TargetID:       "",
			Sorting:        nil,
			Images:         nil,
			CommunityItems: nil,
			Brands:         nil,
			Products:       nil,
			ProductGroups:  pgs,
		},
		{
			Priority:       90,
			Title:          "신규 오픈 브랜드",
			ItemType:       model.HomeItemTypeBrand,
			TargetID:       "",
			Sorting:        nil,
			Images:         nil,
			CommunityItems: nil,
			Brands: []*domain.HomeBrandItemDAO{
				{
					ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/curated/PLASTICISLAND_curation.png",
					Brand:  brandsList["PLASTICISLAND"],
				},
				{
					ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/curated/ONANDON_curation.png",
					Brand:  brandsList["ONANDON"],
				},
				{
					ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/curated/MOE_curation.png",
					Brand:  brandsList["MOE"],
				},
				{
					ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/curated/LYNN_curation.png",
					Brand:  brandsList["LYNN"],
				},
				{
					ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/curated/LINE_curation.png",
					Brand:  brandsList["LINE"],
				},
				{
					ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/curated/LAP_curation.png",
					Brand:  brandsList["LAP"],
				},
				{
					ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/curated/KENNETHLADY_curation.png",
					Brand:  brandsList["KENNETHLADY"],
				},
				{
					ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/curated/EGOIST_curation.png",
					Brand:  brandsList["EGOIST"],
				},
			},
			Products: nil,
		},
		{
			Priority:       80,
			Title:          "매일 업데이트되는 원피스 아울렛 신상품",
			ItemType:       model.HomeItemTypeProduct,
			TargetID:       onepieceCategory.ID.Hex(),
			Sorting:        nil,
			Images:         nil,
			CommunityItems: nil,
			Brands:         nil,
			Products:       onepieceProducts,
		},
		{
			Priority:       70,
			Title:          "곧 다가올 가을, 트렌치 코트 70%~",
			ItemType:       model.HomeItemTypeProduct,
			TargetID:       trenchCategory.ID.Hex(),
			Sorting:        []model.SortingType{model.SortingTypeDiscount70_100},
			Images:         nil,
			CommunityItems: nil,
			Brands:         nil,
			Products:       trenchProducts,
		},
		{
			Priority:       65,
			Title:          "미리 준비하는 한겨울, 코트 70%~",
			ItemType:       model.HomeItemTypeProduct,
			TargetID:       coatCategory.ID.Hex(),
			Sorting:        []model.SortingType{model.SortingTypeDiscount70_100},
			Images:         nil,
			CommunityItems: nil,
			Brands:         nil,
			Products:       coatProducts,
		},
		{
			Priority:       60,
			Title:          "인기 캐주얼 브랜드 아울렛",
			ItemType:       model.HomeItemTypeBrand,
			TargetID:       "",
			Sorting:        nil,
			Images:         nil,
			CommunityItems: nil,
			Brands: []*domain.HomeBrandItemDAO{
				{
					ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/curated/curated_brand_1-1.png",
					Brand:  brandsList["LUCKYCHOUETTE"],
				},
				{
					ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/curated/curated_brand_1-3.png",
					Brand:  brandsList["STUDIOTOMBOY"],
				},
				{
					ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/curated/curated_brand_1_2nd_5.png",
					Brand:  brandsList["VOV"],
				},
			},
			Products: nil,
		},
		{
			Priority:       50,
			Title:          "인기 컨템포러리 브랜드 아울렛",
			ItemType:       model.HomeItemTypeBrand,
			TargetID:       "",
			Sorting:        nil,
			Images:         nil,
			CommunityItems: nil,
			Brands: []*domain.HomeBrandItemDAO{
				{
					ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/curated/curated_brand_2_2nd_1.png",
					Brand:  brandsList["SANDRO"],
				},
				{
					ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/curated/curated_brand_2-2.png",
					Brand:  brandsList["KUHO"],
				},
				{
					ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/curated/curated_brand_2-3.png",
					Brand:  brandsList["JILLSTUARTNY"],
				},
				{
					ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/curated/curated_brand_2_2nd_5.png",
					Brand:  brandsList["JIGOTT"],
				},
			},
			Products: nil,
		},
		{
			Priority:       20,
			Title:          "인기 커리어&트레디셔널 브랜드 아울렛",
			ItemType:       model.HomeItemTypeBrand,
			TargetID:       "",
			Sorting:        nil,
			Images:         nil,
			CommunityItems: nil,
			Brands: []*domain.HomeBrandItemDAO{
				{
					ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/curated/curated_brand_3-1.png",
					Brand:  brandsList["SHESMISS"],
				},
				{
					ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/curated/curated_brand_3-4.png",
					Brand:  brandsList["HAZZYSL"],
				},
				{
					ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/curated/curated_brand_3-5.png",
					Brand:  brandsList["IZZATBABA"],
				},
			},
			Products: nil,
		},
		{
			Priority:       10,
			Title:          "최대 70% 이상 할인 브랜드",
			ItemType:       model.HomeItemTypeBrand,
			TargetID:       "",
			Sorting:        nil,
			Images:         nil,
			CommunityItems: nil,
			Brands: []*domain.HomeBrandItemDAO{
				{
					ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/curated/curated_brand_4-1.png",
					Brand:  brandsList["CLAUDIEPIERLOT"],
				},
				{
					ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/curated/curated_brand_4-2.png",
					Brand:  brandsList["ISABELMARANT"],
				},
				{
					ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/curated/curated_brand_4-3.png",
					Brand:  brandsList["DKNY"],
				},
				{
					ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/curated/curated_brand_4-4.png",
					Brand:  brandsList["MOJOSPHINE"],
				},
				{
					ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/curated/curated_brand_4-5.png",
					Brand:  brandsList["VINCE"],
				},
			},
			Products: nil,
		},
		{
			Priority:       5,
			Title:          "최대 50% ~ 70% 할인 브랜드",
			ItemType:       model.HomeItemTypeBrand,
			TargetID:       "",
			Sorting:        nil,
			Images:         nil,
			CommunityItems: nil,
			Brands: []*domain.HomeBrandItemDAO{
				{
					ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/curated/curated_brand_5-4.png",
					Brand:  brandsList["VANESSABRUNO"],
				},
				{
					ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/curated/curated_brand_5-5.png",
					Brand:  brandsList["THEIZZATCOLLECTION"],
				},
			},
			Products: nil,
		},
	}

	for _, item := range homeItems {
		err := ioc.Repo.HomeItems.Insert(item)
		if err != nil {
			log.Println("homeitem insert err", err)
		}
	}

	log.Println("HomeItems added!")

}
