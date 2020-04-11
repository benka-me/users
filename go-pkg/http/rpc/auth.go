package rpc

import (
	"context"
	"github.com/benka-me/users/go-pkg/users"
)

func (a App) Auth(context.Context, *users.Token) (*users.IsAuth, error) {
	panic("implement me")
}
