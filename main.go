package main

import "example.com/m/routes"

func main() {
	r := routes.Router()
	r.Run(":3000")
}
