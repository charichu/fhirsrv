package main

import (
	"idberlin/fhirsrv/app"
)

func main() {
	app := app.App{}
	app.Initialize("", "", "")

	app.Run(":8080")
}
