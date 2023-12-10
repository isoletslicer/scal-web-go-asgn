package main

import (
	db "assignment2/database"
	routers "assignment2/routers"
)

func main() {
	db.Init()
	routers.ServerInit().Run(":3003")
}
