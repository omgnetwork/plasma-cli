package main

import (
	"github.com/omisego/plasma-cli/parser"
	"github.com/omisego/plasma-cli/util"
	log "github.com/sirupsen/logrus"
)

func main() {
	util.LogFormatter()
	log.Info("Starting OmiseGO Plasma MoreVP CLI")
	parser.ParseArgs()
}
