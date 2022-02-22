package scripts

import (
	"log"
	"strconv"
	"time"

	"github.com/lessbutter/alloff-api/api/apiServer/model"
	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/pkg/exhibition"
	"github.com/lessbutter/alloff-api/pkg/hometab"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddMockExhibitions() {
	log.Println("ADD MOCK EXHIBITIONS START ********")
	pgs, _ := ioc.Repo.ProductGroups.List(5)
	pgIDs := []string{}
	for idx, pg := range pgs {
		log.Println("IDX", idx)
		pgIDs = append(pgIDs, pg.ID.Hex())
		req := getMockExhibition(idx+1, pgIDs)
		exhibitionDao, err := exhibition.AddExhibition(req)
		if err != nil {
			log.Println("err on aadd mock exhibitions", err)
		}
		_, err = ioc.Repo.Exhibitions.Upsert(exhibitionDao)
		if err != nil {
			log.Println("upsert failed on exhibitions")
		}
	}
}

func getMockExhibition(idx int, pgIDs []string) *exhibition.ExhibitionRequest {
	bannerIdx := strconv.Itoa(idx)
	thumbnailIdx := strconv.Itoa(idx * 100)

	return &exhibition.ExhibitionRequest{
		BannerImage:     "https://picsum.photos/seed/" + bannerIdx + "/200/300",
		ThumbnailImage:  "https://picsum.photos/seed/" + thumbnailIdx + "/200/300",
		Title:           "Exhibition" + bannerIdx,
		Description:     thumbnailIdx + "이별보다 더 아픈게 그리움인데, 무시로 무시로, 외로울 때 그때 울어요",
		ProductGroupIDs: pgIDs,
		StartTime:       time.Now(),
		FinishTime:      time.Now().Add(1200 * time.Hour),
	}
}

func AddMockHomeTabs() {
	log.Println("ADD MOCK HOME TABS START ********")
	exhibitions, _, _ := ioc.Repo.Exhibitions.List(0, 10, true)
	exhitibionIDs := []string{}
	for _, ex := range exhibitions {
		exhitibionIDs = append(exhitibionIDs, ex.ID.Hex())
	}
	alloffCategories, _ := ioc.Repo.AlloffCategories.List(nil)
	alloffcatIDs := []string{}
	for _, cat := range alloffCategories {
		alloffcatIDs = append(alloffcatIDs, cat.ID.Hex())
	}

	hometabItems := []*hometab.HomeTabItemRequest{
		{
			Title:        "BRAND SALE",
			Description:  "브랜드가 세일이 몇프로씩 되었는지 사용되는 기획전입니다.",
			Tags:         []string{},
			BackImageUrl: "",
			StartedAt:    time.Now(),
			FinishedAt:   time.Now().Add(1200 * time.Hour),
			Requester: &hometab.BrandsItemRequest{
				BrandKeynames: []string{
					"MAJE", "GAP", "NICECLAUP", "CHLOE", "SHESMISS",
				},
			},
		},
		{
			Title:        "기획전 A",
			Description:  "좀 못생긴 외국 남자가 가방을 들고 서있는 기획전 A페이지 입니다. 아 물론 피그마상에서요.",
			Tags:         []string{},
			BackImageUrl: "https://pds.joongang.co.kr/news/component/htmlphoto_mmdata/202005/25/8ab5037f-8ac6-4597-9197-09b328f2c514.jpg",
			StartedAt:    time.Now(),
			FinishedAt:   time.Now().Add(1200 * time.Hour),
			Requester: &hometab.BrandExhibitionItemRequest{
				ExhibitionID: exhitibionIDs[0],
			},
		},
		{
			Title:        "기획전 모음",
			Description:  "",
			Tags:         []string{},
			BackImageUrl: "",
			StartedAt:    time.Now(),
			FinishedAt:   time.Now().Add(1200 * time.Hour),
			Requester: &hometab.ExhibitionsItemRequest{
				ExhibitionIDs: exhitibionIDs,
			},
		},
		{
			Title:        "기획전 C 짧으면",
			Description:  "",
			Tags:         []string{"여러분", "모두", "부자되세요"},
			BackImageUrl: "https://img.huffingtonpost.com/asset/61494acf240000140118d34e.png",
			StartedAt:    time.Now(),
			FinishedAt:   time.Now().Add(1200 * time.Hour),
			Requester: &hometab.ExhibitionItemRequest{
				ExhibitionID: exhitibionIDs[0],
			},
		},
		{
			Title:        "기획전 C 타이틀 길면",
			Description:  "",
			Tags:         []string{"묻지도 않고", "따지지도않고", "라이나생명보험"},
			BackImageUrl: "https://i.ytimg.com/vi/zK58Ht3OQhM/maxresdefault.jpg",
			StartedAt:    time.Now(),
			FinishedAt:   time.Now().Add(1200 * time.Hour),
			Requester: &hometab.ExhibitionItemRequest{
				ExhibitionID: exhitibionIDs[1],
			},
		},
		{
			Title:        "큐레이션 A",
			Description:  "",
			Tags:         []string{},
			BackImageUrl: "",
			StartedAt:    time.Now(),
			FinishedAt:   time.Now().Add(1200 * time.Hour),
			Requester: &hometab.AlloffCategoryItemRequest{
				AlloffCategoryID: alloffcatIDs[0],
				SortingOptions: []model.SortingType{
					model.SortingTypeDiscount70_100,
				},
			},
		},
		{
			Title:        "큐레이션 B",
			Description:  "",
			Tags:         []string{},
			BackImageUrl: "",
			StartedAt:    time.Now(),
			FinishedAt:   time.Now().Add(1200 * time.Hour),
			Requester: &hometab.AlloffCategoryItemRequest{
				AlloffCategoryID: alloffcatIDs[1],
				SortingOptions: []model.SortingType{
					model.SortingTypeDiscount50_70,
					model.SortingTypeDiscount70_100,
				},
			},
		},
	}

	for idx, item := range hometabItems {
		log.Println("IDX", idx)
		_, err := hometab.AddHometabItem(item)
		if err != nil {
			log.Println(err)
		}
	}
}

func AddMockTopBanners() {
	log.Println("ADD MOCK TOP BANNERS START ********")
	exhibitions, _, _ := ioc.Repo.Exhibitions.List(0, 5, true)
	for id, ex := range exhibitions {
		log.Println("IDX", id)
		idString := strconv.Itoa(id * 77)
		newBanner := &domain.TopBannerDAO{
			ID:           primitive.NewObjectID(),
			ImageUrl:     "https://picsum.photos/seed/" + idString + "/200/300",
			ExhibitionID: ex.ID.Hex(),
			Title:        "위기를 넘어 기회로 " + idString,
			SubTitle:     idString + "번째 탑배너",
			IsLive:       true,
			Weight:       id,
		}

		_, err := ioc.Repo.TopBanners.Insert(newBanner)
		if err != nil {
			log.Println("Err in add top banners")
		}
	}
}

func AddHomeTab() {
	hometabItems := []*hometab.HomeTabItemRequest{
		{
			Title:        "봄까지 입기 좋은 가디건, 아울렛 특가",
			Description:  "",
			Tags:         []string{},
			BackImageUrl: "",
			StartedAt:    time.Now(),
			FinishedAt:   time.Now().Add(240 * time.Hour),
			Requester: &hometab.AlloffCategoryItemRequest{
				AlloffCategoryID: "621008161c35be7a91ccbfc8",
				SortingOptions: []model.SortingType{
					model.SortingTypeDiscount70_100,
					model.SortingTypeDiscount50_70,
				},
			},
		},
		{
			Title:        "올오프 프리미엄 아울렛, 자켓 특가",
			Description:  "",
			Tags:         []string{},
			BackImageUrl: "",
			StartedAt:    time.Now(),
			FinishedAt:   time.Now().Add(1200 * time.Hour),
			Requester: &hometab.AlloffCategoryItemRequest{
				AlloffCategoryID: "621008161c35be7a91ccbfc2",
				SortingOptions: []model.SortingType{
					model.SortingTypeDiscount50_70,
					model.SortingTypeDiscount70_100,
				},
			},
		},
	}

	for idx, item := range hometabItems {
		log.Println("IDX", idx)
		_, err := hometab.AddHometabItem(item)
		if err != nil {
			log.Println(err)
		}
	}
}
