package product

import (
	"github.com/lessbutter/alloff-api/config"
	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

func ProcessAddProductInfoRequests(request *AddMetaInfoRequest) error {
	pdInfo, err := checkNewProducts(request)
	if err != nil {
		return err
	}

	if pdInfo == nil {
		_, err := AddProductInfo(request)
		return err
	}

	_, err = UpdateProductInfo(pdInfo, request)
	return err
}

func checkNewProducts(request *AddMetaInfoRequest) (*domain.ProductMetaInfoDAO, error) {
	pdInfo, err := ioc.Repo.ProductMetaInfos.GetByProductID(request.Brand.KeyName, request.ProductID)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	}

	if err == nil {
		return pdInfo, nil
	}

	config.Logger.Error("err occured on get product meta info "+request.ProductID+" "+request.Brand.KeyName, zap.Error(err))
	return nil, err
}
