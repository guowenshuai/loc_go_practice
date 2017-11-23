package main

import (
	"fmt"

	log "github.com/cihub/seelog"
)

var Logger log.LoggerInterface

func seelogSay() {
	defer log.Flush()
	log.Info("Hello from Seelog!")
}

func init() {
	DisableLog()
	loadAppConfig()
}

func DisableLog() {
	Logger = log.Disabled
	loadAppConfig()
}

func UseLogger(newLogger log.LoggerInterface) {
	Logger = newLogger
}

func loadAppConfig() {
	appConfig := `
<seelog minlevel="warn">
	<outputs formatid="common">
		<rollingfile type="size" filename=
	`
}
