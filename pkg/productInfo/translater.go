package productinfo

import (
	"github.com/lessbutter/alloff-api/config"
	"github.com/lessbutter/alloff-api/config/ioc"
	"go.mongodb.org/mongo-driver/bson"
	"go.uber.org/zap"
	"log"

	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/internal/pkg/translater"
	"golang.org/x/text/language"
)

func TranslateProductInfo(pdInfo *domain.ProductMetaInfoDAO) (*domain.ProductMetaInfoDAO, error) {
	titleInKorean, err := translater.TranslateText(language.Korean.String(), pdInfo.OriginalName)
	if err != nil {
		log.Println("err", err)
		return nil, err
	}

	informationKorean := map[string]string{}
	for key, value := range pdInfo.SalesInstruction.Information {
		keyKorean, err := translater.TranslateText(language.Korean.String(), key)
		if err != nil {
			config.Logger.Error("info translate key err", zap.Error(err))
			return nil, err
		}
		valueKorean, err := translater.TranslateText(language.Korean.String(), value)
		if err != nil {
			config.Logger.Error("info translate key err", zap.Error(err))
			return nil, err
		}
		informationKorean[keyKorean] = valueKorean
	}
	for key, value := range pdInfo.SalesInstruction.Description.Infos {
		if key == "소재" || key == "색상" {
			valueKorean, err := translater.TranslateText(language.Korean.String(), value)
			if err != nil {
				config.Logger.Error("info translate key err", zap.Error(err))
				return nil, err
			}
			pdInfo.SalesInstruction.Description.Infos[key] = valueKorean
		}
	}

	pdInfo.AlloffName = titleInKorean
	pdInfo.SalesInstruction.Information = informationKorean

	return pdInfo, nil
}

// TranslateRequiredProducts : 스크립터에서 사용
func TranslateRequiredProducts() {
	filter := bson.M{"istranslaterequired": true}
	pds, _, err := ioc.Repo.ProductMetaInfos.List(0, 100000, filter, bson.M{})
	if err != nil {
		log.Println(err)
	}

	for _, pd := range pds {
		log.Println(pd.OriginalName, "번역중")
		newPd, err := TranslateProductInfo(pd)
		if err != nil {
			log.Println(err)
		}

		newPd.IsTranslateRequired = false
		_, err = ioc.Repo.ProductMetaInfos.Upsert(newPd)
		if err != nil {
			log.Println(err)
		}
	}
}
