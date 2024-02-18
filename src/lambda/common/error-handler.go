package common

import (
	"log"
	"os"
)

func LogError(msg string, err error) {
	if err != nil {
		log.Fatal(msg+": ", err)
		os.Exit(1)
	}
}
