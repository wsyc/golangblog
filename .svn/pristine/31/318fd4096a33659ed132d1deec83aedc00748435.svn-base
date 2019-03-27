package main

import (
	db "newland/datebase"
	ro "newland/router"
)

func main() {
	defer db.SqlDB.Close()
	router := ro.InitRouter()
	router.Run(":8000")
}
