package main

import (
	"app/internal/stores"
	"app/internal/stores/models"
	"app/internal/stores/mysql"
	postgress "app/internal/stores/postgres"
)

func main() {
	p := postgress.NewConn()
	m := mysql.NewConn()
	u := models.User{
		Id:   1,
		Name: "Dev",
	}
	var s stores.UserRepository = p

	s.Create(u)
	s = m
	s.Create(u)
}
