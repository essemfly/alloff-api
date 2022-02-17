package seeder

import (
	"log"

	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddNewAlloffCats() {
	outerOID := primitive.NewObjectID()
	topOID := primitive.NewObjectID()
	bottomOID := primitive.NewObjectID()
	onePieceOID := primitive.NewObjectID()
	skirtOID := primitive.NewObjectID()
	underwearOID := primitive.NewObjectID()
	bagsOID := primitive.NewObjectID()
	shoesOID := primitive.NewObjectID()
	accessoryOID := primitive.NewObjectID()

	cats := []*domain.AlloffCategoryDAO{
		{
			ID:           outerOID,
			KeyName:      "1_outer",
			Name:         "아우터",
			Level:        1,
			ParentId:     primitive.NilObjectID,
			CategoryType: "NORMAL",
			ImgURL:       "https://alloff.s3.ap-northeast-2.amazonaws.com/alloffcategories/1_outer.svg",
		},
		{
			ID:           topOID,
			KeyName:      "1_top",
			Name:         "상의",
			Level:        1,
			ParentId:     primitive.NilObjectID,
			CategoryType: "NORMAL",
			ImgURL:       "https://alloff.s3.ap-northeast-2.amazonaws.com/alloffcategories/1_top.svg",
		},
		{
			ID:           bottomOID,
			KeyName:      "1_bottom",
			Name:         "바지",
			Level:        1,
			ParentId:     primitive.NilObjectID,
			CategoryType: "NORMAL",
			ImgURL:       "https://alloff.s3.ap-northeast-2.amazonaws.com/alloffcategories/1_bottom.svg",
		},
		{
			ID:           onePieceOID,
			KeyName:      "1_onePiece",
			Name:         "원피스/세트",
			Level:        1,
			ParentId:     primitive.NilObjectID,
			CategoryType: "NORMAL",
			ImgURL:       "https://alloff.s3.ap-northeast-2.amazonaws.com/alloffcategories/1_onePiece.svg",
		},
		{
			ID:           skirtOID,
			KeyName:      "1_skirt",
			Name:         "스커트",
			Level:        1,
			ParentId:     primitive.NilObjectID,
			CategoryType: "NORMAL",
			ImgURL:       "https://alloff.s3.ap-northeast-2.amazonaws.com/alloffcategories/1_skirt.svg",
		},
		{
			ID:           underwearOID,
			KeyName:      "1_underwear",
			Name:         "라운지/언더웨어",
			Level:        1,
			ParentId:     primitive.NilObjectID,
			CategoryType: "NORMAL",
			ImgURL:       "https://alloff.s3.ap-northeast-2.amazonaws.com/alloffcategories/1_underwear.svg",
		},
		{
			ID:           bagsOID,
			KeyName:      "1_bags",
			Name:         "가방",
			Level:        1,
			ParentId:     primitive.NilObjectID,
			CategoryType: "NORMAL",
			ImgURL:       "https://alloff.s3.ap-northeast-2.amazonaws.com/alloffcategories/1_bags.svg",
		},
		{
			ID:           shoesOID,
			KeyName:      "1_shoes",
			Name:         "신발",
			Level:        1,
			ParentId:     primitive.NilObjectID,
			CategoryType: "NORMAL",
			ImgURL:       "https://alloff.s3.ap-northeast-2.amazonaws.com/alloffcategories/1_shoes.svg",
		},
		{
			ID:           accessoryOID,
			KeyName:      "1_accessory",
			Name:         "패션잡화",
			Level:        1,
			ParentId:     primitive.NilObjectID,
			CategoryType: "NORMAL",
			ImgURL:       "https://alloff.s3.ap-northeast-2.amazonaws.com/alloffcategories/1_accessory.svg",
		},
	}

	outer2ndCats := []*domain.AlloffCategoryDAO{
		{
			ID:           primitive.NewObjectID(),
			KeyName:      "2_coat",
			Name:         "코트",
			Level:        2,
			ParentId:     outerOID,
			CategoryType: "NORMAL",
			ImgURL:       "",
		},
		{
			ID:           primitive.NewObjectID(),
			KeyName:      "2_jumper",
			Name:         "점퍼",
			Level:        2,
			ParentId:     outerOID,
			CategoryType: "NORMAL",
			ImgURL:       "",
		},
		{
			ID:           primitive.NewObjectID(),
			KeyName:      "2_jacket",
			Name:         "자켓",
			Level:        2,
			ParentId:     outerOID,
			CategoryType: "NORMAL",
			ImgURL:       "",
		},
		{
			ID:           primitive.NewObjectID(),
			KeyName:      "2_vest",
			Name:         "베스트",
			Level:        2,
			ParentId:     outerOID,
			CategoryType: "NORMAL",
			ImgURL:       "",
		},
		{
			ID:           primitive.NewObjectID(),
			KeyName:      "2_field",
			Name:         "아우터",
			Level:        2,
			ParentId:     outerOID,
			CategoryType: "NORMAL",
			ImgURL:       "",
		},
		{
			ID:           primitive.NewObjectID(),
			KeyName:      "2_padding",
			Name:         "패딩",
			Level:        2,
			ParentId:     outerOID,
			CategoryType: "NORMAL",
			ImgURL:       "",
		},
	}

	cats = append(cats, outer2ndCats...)

	top2ndCats := []*domain.AlloffCategoryDAO{
		{
			ID:           primitive.NewObjectID(),
			KeyName:      "2_tshirt",
			Name:         "티셔츠",
			Level:        2,
			ParentId:     topOID,
			CategoryType: "NORMAL",
			ImgURL:       "",
		},
		{
			ID:           primitive.NewObjectID(),
			KeyName:      "2_knit",
			Name:         "니트/스웨터",
			Level:        2,
			ParentId:     topOID,
			CategoryType: "NORMAL",
			ImgURL:       "",
		},
		{
			ID:           primitive.NewObjectID(),
			KeyName:      "2_cardigan",
			Name:         "가디건",
			Level:        2,
			ParentId:     topOID,
			CategoryType: "NORMAL",
			ImgURL:       "",
		},
		{
			ID:           primitive.NewObjectID(),
			KeyName:      "2_blouse",
			Name:         "블라우스",
			Level:        2,
			ParentId:     topOID,
			CategoryType: "NORMAL",
			ImgURL:       "",
		},
		{
			ID:           primitive.NewObjectID(),
			KeyName:      "2_shirt",
			Name:         "셔츠",
			Level:        2,
			ParentId:     topOID,
			CategoryType: "NORMAL",
			ImgURL:       "",
		},
		{
			ID:           primitive.NewObjectID(),
			KeyName:      "2_mantoman",
			Name:         "맨투맨",
			Level:        2,
			ParentId:     topOID,
			CategoryType: "NORMAL",
			ImgURL:       "",
		},
		{
			ID:           primitive.NewObjectID(),
			KeyName:      "2_hood",
			Name:         "후드",
			Level:        2,
			ParentId:     topOID,
			CategoryType: "NORMAL",
			ImgURL:       "",
		},
		{
			ID:           primitive.NewObjectID(),
			KeyName:      "2_sleeveless",
			Name:         "민소매",
			Level:        2,
			ParentId:     topOID,
			CategoryType: "NORMAL",
			ImgURL:       "",
		},
	}

	cats = append(cats, top2ndCats...)

	bottom2ndCats := []*domain.AlloffCategoryDAO{
		{
			ID:           primitive.NewObjectID(),
			KeyName:      "2_slacks",
			Name:         "슬랙스",
			Level:        2,
			ParentId:     bottomOID,
			CategoryType: "NORMAL",
			ImgURL:       "",
		},
		{
			ID:           primitive.NewObjectID(),
			KeyName:      "2_denim",
			Name:         "데님",
			Level:        2,
			ParentId:     bottomOID,
			CategoryType: "NORMAL",
			ImgURL:       "",
		},
		{
			ID:           primitive.NewObjectID(),
			KeyName:      "2_longpants",
			Name:         "팬츠",
			Level:        2,
			ParentId:     bottomOID,
			CategoryType: "NORMAL",
			ImgURL:       "",
		},
		// {
		// 	ID:           bottomOID,
		// 	KeyName:      "2_shortpants",
		// 	Name:         "숏팬츠",
		// 	Level:        2,
		// 	ParentId:     bottomOID,
		// 	CategoryType: "NORMAL",
		// 	ImgURL:       "",
		// },
		// {
		// 	ID:           bottomOID,
		// 	KeyName:      "2_leggings",
		// 	Name:         "레깅스",
		// 	Level:        2,
		// 	ParentId:     bottomOID,
		// 	CategoryType: "NORMAL",
		// 	ImgURL:       "",
		// },
	}

	cats = append(cats, bottom2ndCats...)

	onepiece2ndCats := []*domain.AlloffCategoryDAO{
		// {
		// 	ID:           primitive.NewObjectID(),
		// 	KeyName:      "2_mini",
		// 	Name:         "미니원피스",
		// 	Level:        2,
		// 	ParentId:     onePieceOID,
		// 	CategoryType: "NORMAL",
		// 	ImgURL:       "",
		// },
		// {
		// 	ID:           primitive.NewObjectID(),
		// 	KeyName:      "2_midi",
		// 	Name:         "미디원피스",
		// 	Level:        2,
		// 	ParentId:     onePieceOID,
		// 	CategoryType: "NORMAL",
		// 	ImgURL:       "",
		// },
		// {
		// 	ID:           primitive.NewObjectID(),
		// 	KeyName:      "2_maxi",
		// 	Name:         "맥시원피스",
		// 	Level:        2,
		// 	ParentId:     onePieceOID,
		// 	CategoryType: "NORMAL",
		// 	ImgURL:       "",
		// },
		// {
		// 	ID:           primitive.NewObjectID(),
		// 	KeyName:      "2_jumpsuit",
		// 	Name:         "점프수트",
		// 	Level:        2,
		// 	ParentId:     onePieceOID,
		// 	CategoryType: "NORMAL",
		// 	ImgURL:       "",
		// },
		// {
		// 	ID:           primitive.NewObjectID(),
		// 	KeyName:      "2_set",
		// 	Name:         "투피스/세트",
		// 	Level:        2,
		// 	ParentId:     onePieceOID,
		// 	CategoryType: "NORMAL",
		// 	ImgURL:       "",
		// },
	}
	cats = append(cats, onepiece2ndCats...)

	skirt2ndCats := []*domain.AlloffCategoryDAO{
		// {
		// 	ID:           primitive.NewObjectID(),
		// 	KeyName:      "2_miniskirt",
		// 	Name:         "미니스커트",
		// 	Level:        2,
		// 	ParentId:     skirtOID,
		// 	CategoryType: "NORMAL",
		// 	ImgURL:       "",
		// },
		// {
		// 	ID:           primitive.NewObjectID(),
		// 	KeyName:      "2_midiskirt",
		// 	Name:         "미디스커트",
		// 	Level:        2,
		// 	ParentId:     skirtOID,
		// 	CategoryType: "NORMAL",
		// 	ImgURL:       "",
		// },
		// {
		// 	ID:           primitive.NewObjectID(),
		// 	KeyName:      " 2_longskirt",
		// 	Name:         "롱스커트",
		// 	Level:        2,
		// 	ParentId:     skirtOID,
		// 	CategoryType: "NORMAL",
		// 	ImgURL:       "",
		// },
	}
	cats = append(cats, skirt2ndCats...)

	underwear2ndCats := []*domain.AlloffCategoryDAO{
		// {
		// 	ID:           primitive.NewObjectID(),
		// 	KeyName:      "2_rounge",
		// 	Name:         "라운지웨어",
		// 	Level:        2,
		// 	ParentId:     underwearOID,
		// 	CategoryType: "NORMAL",
		// 	ImgURL:       "",
		// },
		// {
		// 	ID:           primitive.NewObjectID(),
		// 	KeyName:      "2_under",
		// 	Name:         "언더웨어",
		// 	Level:        2,
		// 	ParentId:     underwearOID,
		// 	CategoryType: "NORMAL",
		// 	ImgURL:       "",
		// },
	}
	cats = append(cats, underwear2ndCats...)

	accessory2ndCats := []*domain.AlloffCategoryDAO{
		// {
		// 	ID:           primitive.NewObjectID(),
		// 	KeyName:      "2_belt",
		// 	Name:         "밸트",
		// 	Level:        2,
		// 	ParentId:     accessoryOID,
		// 	CategoryType: "NORMAL",
		// 	ImgURL:       "",
		// },
		// {
		// 	ID:           primitive.NewObjectID(),
		// 	KeyName:      "2_gloves",
		// 	Name:         "장갑",
		// 	Level:        2,
		// 	ParentId:     accessoryOID,
		// 	CategoryType: "NORMAL",
		// 	ImgURL:       "",
		// },
		// {
		// 	ID:           primitive.NewObjectID(),
		// 	KeyName:      "2_hats",
		// 	Name:         "모자",
		// 	Level:        2,
		// 	ParentId:     accessoryOID,
		// 	CategoryType: "NORMAL",
		// 	ImgURL:       "",
		// },
		// {
		// 	ID:           primitive.NewObjectID(),
		// 	KeyName:      "2_scarf",
		// 	Name:         "스카프/머플러",
		// 	Level:        2,
		// 	ParentId:     accessoryOID,
		// 	CategoryType: "NORMAL",
		// 	ImgURL:       "",
		// },
		// {
		// 	ID:           primitive.NewObjectID(),
		// 	KeyName:      "2_socks",
		// 	Name:         "양말",
		// 	Level:        2,
		// 	ParentId:     accessoryOID,
		// 	CategoryType: "NORMAL",
		// 	ImgURL:       "",
		// },
		// {
		// 	ID:           primitive.NewObjectID(),
		// 	KeyName:      "2_wallet",
		// 	Name:         "지갑",
		// 	Level:        2,
		// 	ParentId:     accessoryOID,
		// 	CategoryType: "NORMAL",
		// 	ImgURL:       "",
		// },
		// {
		// 	ID:           primitive.NewObjectID(),
		// 	KeyName:      "2_etc",
		// 	Name:         "기타",
		// 	Level:        2,
		// 	ParentId:     accessoryOID,
		// 	CategoryType: "NORMAL",
		// 	ImgURL:       "",
		// },
	}
	cats = append(cats, accessory2ndCats...)

	for _, cat := range cats {
		_, err := ioc.Repo.AlloffCategories.Upsert(cat)
		if err != nil {
			log.Println("err in adding alloff cats", err)
		}
	}
}
