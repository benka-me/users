package rpc

import (
	"context"
	"github.com/benka-me/users/go-pkg/users"
)

func (a App) Register(context.Context, *users.RegisterReq) (*users.RegisterRes, error) {
	panic("implement me")
}
