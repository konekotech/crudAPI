package main

import (
	"github.com/konekotech/crudAPI/db"
	"github.com/konekotech/crudAPI/server"
)

func main() {
	db.Init()
	server.Init()
	defer db.Close()
}
