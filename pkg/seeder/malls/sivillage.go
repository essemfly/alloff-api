package malls

import (
	"log"
	"time"

	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
)

func AddSiVillages() {
	modulename := "sivillage"
	categories := map[string]string{
		"코트":       "010000029408",
		"점퍼":       "010000029415",
		"자켓/베스트":   "010000029421",
		"수트":       "010000029427",
		"원피스/점프수트": "010000029453",
		"셔츠/블라우스":  "010000029439",
		"니트":       "010000029446",
		"티셔츠":      "010000029431",
		"팬츠":       "010000029460",
		"스커트":      "010000029476",
		"라운지/언더웨어": "010000029481",
		"골프/스포츠웨어": "010000029488",
	}

	brands := map[string]domain.BrandDAO{
		"72%3AG-CUT": {
			KorName:       "지컷",
			EngName:       "g-cut",
			KeyName:       "GCUT",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/GCUT.png",
			Onpopular:     false,
			Description:   "컨템포러리 캐주얼",
			Created:       time.Now(),
			IsOpen:        true,
			InMaintenance: false, SizeGuide: []domain.SizeGuideDAO{
				{Label: "상의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/g-cut_TOP.png"},
				{Label: "하의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/g-cut_BOTTOM.png"},
			},
		},
		"V00130%3AALLSAINTS": {
			KorName:       "올세인츠",
			EngName:       "ALLSAINTS",
			KeyName:       "ALLSAINTS",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/ALLSAINTS.png",
			Onpopular:     false,
			Description:   "컨템포러리 디자이너",
			Created:       time.Now(),
			IsOpen:        true,
			InMaintenance: false, SizeGuide: []domain.SizeGuideDAO{
				{Label: "상의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/ALLSAINTS_TOP.png"},
				{Label: "하의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/ALLSAINTS_BOTTOM.png"},
			},
		},
		"71%3AVOV": {
			KorName:       "보브",
			EngName:       "VOV",
			KeyName:       "VOV",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/VOV.png",
			Onpopular:     false,
			Description:   "컨템포러리 캐주얼",
			Created:       time.Now(),
			IsOpen:        true,
			InMaintenance: false, SizeGuide: []domain.SizeGuideDAO{
				{Label: "상의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/VOV_TOP.png"},
				{Label: "하의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/VOV_BOTTOM.png"},
			},
		},
		"VP%3AVPPLEMENT": {
			KorName:       "브블먼트",
			EngName:       "VPPLEMENT",
			KeyName:       "VPPLEMENT",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/VPPLEMENT.png",
			Onpopular:     false,
			Description:   "라이프 캐주얼",
			Created:       time.Now(),
			IsOpen:        true,
			InMaintenance: false, SizeGuide: []domain.SizeGuideDAO{
				{Label: "상의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/VPPLEMENT_TOP.png"},
				{Label: "하의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/VPPLEMENT_BOTTOM.png"},
			},
		},
		"V00221%3AVIS+A+VIS": {
			KorName:       "비자비",
			EngName:       "VIS A VIS",
			KeyName:       "VISAVIS",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/VISAVIS.png",
			Onpopular:     false,
			Description:   "컨템포러리 캐주얼",
			Created:       time.Now(),
			IsOpen:        true,
			InMaintenance: false, SizeGuide: []domain.SizeGuideDAO{
				{Label: "상의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/VIS+A+VIS_TOP.png"},
				{Label: "하의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/VIS+A+VIS_BOTTOM.png"},
			},
		},
		"V00137%3ATATRAS": {
			KorName:       "타트라스",
			EngName:       "TATRAS",
			KeyName:       "TATRAS",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/TATRAS.png",
			Onpopular:     false,
			Description:   "컨템포러리 캐주얼",
			Created:       time.Now(),
			IsOpen:        false,
			InMaintenance: false, SizeGuide: []domain.SizeGuideDAO{
				{Label: "상의/하의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/TATRAS_TOP:BOTTOM.png"},
			},
		},
		"91%3ASTUDIO+TOMBOY": {
			KorName:       "스튜디오 톰보이",
			EngName:       "STUDIO TOMBOY",
			KeyName:       "STUDIOTOMBOY",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/STUDIOTOMBOY.png",
			Onpopular:     false,
			Description:   "라이프 캐주얼",
			Created:       time.Now(),
			IsOpen:        true,
			InMaintenance: false, SizeGuide: []domain.SizeGuideDAO{
				{Label: "상의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/STUDIO+TOMBOY_TOP.png"},
				{Label: "하의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/STUDIO+TOMBOY_BOTTOM.png"},
			},
		},
		// "24%3ASTELLA+McCARTNEY": {
		// 	KorName:       "스텔라 맥카트니",
		// 	EngName:       "STELLA McCARTNEY",
		// 	KeyName:       "STELLAMCCARTNEY",
		// 	LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/STELLAMCCARTNEY.png",
		// 	Onpopular:     false,
		// 	Description:   "컨템포러리 디자이너",
		// 	Created:       time.Now(),
		// 	IsOpen:        true,
		// 	InMaintenance: false, SizeGuide: []domain.SizeGuideDAO{
		// 		{Label: "상의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/STELLA+McCARTNEY_TOP.png"},
		// 		{Label: "하의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/STELLA+McCARTNEY_BOTTOM.png"},
		// 	},
		// },
		"12%3AST.JOHN": {
			KorName:       "세인트존",
			EngName:       "St.John",
			KeyName:       "STJOHN",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/STJOHN.png",
			Onpopular:     false,
			Description:   "컨템포러리 캐주얼",
			Created:       time.Now(),
			IsOpen:        false,
			InMaintenance: false, SizeGuide: []domain.SizeGuideDAO{
				{Label: "상의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/St.John_TOP.png"},
			},
		},
		"SD%3ASAVE+THE+DUCK": {
			KorName:       "세이브 더 덕",
			EngName:       "SAVE THE DUCK",
			KeyName:       "SAVETHEDUCK",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/SAVETHEDUCK.png",
			Onpopular:     false,
			Description:   "라이프 캐주얼",
			Created:       time.Now(),
			IsOpen:        true,
			InMaintenance: false, SizeGuide: []domain.SizeGuideDAO{
				{Label: "상의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/SAVE+THE+DUCK_TOP.png"},
			},
		},
		"64%3ASACAI": {
			KorName:       "사카이",
			EngName:       "sacai",
			KeyName:       "SACAI",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/SACAI.png",
			Onpopular:     false,
			Description:   "컨템포러리 디자이너",
			Created:       time.Now(),
			IsOpen:        true,
			InMaintenance: false, SizeGuide: []domain.SizeGuideDAO{
				{Label: "상의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/sacai_TOP.png"},
				{Label: "하의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/sacai_BOTTOM.png"},
			},
		},
		// "RO%3ARICK+OWENS": {
		// 	KorName:       "릭 오웬스",
		// 	EngName:       "Rick Owens",
		// 	KeyName:       "RICKOWENS",
		// 	LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/RICKOWENS.png",
		// 	Onpopular:     false,
		// 	Description:   "컨템포러리 디자이너",
		// 	Created:       time.Now(),
		// 	IsOpen:        true,
		// 	InMaintenance: false, SizeGuide: []domain.SizeGuideDAO{
		// 		{Label: "상의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/Rick+Owens_TOP.png"},
		// 		{Label: "하의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/Rick+Owens_BOTTOM.png"},
		// 	},
		// },
		"67%3AREISS": {
			KorName:       "리스",
			EngName:       "REISS",
			KeyName:       "REISS",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/REISS.png",
			Onpopular:     false,
			Description:   "컨템포러리",
			Created:       time.Now(),
			IsOpen:        true,
			InMaintenance: false, SizeGuide: []domain.SizeGuideDAO{
				{Label: "상의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/REISS_TOP.png"},
				{Label: "하의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/REISS_BOTTOM.png"},
			},
		},
		"40%3APROENZA+SCHOULER": {
			KorName:       "프로엔자 스쿨러",
			EngName:       "PROENZA SCHOULER",
			KeyName:       "PROENZASCHOULER",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/PROENZASCHOULER.png",
			Onpopular:     false,
			Description:   "컨템포러리 디자이너",
			Created:       time.Now(),
			IsOpen:        true,
			InMaintenance: false, SizeGuide: []domain.SizeGuideDAO{
				{Label: "상의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/PROENZA+SCHOULER_TOP.png"},
				{Label: "하의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/PROENZA+SCHOULER_BOTTOM.png"},
			},
		},
		// "57%3APAUL+SMITH": {
		// 	KorName:       "폴스미스",
		// 	EngName:       "Paul Smith",
		// 	KeyName:       "PAULSMITH",
		// 	LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/PAULSMITH.png",
		// 	Onpopular:     false,
		// 	Description:   "컨템포러리 디자이너",
		// 	Created:       time.Now(),
		// 	IsOpen:        true,
		// 	InMaintenance: false, SizeGuide: []domain.SizeGuideDAO{
		// 		{Label: "상의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/Paul+Smith_TOP.png"},
		// 		{Label: "하의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/Paul+Smith_BOTTOM.png"},
		// 	},
		// },
		"V00196%3ANOBIS": {
			KorName:       "노비스",
			EngName:       "nobis",
			KeyName:       "NOBIS",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/NOBIS.png",
			Onpopular:     false,
			Description:   "라이프 캐주얼",
			Created:       time.Now(),
			IsOpen:        true,
			InMaintenance: false, SizeGuide: []domain.SizeGuideDAO{
				{Label: "상의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/NOBIS_TOP.png"},
			},
		},
		"V00112%3AMOOSE+KNUCKLES": {
			KorName:       "무스너클",
			EngName:       "MOOSE KNUCKLES",
			KeyName:       "MOOSEKNUCKLES",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/MOOSEKNUCKLES.png",
			Onpopular:     false,
			Description:   "컨템포러리 디자이너",
			Created:       time.Now(),
			IsOpen:        true,
			InMaintenance: false, SizeGuide: []domain.SizeGuideDAO{
				{Label: "상의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/MOOSE+KNUCKLES_TOP.png"},
			},
		},
		// "31%3AMARNI": {
		// 	KorName:       "마르니",
		// 	EngName:       "MARNI",
		// 	KeyName:       "MARNI",
		// 	LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/MARNI.png",
		// 	Onpopular:     false,
		// 	Description:   "컨템포러리 디자이너",
		// 	Created:       time.Now(),
		// 	IsOpen:        true,
		// 	InMaintenance: false, SizeGuide: []domain.SizeGuideDAO{
		// 		{Label: "상의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/MARNI_TOP.png"},
		// 		{Label: "하의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/MARNI_BOTTOM.png"},
		// 	},
		// },
		// "23%3AMAISON+MARGIELA": {
		// 	KorName:       "메종 마르지엘라",
		// 	EngName:       "Maison Margiela",
		// 	KeyName:       "MAISONMARGIELA",
		// 	LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/MAISONMARGIELA.png",
		// 	Onpopular:     false,
		// 	Description:   "컨템포러리 디자이너",
		// 	Created:       time.Now(),
		// 	IsOpen:        true,
		// 	InMaintenance: false, SizeGuide: []domain.SizeGuideDAO{
		// 		{Label: "상의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/Maison+Margiela_TOP.png"},
		// 		{Label: "하의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/Maison+Margiela_BOTTOM.png"},
		// 	},
		// },
		"JC%3AJuicy+Couture": {
			KorName:       "쥬시꾸뛰르",
			EngName:       "Juicy Couture",
			KeyName:       "JUICYCOUTURE",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/JUICYCOUTURE.png",
			Onpopular:     false,
			Description:   "컨템포러리 캐주얼",
			Created:       time.Now(),
			IsOpen:        true,
			InMaintenance: false, SizeGuide: []domain.SizeGuideDAO{
				{Label: "상의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/Juicy+Couture_TOP.png"},
				{Label: "하의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/Juicy+Couture_BOTTOM.png"},
			},
		},
		"IA%3AILAIL": {
			KorName:       "일라일",
			EngName:       "ILAIL",
			KeyName:       "ILAIL",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/ILAIL.png",
			Onpopular:     false,
			Description:   "컨템포러리 캐주얼",
			Created:       time.Now(),
			IsOpen:        true,
			InMaintenance: false, SizeGuide: []domain.SizeGuideDAO{
				{Label: "상의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/ILAIL_TOP.png"},
				{Label: "하의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/ILAIL_BOTTOM.png"},
			},
		},
		"46%3AHERNO": {
			KorName:       "에르노",
			EngName:       "HERNO",
			KeyName:       "HERNO",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/HERNO.png",
			Onpopular:     false,
			Description:   "컨템포러리 디자이너",
			Created:       time.Now(),
			IsOpen:        true,
			InMaintenance: false, SizeGuide: []domain.SizeGuideDAO{
				{Label: "상의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/HERNO_TOP.png"},
				{Label: "하의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/HERNO_BOTTOM.png"},
			},
		},
		"A1%3AGIORGIO+ARMANI": {
			KorName:       "조르지오 아르마니",
			EngName:       "GIORGIO ARMANI",
			KeyName:       "GIORGIOARMANI",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/GIORGIOARMANI.png",
			Onpopular:     false,
			Description:   "클래식",
			Created:       time.Now(),
			IsOpen:        true,
			InMaintenance: false, SizeGuide: []domain.SizeGuideDAO{
				{Label: "상의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/GIORGIO+ARMANI_TOP.png"},
				{Label: "하의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/GIORGIO+ARMANI_BOTTOM.png"},
			},
		},
		"51%3AGAP+Adults": {
			KorName:       "갭",
			EngName:       "GAP",
			KeyName:       "GAP",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/GAP.png",
			Onpopular:     false,
			Description:   "라이프 캐주얼",
			Created:       time.Now(),
			IsOpen:        true,
			InMaintenance: false, SizeGuide: []domain.SizeGuideDAO{
				{Label: "상의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/GAP_TOP.png"},
				{Label: "하의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/GAP_BOTTOM.png"},
				{Label: "언더웨어", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/GAP_Underwear.png"},
			},
		},
		"A3%3AEMPORIO+ARMANI(D)": {
			KorName:       "엠포리오 아르마니",
			EngName:       "EMPORIO ARMANI",
			KeyName:       "EMPORIOARMANI",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/EMPORIOARMANI.png",
			Onpopular:     false,
			Description:   "컨템포러리 디자이너",
			Created:       time.Now(),
			IsOpen:        true,
			InMaintenance: false, SizeGuide: []domain.SizeGuideDAO{
				{Label: "상의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/EMPORIO+ARMANI_TOP.png"},
				{Label: "하의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/EMPORIO+ARMANI_BOTTOM.png"},
			},
		},
		"22%3ADSQUARED2": {
			KorName:       "디스퀘어드2",
			EngName:       "DSQUARED2",
			KeyName:       "DSQUARED2",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/DSQUARED2.png",
			Onpopular:     false,
			Description:   "진 캐주얼",
			Created:       time.Now(),
			IsOpen:        true,
			InMaintenance: false, SizeGuide: []domain.SizeGuideDAO{
				{Label: "상의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/DSQUARED2_TOP.png"},
				{Label: "하의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/DSQUARED2_BOTTOM.png"},
			},
		},
		"32%3ADRIES+VAN+NOTEN": {
			KorName:       "드리스반노튼",
			EngName:       "DRIES VAN NOTEN",
			KeyName:       "DRIESVANNOTEN",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/DRIESVANNOTEN.png",
			Onpopular:     false,
			Description:   "컨템포러리 디자이너",
			Created:       time.Now(),
			IsOpen:        true,
			InMaintenance: false, SizeGuide: []domain.SizeGuideDAO{
				{Label: "상의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/DIRES+VAN+NOTEN_TOP.png"},
				{Label: "하의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/DIRES+VAN+NOTEN_BOTTOM.png"},
			},
		},
		"13%3ADIESEL": {
			KorName:       "디젤",
			EngName:       "DIESEL",
			KeyName:       "DIESEL",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/DIESEL.png",
			Onpopular:     false,
			Description:   "진 캐주얼",
			Created:       time.Now(),
			IsOpen:        true,
			InMaintenance: false, SizeGuide: []domain.SizeGuideDAO{
				{Label: "상의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/DIESEL_TOP.png"},
				{Label: "하의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/DIESEL_BOTTOM.png"},
			},
		},
		"DL%3ADELLA+LANA": {
			KorName:       "델라라나",
			EngName:       "Della Lana",
			KeyName:       "DELLALANA",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/DELLALANA.png",
			Onpopular:     false,
			Description:   "컨템포러리",
			Created:       time.Now(),
			IsOpen:        true,
			InMaintenance: false, SizeGuide: []domain.SizeGuideDAO{
				{Label: "상의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/Della+Lana_TOP.png"},
				{Label: "하의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/Della+Lana_BOTTOM.png"},
			},
		},
		// "58%3AChlo%C3%A9": {
		// 	KorName:       "끌로에",
		// 	EngName:       "Chloé",
		// 	KeyName:       "CHLOE",
		// 	LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/CHLOE.png",
		// 	Onpopular:     false,
		// 	Description:   "컨템포러리 디자이너",
		// 	Created:       time.Now(),
		// 	IsOpen:        true,
		// 	InMaintenance: false, SizeGuide: []domain.SizeGuideDAO{
		// 		{Label: "상의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/Chloe_TOP.png"},
		// 		{Label: "하의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/Chloe_BOTTOM.png"},
		// 	},
		// },
		// "45%3ABRUNELLO+CUCINELLI": {
		// 	KorName:       "브루넬로 쿠치넬리",
		// 	EngName:       "Brunello Cucinelli",
		// 	KeyName:       "BRUNELLOCUCINELLI",
		// 	LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/BRUNELLOCUCINELLI.png",
		// 	Onpopular:     false,
		// 	Description:   "클래식",
		// 	Created:       time.Now(),
		// 	IsOpen:        true,
		// 	InMaintenance: false, SizeGuide: []domain.SizeGuideDAO{
		// 		{Label: "상의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/Brunello+Cucinelli_TOP.png"},
		// 		{Label: "하의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/Brunello+Cucinelli_BOTTOM.png"},
		// 	},
		// },
		"A4%3AARMANI+EXCHANGE": {
			KorName:       "아르마니 익스체인지",
			EngName:       "ARMANI EXCHANGE",
			KeyName:       "ARMANIEXCHANGE",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/ARMANIEXCHANGE.png",
			Onpopular:     false,
			Description:   "진 캐주얼",
			Created:       time.Now(),
			IsOpen:        true,
			InMaintenance: false, SizeGuide: []domain.SizeGuideDAO{
				{Label: "상의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/ARMANI+EXCHANGE_TOP.png"},
				{Label: "하의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/ARMANI+EXCHANGE_BOTTOM.png"},
			},
		},
		"36%3AALEXANDER+WANG": {
			KorName:       "알렉산더왕",
			EngName:       "alexanderwang",
			KeyName:       "ALEXANDERWANG",
			LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/ALEXANDERWANG.png",
			Onpopular:     false,
			Description:   "컨템포러리 디자이너",
			Created:       time.Now(),
			IsOpen:        true,
			InMaintenance: false, SizeGuide: []domain.SizeGuideDAO{
				{Label: "상의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/alexanderwang_TOP.png"},
				{Label: "하의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/alexanderwang_BOTTOM.png"},
			},
		},
		// "39%3AACNE+STUDIOS": {
		// 	KorName:       "아크네 스튜디오",
		// 	EngName:       "Acne Studios",
		// 	KeyName:       "ACNESTUDIOS",
		// 	LogoImgUrl:    "https://alloff.s3.ap-northeast-2.amazonaws.com/brands/ACNESTUDIOS.png",
		// 	Onpopular:     false,
		// 	Description:   "컨템포러리 디자이너",
		// 	Created:       time.Now(),
		// 	IsOpen:        true,
		// 	InMaintenance: false, SizeGuide: []domain.SizeGuideDAO{
		// 		{Label: "상의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/Acne+Studios_TOP.png"},
		// 		{Label: "하의", ImgUrl: "https://alloff.s3.ap-northeast-2.amazonaws.com/sizeguides/Acne+Studios_BOTTOM.png"},
		// 	},
		// },
	}

	for shopId, brand := range brands {
		_, err := ioc.Repo.Brands.Upsert(&brand)
		if err != nil {
			log.Println(err)
		}
		for keyname, catId := range categories {
			category := domain.CategoryDAO{
				Name:          keyname,
				KeyName:       brand.KeyName + "-" + keyname,
				CatIdentifier: catId,
				BrandKeyname:  brand.KeyName,
			}

			updatedCat, err := ioc.Repo.Categories.Upsert(&category)
			if err != nil {
				log.Println(err)
			}

			source := domain.CrawlSourceDAO{
				BrandKeyname:         brand.KeyName,
				BrandIdentifier:      shopId,
				MainCategoryKey:      "",
				Category:             *updatedCat,
				CrawlUrl:             "https://www.sivillage.com/dispctg/initDispCtg.siv?brand_info=" + shopId + "&disp_ctg_no=" + catId + "&outlet_yn=Y&goods_divi=all_goods&disp_clss_cd=10",
				CrawlModuleName:      modulename,
				IsSalesProducts:      true,
				IsForeignDelivery:    false,
				PriceMarginPolicy:    "NORMAL",
				DeliveryPrice:        0,
				EarliestDeliveryDays: 2,
				LatestDeliveryDays:   7,
				DeliveryDesc:         nil,
				RefundAvailable:      true,
				ChangeAvailable:      true,
				RefundFee:            5000,
				ChangeFee:            5000,
			}

			_, err = ioc.Repo.CrawlSources.Upsert(&source)
			if err != nil {
				log.Println(err)
			}
		}

	}
	log.Println("SI Village brands, categories & sources are added")
}
