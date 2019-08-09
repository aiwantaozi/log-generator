package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/Pallinder/go-randomdata"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

const VERSION = "v0.0.0-dev"

const jsonLogDir = "/var/mylogs/json"
const textLogDir = "/var/mylogs/text"

const infoJSONLogPath = "/var/mylogs/json/info"
const errorJSONLogPath = "/var/mylogs/json/error"

const infoTextLogPath = "/var/mylogs/text/info"
const errorTextLogPath = "/var/mylogs/text/error"

var infoJSONLog = log.New()
var errorJSONLog = log.New()
var infoTextLog = log.New()
var errorTextLog = log.New()
var jsonFormat = &log.JSONFormatter{DisableTimestamp: true}

func main() {
	app := cli.NewApp()
	app.Name = "log-generator"
	app.Version = VERSION
	app.Usage = "You need help!"
	app.Action = func(c *cli.Context) error {
		if err := os.MkdirAll(jsonLogDir, 0777); err != nil {
			return err
		}

		fmt.Println("2")
		if err := os.MkdirAll(textLogDir, 0777); err != nil {
			return err
		}

		fmt.Println("3")
		infoJSONLogFile, err := os.OpenFile(infoJSONLogPath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
		if err != nil {
			return err
		}
		infoJSONLog.SetOutput(infoJSONLogFile)
		infoJSONLog.SetFormatter(jsonFormat)
		defer infoJSONLogFile.Close()

		errorJSONLogFile, _ := os.OpenFile(errorJSONLogPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		errorJSONLog.SetOutput(errorJSONLogFile)
		errorJSONLog.SetFormatter(jsonFormat)
		defer errorJSONLogFile.Close()

		infoTextLogFile, _ := os.OpenFile(infoTextLogPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		infoTextLog.SetOutput(infoTextLogFile)
		defer infoTextLogFile.Close()

		errorTextLogFile, _ := os.OpenFile(errorTextLogPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		errorTextLog.SetOutput(errorTextLogFile)
		defer errorTextLogFile.Close()

		server()
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func sayHelloToError(w http.ResponseWriter, r *http.Request) {
	str := genHelloTo("error")
	fmt.Println(fmt.Sprintf(`{"log", "%s"}`, str))
	errorJSONLog.Errorln(str)
}

func sayHelloToInfo(w http.ResponseWriter, r *http.Request) {
	str := genHelloTo("info")
	infoJSONLog.Infoln(str)
	fmt.Println(fmt.Sprintf(`{"log", "%s"}`, str))
}

func sayHelloToErrorText(w http.ResponseWriter, r *http.Request) {
	str := genHelloTo("error")
	fmt.Println(str)
	errorTextLog.Errorln(str)
}

func sayHelloToInfoText(w http.ResponseWriter, r *http.Request) {
	str := genHelloTo("infp")
	fmt.Println(genHelloTo("info"))
	infoTextLog.Errorln(str)
}

func genHelloTo(level string) string {
	return fmt.Sprintf("Hi %s, %s log, time: %s", randomdata.FullName(randomdata.RandomGender), level, time.Now().Format(time.RFC1123Z))
}

func server() {
	fmt.Println("Starting Server, listen at port 9090")
	http.HandleFunc("/error/json", sayHelloToError)
	http.HandleFunc("/info/json", sayHelloToInfo)
	http.HandleFunc("/error/text", sayHelloToErrorText)
	http.HandleFunc("/info/text", sayHelloToInfoText)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
