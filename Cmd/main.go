package main

import (
	"HiringSystem/DataBaseService"
	"HiringSystem/Router"
)

func main() {
	DataBaseService.InitalDataBase()
	r := Router.Router()
	r.Run()
}
