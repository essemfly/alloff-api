package main

import (
	"log"

	"github.com/lessbutter/alloff-api/config"
	"github.com/lessbutter/alloff-api/internal/storage/mongo"
	"github.com/lessbutter/alloff-api/pkg/seeder/malls"
)

func main() {
	conf := config.GetConfiguration()
	log.Println(conf)

	conn := mongo.NewMongoDB(conf)
	conn.RegisterRepos()

	malls.AddLotteFashion()
	malls.AddSSFMall()
	malls.AddHandsome()
	malls.AddIdLook()
	malls.AddSiVillages()
	malls.AddKolonMall()
	malls.AddMichaa()
	malls.AddBabathe()
	malls.AddHfashion()
	malls.AddIDFMall()
	malls.AddDaehyun()
	malls.AddNiceClaup()
	malls.AddLacoste()
	malls.AddSisley()
	malls.AddBenetton()
	malls.AddLoungeB()
	malls.AddBylynn()
	malls.AddTheAMall()
	malls.AddOthers()
}
