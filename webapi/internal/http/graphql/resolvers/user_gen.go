package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.47

import (
	"context"
	"fmt"
	"webapi/internal/entity"
	graph_gen "webapi/internal/http/graphql/generated"
)

// SignIn is the resolver for the signIn field.
func (r *mutationResolver) SignIn(ctx context.Context, input graph_gen.SignInInput) (*graph_gen.Session, error) {
	panic(fmt.Errorf("not implemented: SignIn - signIn"))
}

// SignUp is the resolver for the signUp field.
func (r *mutationResolver) SignUp(ctx context.Context, input graph_gen.SignUpInput) (*graph_gen.Session, error) {
	panic(fmt.Errorf("not implemented: SignUp - signUp"))
}

// Me is the resolver for the me field.
func (r *queryResolver) Me(ctx context.Context) (*entity.User, error) {
	panic(fmt.Errorf("not implemented: Me - me"))
}
