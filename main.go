package main

import (
	"app/pkg/common/db"
	"app/pkg/controller"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("./pkg/common/envs/.env")
	viper.ReadInConfig()

	port := viper.Get("PORT").(string)

	dbUrl := viper.Get("DBURL").(string)

	r := gin.Default()

	h := db.Init(dbUrl)

	controller.RegisterRoutes(r, h)

	r.Run(port)
}
