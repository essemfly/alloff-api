package malls

import (
	"log"
	"time"

	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
)

func AddOthers() {
	brands := []domain.BrandDAO{
		{
			KorName:       "비씨비지",
			EngName:       "BCBG",
			KeyName:       "BCBG",
			Description:   "커리어",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/BCBG.png",
			Onpopular:     false,
			Created:       time.Now(),
			IsOpen:        false,
			InMaintenance: false,
			SizeGuide:     []domain.SizeGuideDAO{},
		},
		{
			KorName:       "써스데이 아일랜드",
			EngName:       "Thursday Island",
			KeyName:       "THURSDAYISLAND",
			Description:   "컨템포러리 캐주얼",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/THURSDAYISLAND.png",
			Onpopular:     false,
			Created:       time.Now(),
			IsOpen:        false,
			InMaintenance: false,
			SizeGuide:     []domain.SizeGuideDAO{},
		},
		{
			KorName:       "잇미샤",
			EngName:       "itMICHAA",
			KeyName:       "ITMICHAA",
			Description:   "컨템포러리 캐주얼",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/ITMICHAA.png",
			Onpopular:     false,
			Created:       time.Now(),
			IsOpen:        false,
			InMaintenance: false,
			SizeGuide:     []domain.SizeGuideDAO{},
		},
		{
			KorName:       "버버리",
			EngName:       "BURBERRY",
			KeyName:       "BURBERRY",
			Description:   "럭셔리",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/BURBERRY.png",
			Onpopular:     false,
			Created:       time.Now(),
			IsOpen:        false,
			InMaintenance: false,
			SizeGuide:     []domain.SizeGuideDAO{},
		}, {
			KorName:       "몽클레르",
			EngName:       "Moncler",
			KeyName:       "MONCLER",
			Description:   "럭셔리",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/MONCLER.png",
			Onpopular:     false,
			Created:       time.Now(),
			IsOpen:        false,
			InMaintenance: false,
			SizeGuide:     []domain.SizeGuideDAO{},
		},
		{
			KorName:       "블루독 베이비",
			EngName:       "BLUEDOG baby",
			KeyName:       "BLUEDOGBABY",
			Description:   "키즈",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/BLUEDOGBABY.png",
			Onpopular:     false,
			Created:       time.Now(),
			IsOpen:        false,
			IsHide:        true,
			InMaintenance: false,
			SizeGuide:     []domain.SizeGuideDAO{},
		},
		{
			KorName:       "밍크뮤",
			EngName:       "minkmui",
			KeyName:       "MINKMUI",
			Description:   "키즈",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/MINKMUI.png",
			Onpopular:     false,
			Created:       time.Now(),
			IsOpen:        false,
			IsHide:        true,
			InMaintenance: false,
			SizeGuide:     []domain.SizeGuideDAO{},
		},
		{
			KorName:       "이세이미야케",
			EngName:       "ISSEY MIYAKE",
			KeyName:       "ISSEYMIYAKE",
			Description:   "컨템포러리",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/ISSEYMIYAKE.png",
			Onpopular:     false,
			Created:       time.Now(),
			IsOpen:        false,
			IsHide:        false,
			InMaintenance: false,
			SizeGuide:     []domain.SizeGuideDAO{},
		},
		{
			KorName:       "플리츠플리즈 이세이미야케",
			EngName:       "PLEATS PLEASE ISSEY MIYAKE",
			KeyName:       "PLEATSPLEASE",
			Description:   "클래식",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/PLEATSPLEASE.png",
			Onpopular:     false,
			Created:       time.Now(),
			IsOpen:        false,
			IsHide:        false,
			InMaintenance: false,
			SizeGuide:     []domain.SizeGuideDAO{},
		},
		{
			KorName:       "바오바오 이세이미야케",
			EngName:       "BAOBAO ISSEY MIYAKE",
			KeyName:       "BAOBAO",
			Description:   "컨템포러리 디자이너",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/BAOBAO.png",
			Onpopular:     false,
			Created:       time.Now(),
			IsOpen:        false,
			IsHide:        false,
			InMaintenance: false,
			SizeGuide:     []domain.SizeGuideDAO{},
		},
		{
			KorName:       "제이린드버그",
			EngName:       "J.LINDEBERG",
			KeyName:       "JLINDEBERG",
			Description:   "스포츠/아웃도어",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/JLINDEBERG.png",
			Onpopular:     false,
			Created:       time.Now(),
			IsOpen:        false,
			IsHide:        false,
			InMaintenance: false,
			SizeGuide:     []domain.SizeGuideDAO{},
		},
		{
			KorName:       "보테가베네타",
			EngName:       "BOTTEGA VENETA",
			KeyName:       "BOTTEGAVENETA",
			Description:   "럭셔리",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/BOTTEGAVENETA.png",
			Onpopular:     false,
			Created:       time.Now(),
			IsOpen:        false,
			IsHide:        false,
			InMaintenance: false,
			SizeGuide:     []domain.SizeGuideDAO{},
		},
		{
			KorName:       "프라다",
			EngName:       "PRADA",
			KeyName:       "PRADA",
			Description:   "럭셔리",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/PRADA.png",
			Onpopular:     false,
			Created:       time.Now(),
			IsOpen:        false,
			IsHide:        false,
			InMaintenance: false,
			SizeGuide:     []domain.SizeGuideDAO{},
		},
		{
			KorName:       "토리버치",
			EngName:       "TORYBURCH",
			KeyName:       "TORYBURCH",
			Description:   "컨템포러리 디자이너",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/TORYBURCH.png",
			Onpopular:     false,
			Created:       time.Now(),
			IsOpen:        false,
			IsHide:        false,
			InMaintenance: false,
			SizeGuide:     []domain.SizeGuideDAO{},
		},
		{
			KorName:       "로에베",
			EngName:       "LOEWE",
			KeyName:       "LOEWE",
			Description:   "컨템포러리 디자이너",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/LOEWE.png",
			Onpopular:     false,
			Created:       time.Now(),
			IsOpen:        false,
			IsHide:        false,
			InMaintenance: false,
			SizeGuide:     []domain.SizeGuideDAO{},
		},
		{
			KorName:       "가니",
			EngName:       "GANNI",
			KeyName:       "GANNI",
			Description:   "컨템포러리 디자이너",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/GANNI.png",
			Onpopular:     false,
			Created:       time.Now(),
			IsOpen:        false,
			IsHide:        false,
			InMaintenance: false,
			SizeGuide:     []domain.SizeGuideDAO{},
		},
	}

	for _, brand := range brands {
		_, err := ioc.Repo.Brands.Upsert(&brand)
		if err != nil {
			log.Println(err)
		}
	}

	log.Println("All not open brands are added")
}
