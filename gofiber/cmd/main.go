package main

import internal "github.com/risoluzumaki/jwt/go/internal/bootstrapp"

func main() {
	app := internal.NewApp()
	app.Listen(":3000")
}
