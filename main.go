package main

import (
	"github.com/tarokamikaze/kimoi-ddd/domain"
	"github.com/tarokamikaze/kimoi-ddd/infrastructure"
)

func main() {
	ds := domain.DomainService{Repo: &infrastructure.UserRepositoryImpl{}}
	if err := ds.Kimoi("newaddr@example.com"); err != nil {
		panic(err.Error())
	}
}
