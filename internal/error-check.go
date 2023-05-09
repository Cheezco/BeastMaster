package internal

import (
	"BeastMaster/internal/debug"
	"log"
)

func CheckError(err error) {
	if err != nil {
		debug.Log(err)
		log.Fatal(err)
	}
}
