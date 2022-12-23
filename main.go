package main

import (
	"dewetour/database"
	"dewetour/pkg/mysql"
)

func main() {

	mysql.DatabaseInit()

	database.RunMigration()
}
