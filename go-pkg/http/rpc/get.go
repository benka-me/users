package rpc

import (
	"context"
	"github.com/benka-me/users/go-pkg/users"
)

func (app *App) Get(ctx context.Context, req *users.GetReq) (*users.Data, error) {
	u := &users.User{}
	err := u.Find(app.MongoUsers, ctx, req)

	return &u.Data, err
}

func (app *App) GetAll(ctx context.Context, req *users.Empty) (*users.All, error) {
	us, err := users.FindAll(app.MongoUsers, ctx)

	return us, err
}
