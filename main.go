package main

import (
	"blog/conf"
	"blog/routes"
)

func main() {
	conf.Init()
	var r = routes.NewRouter()
	_ = r.Run(conf.HttpPort)
}
