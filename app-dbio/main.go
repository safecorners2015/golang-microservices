package main

import (
	"fmt"
	"net/http"

	"github.com/pottava/golang-microservices/app-dbio/app/config"
	_ "github.com/pottava/golang-microservices/app-dbio/app/controllers"
	"github.com/pottava/golang-microservices/app-dbio/app/logs"
)

func main() {
	cfg := config.NewConfig()
	logs.Debug.Print("[config] " + cfg.String())
	logs.Info.Printf("[service] listening on port %v", cfg.Port)
	logs.Fatal.Print(http.ListenAndServe(":"+fmt.Sprint(cfg.Port), nil))
}
