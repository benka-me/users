package db

import (
	"github.com/benka-me/laruche/go-pkg/laruche/faker"
	"github.com/benka-me/users/go-pkg/config"
	"testing"
)

var db = Init(config.Init(true), true)

func Test1(t *testing.T) {
	t.Run("migrate db", func(t *testing.T) {
		db.AutoMigrate(&Bee{})
	})
	tests := []struct {
		name string
		new  *Bee
	}{
		{
			name: "aaaa",
			new:  BeeFrom(faker.AlphaBeesMap["benka-me/aaaa"]),
		},
	}
	for _, test := range tests {
		got := &Bee{}
		t.Run(test.name, func(t *testing.T) {
			got = db.Create(test.new).Value.(*Bee)
		})
		t.Run(test.name+"/recover", func(t *testing.T) {
			b := &Bee{}
			db.Find(b, got.ID)
		})
	}
}
