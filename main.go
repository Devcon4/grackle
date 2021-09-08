package main

import (
	"net/http"

	"github.com/Devcon4/grackle/packages/api"
	"github.com/Devcon4/grackle/packages/framework"
)

func main() {
	app := framework.NewServer(&http.Server{
		Addr: ":3000",
	})

	app.RegisterRouter(api.NewMainRouter())

	if err := app.ListenAndServe(); err != nil {
		panic(err)
	}
}
