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

		fmt.Println("Register process revceived :", r.Username)
		pwd, _ := hash.HashPassword(r.Password)
		entry := users.User{
			Data: users.Data{
				Username: r.Username,
				Email:    r.Email,
			},
			Password: pwd,
		}

		if app.usernameAvailable(context.TODO(), r) {
			inserted, err := app.MongoUsers.InsertOne(context.TODO(), entry)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("inserted: ", inserted.InsertedID)
			}

			//TODO: Send email validation
		} else {
			fmt.Println("cant insert: ", r.Username)
			//TODO: Implement case of someone registered faster
		}
	}
}

func (app *App) usernameAvailable(ctx context.Context, req *users.RegisterReq) bool {
	filter := bson.D{{"data.username", req.Username}}
	return nil != app.MongoUsers.FindOne(context.TODO(), filter).Err()
}
func (app *App) emailAvailable(ctx context.Context, req *users.RegisterReq) bool {
	filter := bson.D{{"data.email", req.Username}}
	return nil != app.MongoUsers.FindOne(context.TODO(), filter).Err()
}

func (app *App) Register(ctx context.Context, req *users.RegisterReq) (*users.RegisterRes, error) {
	if !app.usernameAvailable(ctx, req) {
		return &users.RegisterRes{}, status.Error(codes.AlreadyExists, "username already used")
	}
	if !app.emailAvailable(ctx, req) {
		return &users.RegisterRes{}, status.Error(codes.AlreadyExists, "email already used")
	}
	//TODO validate pwd / email / username

	//TODO add event register request

	app.RegisterChan <- req

	return &users.RegisterRes{}, status.Error(codes.OK, "ok")
}
