package db

import (
	"fmt"
	"github.com/benka-me/laruche/go-pkg/laruche/faker"
	"github.com/benka-me/users/go-pkg/config"
	"testing"
)

var db = Init(config.Init(true), true)

func Test1(t *testing.T) {
	t.Run("migrate db", func(t *testing.T) {
		db.DropTable("bees")
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
		{
			name: "bbbb",
			new:  BeeFrom(faker.AlphaBeesMap["benka-me/bbbb"]),
		},
	}
	tests[1].new.Deps = []*Bee{{ID: "benka-me/aaaa"}}
	for _, test := range tests {
		fmt.Println("test: ", test.new.ID)
		got := &Bee{}
		t.Run(test.name, func(t *testing.T) {
			got = db.Create(test.new).Value.(*Bee)
		})
		t.Run(test.name+"/find", func(t *testing.T) {
			find := &Bee{}
			db.Where("id = ?", got.ID).First(find)
			if find.ID != got.ID {
				t.Error(" b.ID != got.ID : ", find.ID, got.ID)
			}
		})
	}
}
