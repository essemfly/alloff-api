package hometab

import (
	"errors"
	"time"

	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddHometabItem(req *HomeTabItemRequest) (*domain.HomeTabItemDAO, error) {
	if req.Requester == nil {
		return nil, errors.New("requester field not set")
	}

	item := &domain.HomeTabItemDAO{
		ID:           primitive.NewObjectID(),
		Title:        req.Title,
		Description:  req.Description,
		Tags:         req.Tags,
		BackImageUrl: req.BackImageUrl,
		StartedAt:    req.StartedAt,
		EndedAt:      req.EndedAt,
		CreatedAt:    time.Time{},
		UpdatedAt:    time.Time{},
	}

	item = req.Requester.fillItemContents(item)

	result, err := ioc.Repo.HomeTabItems.Insert(item)
	if err != nil {
		return nil, err
	}

	return result, nil
}
