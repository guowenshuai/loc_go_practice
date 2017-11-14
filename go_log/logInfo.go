package main

import (
	log "github.com/Sirupsen/lo grus"
)

func main() {
	log.WithFields(log.Fields {
		"animal": "walrus",
	}).Info("A walrus appears")
}
