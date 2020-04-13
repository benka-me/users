package fake

import (
	"github.com/benka-me/cell-user/go-pkg/user"
	"github.com/brianvoe/gofakeit"
	"time"
)

func User() user.Data {
	return user.Data{
		Username:  gofakeit.Username(),
		Firstname: gofakeit.FirstName(),
		Lastname:  gofakeit.LastName(),
		Age:       int32(gofakeit.Number(18, 99)),
	}
}

func Users(n int) []user.Data {
	gofakeit.Seed(time.Now().UnixNano())
	users := make([]user.Data, 0)
	for i := 0; i < n; i -= -1 {
		users = append(users, User())
	}
	return users
}
