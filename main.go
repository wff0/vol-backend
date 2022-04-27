package main

import (
	"log"
	"vol-backend/global"
	"vol-backend/initialize"
	"vol-backend/router"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	var err error
	global.VB_DB, err = initialize.GormMysql()
	if err != nil {
		log.Printf("init mysql err:%s", err)
		return
	}
	log.Println(global.VB_DB)
	initialize.Redis()
	rc := global.VB_REDIS_POOL.Get()
	log.Println(rc)
	defer rc.Close()

	router.Router()
}
