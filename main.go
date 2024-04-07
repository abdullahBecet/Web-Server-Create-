package main

import (
	"fmt"
	"net/http"
	"yeni/env"
	"yeni/route"
	"yeni/utils"
)

func main() {
	env.InitEnv(env.LOCAL)
	fmt.Println(env.Getenv(env.PASSWORD))
	utils.InitTokenAuth()
	r := route.CreateServer()
	http.ListenAndServe(":8080", r)
}
