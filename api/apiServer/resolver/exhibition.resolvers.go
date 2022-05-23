package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/lessbutter/alloff-api/api/apiServer/mapper"
	"github.com/lessbutter/alloff-api/api/apiServer/middleware"
	"github.com/lessbutter/alloff-api/api/apiServer/model"
	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/internal/pkg/alimtalk"
)

func (r *mutationResolver) SetAlarm(ctx context.Context, id string) (bool, error) {
	user := middleware.ForContext(ctx)
	if user == nil {
		return false, fmt.Errorf("ERR000:invalid token")
	}

	exhibitionDao, err := ioc.Repo.Exhibitions.Get(id)
	if err != nil {
		return false, err
	}

	alimtalkRegisterd, err := alimtalk.ChangeExhibitionNotifyStatus(user, exhibitionDao)
	if err != nil {
		return false, err
	}

	// 기존의 알림톡을 취소하는 경우
	if alimtalkRegisterd == nil {
		return false, nil
	}
	return true, nil
}

func (r *queryResolver) Exhibition(ctx context.Context, id string) (*model.Exhibition, error) {
	exhibitionDao, err := ioc.Repo.Exhibitions.Get(id)
	if err != nil {
		return nil, err
	}

	return mapper.MapExhibition(exhibitionDao, false), nil
}

func (r *queryResolver) Exhibitions(ctx context.Context, input model.ExhibitionInput) (*model.ExhibitionOutput, error) {
	offset, limit := 0, 100
	query := ""

	var exhibitionStatus domain.ExhibitionStatus
	switch input.Status {
	case model.ExhibitionStatusNotOpen:
		exhibitionStatus = domain.EXHIBITION_NOTOPEN
	case model.ExhibitionStatusLive:
		exhibitionStatus = domain.EXHIBITION_LIVE
	case model.ExhibitionStatusClosed:
		exhibitionStatus = domain.EXHIBITION_CLOSED
	}

	// only live
	liveDaos, liveCnt, err := ioc.Repo.Exhibitions.List(offset, limit, false, domain.EXHIBITION_LIVE, domain.EXHIBITION_TIMEDEAL, query)
	if err != nil {
		return nil, err
	}
	// not open
	notOpenDaos, notOpenCnt, err := ioc.Repo.Exhibitions.List(offset, limit, false, domain.EXHIBITION_NOTOPEN, domain.EXHIBITION_TIMEDEAL, query)
	if err != nil {
		return nil, err
	}

	exs := []*model.Exhibition{}

	if exhibitionStatus == domain.EXHIBITION_LIVE {
		for _, exhibitionDao := range liveDaos {
			exs = append(exs, mapper.MapExhibition(exhibitionDao, true))
		}
	} else {
		for _, exhibitionDao := range notOpenDaos {
			exs = append(exs, mapper.MapExhibition(exhibitionDao, true))
		}
	}

	return &model.ExhibitionOutput{
		Exhibitions:   exs,
		Status:        input.Status,
		LiveCounts:    liveCnt,
		NotOpenCounts: notOpenCnt,
	}, nil
}
