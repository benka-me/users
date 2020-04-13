package rpc

import (
	"context"
	"fmt"
	"github.com/benka-me/users/go-pkg/hash"
	"github.com/benka-me/users/go-pkg/users"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (app *App) insertRegisterProcess() {
	for {
		r := <-app.RegisterChan

		pwd, _ := hash.HashPassword(r.Password)
		entry := users.User{
			Data: users.Data{
				Username: r.Username,
				Email:    r.Email,
			},
			Password: pwd,
		}

		if app.available(context.TODO(), r) {
			_, _ = app.MongoUsers.InsertOne(context.TODO(), entry)

			//TODO: Send email validation
		} else {
			fmt.Println("cant insert: ", r.Username)
			//TODO: Implement case of someone registered faster
		}
	}
}

func (app *App) available(ctx context.Context, req *users.RegisterReq) bool {
	filter := bson.D{
		{"$or", bson.A{
			bson.D{{"data.username", req.Username}},
			bson.D{{"data.email", req.Email}},
		}},
	}

	return nil != app.MongoUsers.FindOne(context.TODO(), filter).Err()
}

func (app *App) Register(ctx context.Context, req *users.RegisterReq) (*users.RegisterRes, error) {
	if !app.available(ctx, req) {
		return &users.RegisterRes{}, status.Error(codes.AlreadyExists, "username already exist")
	}
	//TODO validate pwd / email / username

	//TODO add event register request

	app.RegisterChan <- req

	return &users.RegisterRes{}, status.Error(codes.OK, "ok")
}
