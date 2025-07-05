package main

import(
	"yuukisan/yuukisan/config"
	"yuukisan/yuukisan/database"
	"yuukisan/yuukisan/routes"
)

func main() {
	config.LoadEnv()

	database.InitDB()

	r := routes.SetupRouter()

	r.Run("0.0.0.0:" + config.GetEnv("APP_PORT", "3000"))
}