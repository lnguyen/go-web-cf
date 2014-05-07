package main

import (
	"fmt"
	"github.com/cloudfoundry-community/go-cfenv"
	"github.com/go-martini/martini"
	"github.com/hoisie/redis"
)

func main() {
	appEnv := cfenv.Current()
	var client redis.Client
	redis, _ := appEnv.Services.FindByTagName("redis")
	client.Addr = redis.Credentials["hostname"] + ":" + redis.Credentials["port"]
	client.Password = redis.Credentials["password"]
	fmt.Println(redis)
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
