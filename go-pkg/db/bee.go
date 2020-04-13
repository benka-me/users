package db

import (
	"github.com/benka-me/laruche/go-pkg/laruche"
	"time"
)

type Bee struct {
	ID           string `gorm:"primary_key"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time `sql:"index"`
	Name         string     `gorm:""`
	PkgName      string     `gorm:""`
	PkgNameCamel string     `gorm:""`
	Repo         string     `gorm:""`
	Author       string     `gorm:""`
	Port         int        `gorm:""`
	Public       bool       `gorm:""`
	License      string     `gorm:""`
	Description  string     `gorm:""`
	Keywords     string     `gorm:""`
	Tag          string     `gorm:""`
	DevLang      int        `gorm:""`
	Languages    int        `gorm:""`
	IsGateway    bool       `gorm:""`
	Deps         []*Bee     `gorm:"many2many:dependencies;association_jointable_foreignkey:dep_id"`
	//Cons         []*Bee     `gorm:"many2many:bee_bee;"`
}

func BeeFrom(from *laruche.Bee) *Bee {
	return &Bee{
		ID:           from.GetNamespaceStr(),
		CreatedAt:    time.Time{},
		UpdatedAt:    time.Time{},
		DeletedAt:    nil,
		Name:         from.Name,
		PkgName:      from.PkgName,
		PkgNameCamel: from.PkgNameCamel,
		Repo:         from.Repo,
		Author:       from.Author,
		Port:         int(from.Port),
		Public:       from.Public,
		License:      from.License,
		Description:  from.Description,
		Keywords:     from.Keywords,
		Tag:          from.Tag,
		DevLang:      int(from.DevLang),
		Languages:    0,
		IsGateway:    from.IsGateway,
		//Deps:         format(from.Deps),
		//Cons:         format(from.Cons),
	}
}

func format(deps []string) []*Bee {
	ret := make([]*Bee, len(deps))
	for i, d := range deps {
		ret[i] = &Bee{ID: d}
	}
	return ret
}
