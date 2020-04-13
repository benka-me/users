package rpc

import (
	"context"
	"fmt"
	"github.com/benka-me/users/go-pkg/hash"
	"github.com/benka-me/users/go-pkg/jwt"
	"github.com/benka-me/users/go-pkg/users"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (app *App) Login(ctx context.Context, req *users.LoginReq) (*users.LoginRes, error) {
	fmt.Println("login request")
	u := &users.User{}
	res := &users.LoginRes{}

	var err error

	filter := bson.D{{"data.username", req.Identifier}}
	err = app.MongoUsers.FindOne(ctx, filter).Decode(&u)
	if err != nil {
		return res, status.Error(codes.Code(34), "bad credential")
	}

	ok := hash.CheckPasswordHash(req.GetPassword(), u.Password)
	if ok {
		res.Data = &u.Data
	} else {
		return res, status.Error(codes.Code(34), "bad credential")
	}

	res.Auth, err = jwt.GenerateToken(res.Data.Username, res.Data.Email)
	if err != nil {
		return nil, status.Error(codes.Aborted, "cannot generate auth token")
	}

	return res, nil
}
