package domain

type CrawlSourceDAO struct {
	ID string `bson:"_id,omitempty"`
	// 크롤링하는 브랜드의  Keyname
	BrandKeyname string
	// 브랜드 Identifier
	BrandIdentifier string
	//  여성의류 / 남성의류 / 키즈 같은 분류
	MainCategoryKey string
	// 크롤링하는 세부 카테고리
	Category CategoryDAO
	// 크롤링 main url
	CrawlUrl string
	// 어떤 몰에서 어떤 모듈로 크롤링하는지 내용
	CrawlModuleName string
	// 아울렛 상품인지 유무
	IsOutletProducts bool
}
