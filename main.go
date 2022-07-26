package main

import (
	"github.com/jackthepanda96/learn-deploy/config"
	"github.com/jackthepanda96/learn-deploy/infrastructure/database"
)

func main() {
	var cfg = config.GetConfig()
	var db = database.InitDB(cfg)
}
