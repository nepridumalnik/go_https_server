package main

import (
	"log"
	"site-1/web"
)

func main() {
	ws, err := web.MakeWebServer("config.json")

	if err != nil {
		log.Fatal(err)
	}

	err = ws.Run()

	if err != nil {
		log.Fatal(err)
	}
}
