package main

import (
	"github.com/omise-go/plasma-cli/parser"
	"github.com/omise-go/plasma-cli/util"
	log "github.com/sirupsen/logrus"
)

func main() {
	util.LogFormatter()
	log.Info("Starting OmiseGO Plasma MoreVP CLI")
	parser.ParseArgs()
}
