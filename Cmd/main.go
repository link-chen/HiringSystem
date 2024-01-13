package main

import (
	"HiringSystem/DataBaseService"
	"HiringSystem/Router"
	"HiringSystem/Utils"
)

func main() {
	DataBaseService.InitalDataBase()
	DataBaseService.InitRedis()
	Utils.CreateDir("Resume")
	r := Router.Router()
	r.Run()
}
