package main

import (
	c "hi-story/config"
	"hi-story/routes"
)

func main(){
	c.InitDB()
	c.InitMigrate()
	e := routes.New()

	e.Start(":8000")
}