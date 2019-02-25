package main

import (
	"parser"
	"util"

	log "github.com/sirupsen/logrus"
)

func main() {
	util.LogFormatter()
	log.Info("Starting OmiseGO Plasma MoreVP CLI")
	parser.ParseArgs()
}
