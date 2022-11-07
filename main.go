package main

import (
	"example.com/m/pkg/api"
)

func main() {
	r := api.Router()
	r.Run(":3000")
}
