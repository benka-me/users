package rpc

import (
	"context"
	"github.com/benka-me/users/go-pkg/users"
)

func (a App) Get(context.Context, *users.GetReq) (*users.Data, error) {
	panic("implement me")
}

func (a App) GetAll(context.Context, *users.Empty) (*users.All, error) {
	panic("implement me")
}
