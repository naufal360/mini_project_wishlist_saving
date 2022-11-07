package main

import (
	"os"
	r "wishlist/route"
)

func main() {

	routes := r.Init()

	port := ":" + os.Getenv("PORT")

	routes.Logger.Fatal(routes.Start(port))
}
