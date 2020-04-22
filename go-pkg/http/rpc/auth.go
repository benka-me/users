package rpc

import (
	"context"
	"fmt"
	"github.com/benka-me/users/go-pkg/jwt"
	"github.com/benka-me/users/go-pkg/users"
)

func (app *App) Auth(ctx context.Context, req *users.Token) (*users.IsAuth, error) {
	err := jwt.CheckJwt(req.Val)
	fmt.Println("auth err:", err, req.Val)

	return &users.IsAuth{Val: err == nil}, nil
}
