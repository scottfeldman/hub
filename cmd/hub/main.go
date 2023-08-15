package main

import (
	"os"

	"github.com/merliot/dean"
	"github.com/merliot/sw-poc/models/hub"
	"github.com/merliot/sw-poc/models/ps30m"
	"github.com/merliot/sw-poc/models/gps"
)

func main() {

	hub := hub.New("swpoc01", "hub", "swpoc01").(*hub.Hub)

	server := dean.NewServer(hub)
	hub.UseServer(server)

	server.Addr = ":8000"
	if port, ok := os.LookupEnv("PORT"); ok {
		server.Addr = ":" + port
	}

	if user, ok := os.LookupEnv("USER"); ok {
		if passwd, ok := os.LookupEnv("PASSWD"); ok {
			server.BasicAuth(user, passwd)
		}
	}

	server.RegisterModel("ps30m", ps30m.New)
	server.RegisterModel("gps", gps.New)

	go server.ListenAndServe()
	server.Run()
}
