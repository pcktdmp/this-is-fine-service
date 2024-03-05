package main

import (
	log "github.com/charmbracelet/log"
	"net/http"
	"os"
	"strings"
)

var listStr = os.Getenv("THIS_IS_FINE_STDOUT_LOGGING")
var logItems = strings.Split(listStr, ",")

func handler(w http.ResponseWriter, r *http.Request) {
	for _, logItem := range logItems {
		logMessages := strings.Split(logItem, ":")
		if len(logMessages) != 2 {
			log.Fatal("Please define log level via 'INFO:example-log-message'. Supported levels are: INFO, WARN, ERROR and DEBUG")
		}
		switch sev := logMessages[0]; sev {
		case "INFO":
			log.Info(logMessages[1])
		case "WARN":
			log.Warn(logMessages[1])
		case "ERROR":
			log.Error(logMessages[1])
		case "DEBUG":
			log.Debug(logMessages[1])
		default:
			log.Print(logMessages[1])
		}
	}

}

func main() {
	if listStr == "" {
		log.Fatal("Environment variable THIS_IS_FINE_STDOUT_LOGGING is not set or is empty. Allowed format: THIS_IS_FINE_STDOUT_LOGGING=<LEVEL>:<MESSAGE>,<LEVEL>:<MESSAGE>")
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
