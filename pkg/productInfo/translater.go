package productinfo

import (
	"github.com/lessbutter/alloff-api/config"
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
