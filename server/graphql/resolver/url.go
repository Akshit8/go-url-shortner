package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/Akshit8/url-shortner/server/graphql/generated"
	"github.com/Akshit8/url-shortner/server/graphql/model"
)

func (r *mutationResolver) CreateURL(ctx context.Context, input model.CreateURL) (*model.URL, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateURL(ctx context.Context, code string, input model.UpdateURL) (*model.URL, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteURL(ctx context.Context, code string) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetURLByID(ctx context.Context, code string) (*model.URL, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetAllURL(ctx context.Context, limit *int, offset *int) ([]*model.URL, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
