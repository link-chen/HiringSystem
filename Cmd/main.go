package main

import (
	"HiringSystem/DataBaseService"
	"HiringSystem/Router"
	"HiringSystem/Utils"
)

func main() {
	DataBaseService.InitalDataBase()
	Utils.CreateDir("Resume")
	r := Router.Router()
	r.Run()
}
