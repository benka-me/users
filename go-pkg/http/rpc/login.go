package rpc

import (
	"context"
	"github.com/benka-me/users/go-pkg/users"
)

func (a App) Login(context.Context, *users.LoginReq) (*users.LoginRes, error) {
	panic("implement me")
}
