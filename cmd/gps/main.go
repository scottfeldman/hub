package main

import (
	"os"

	"github.com/merliot/dean"
	"github.com/merliot/hub/models/gps"
)

func main() {
	port, _ := os.LookupEnv("PORT")
	user, _ := os.LookupEnv("USER")
	passwd, _ := os.LookupEnv("PASSWD")

	thing := gps.New("foo", "gps", "foo").(*gps.Gps)

	server := dean.NewServer(thing)
	server.BasicAuth(user, passwd)

	server.DialWebSocket(user, passwd, "ws://127.0.0.1:8000/ws/1500", thing.Announce())

	if port != "" {
		server.Addr = ":" + port
		go server.ListenAndServe()
	}

	server.Run()
}
