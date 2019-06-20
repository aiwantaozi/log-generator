package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Pallinder/go-randomdata"
	"github.com/urfave/cli"
)

var VERSION = "v0.0.0-dev"

func main() {
	app := cli.NewApp()
	app.Name = "log-generator"
	app.Version = VERSION
	app.Usage = "You need help!"
	app.Action = func(c *cli.Context) error {
		server()
		return nil
	}

	app.Run(os.Args)
}

func sayHelloToError(w http.ResponseWriter, r *http.Request) {
	fmt.Println(fmt.Sprintf(`{"log", "Hi %s, you got error log, time: %s"}`, randomdata.FullName(randomdata.RandomGender), time.Now().Format(time.RFC1123Z)))
}

func sayHelloToInfo(w http.ResponseWriter, r *http.Request) {
	fmt.Println(fmt.Sprintf(`{"log", "Hi %s, normanl log, time: %s"}`, randomdata.FullName(randomdata.RandomGender), time.Now().Format(time.RFC1123Z)))
}

func server() {
	fmt.Println("Starting Server, listen at port 9090")
	http.HandleFunc("/error", sayHelloToError)
	http.HandleFunc("/info", sayHelloToInfo)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
