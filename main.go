package main

import (
	cfenv "github.com/cloudfoundry-community/go-cfenv"
	"github.com/go-martini/martini"
	"github.com/hoisie/redis"
)

func main() {
	appEnv, _ := cfenv.Current()
	var client redis.Client
	services, _ := appEnv.Services.WithTag("redis")
	redis := services[0]
	client.Addr = redis.Credentials["hostname"] + ":" + redis.Credentials["port"]
	client.Password = redis.Credentials["password"]
	m := martini.Classic()
	m.Post("/:name", func(params martini.Params) string {
		client.Set(params["name"], []byte(params["name"]))
		return "Received " + params["name"]
	})
	m.Get("/:name", func(params martini.Params) string {
		val, _ := client.Get(params["name"])
		return "Hello " + string(val)
	})
	m.Run()
}
