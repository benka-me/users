package users

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type User struct {
	ObjectId primitive.ObjectID `bson:"_id,omitempty"`
	Password string
	Data
}

func (u *User) Find(users *mongo.Collection, ctx context.Context, identifier *GetReq) error {
	var err error

	filter := bson.D{{"data.username", identifier.GetName()}}
	err = users.FindOne(context.TODO(), filter).Decode(&u)
	if err != nil {
		return status.Error(codes.NotFound, "user not found")
	} else {
		u.Id = u.ObjectId.Hex()
		return nil
	}
}

func FindAll(users *mongo.Collection, ctx context.Context) (*All, error) {
	all := &All{
		Users: make([]*Data, 0),
	}
	cur, err := users.Find(ctx, bson.D{})
	if err != nil {
		return all, err
	}

	for cur.Next(ctx) {
		var user User
		err := cur.Decode(&user)
		if err != nil {
			return all, err
		}
		user.Id = user.ObjectId.Hex()
		all.Users = append(all.Users, &user.Data)
	}
	return all, err
}
