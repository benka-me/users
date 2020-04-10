package rpc

import (
	"context"
	"fmt"
	"github.com/benka-me/users/go-pkg/users"
)

func (c *App) HelloWorld(ctx context.Context, req *users.Request) (*users.Greeting, error) {

	return &users.Greeting{Msg: fmt.Sprintf("Hello %s", req.Msg)}, nil
}
