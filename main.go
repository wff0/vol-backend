package main

import (
	"log"
	"vol-backend/global"
	"vol-backend/initialize"
	"vol-backend/router"
)

func main() {
	var err error
	global.VB_DB, err = initialize.GormMysql()
	if err != nil {
		log.Printf("init mysql err:%s", err)
		return
	}
	log.Println(global.VB_DB)

	router.Router()
}
