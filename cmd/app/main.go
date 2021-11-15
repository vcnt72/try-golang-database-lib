package main

import "github.com/vcnt72/try-golang-database-lib/di"

func main() {
	app := di.NewApp()

	app.Run()
}
