package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/lessbutter/alloff-api/api/server"
	"github.com/lessbutter/alloff-api/api/server/model"
	"github.com/lessbutter/alloff-api/config/ioc"
)

func (r *queryResolver) Alloffcategories(ctx context.Context, input *model.AlloffCategoryInput) ([]*model.AlloffCategory, error) {
	alloffCatDaos, _ := ioc.Repo.AlloffCategories.List(input.ParentID)
	cats := []*model.AlloffCategory{}
	for _, catDao := range alloffCatDaos {
		cat := catDao.ToDTO()
		if cat != nil {
			cats = append(cats, cat)
		}
	}
	return cats, nil
}

func (r *queryResolver) Alloffcategory(ctx context.Context, input *model.AlloffCategoryID) (*model.AlloffCategory, error) {
	catDao, _ := ioc.Repo.AlloffCategories.Get(input.ID)
	return catDao.ToDTO(), nil
}

func (r *queryResolver) AlloffcategoryProducts(ctx context.Context, input model.CategoryProductsInput) (*model.AlloffCategoryProducts, error) {
	panic(fmt.Errorf("not implemented"))
}

// Query returns server.QueryResolver implementation.
func (r *Resolver) Query() server.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
