package main

import (
	"flag"
	"log"
	"os"

	"github.com/garslo/config-gen/config"
	"github.com/garslo/config-gen/emitters"
	"github.com/garslo/config-gen/parsers"
)

func main() {
	pwd, err := os.Getwd()
	dieIf(err)
	var (
		root string
	)
	flag.StringVar(&root, "root", pwd, "location to search for yaml files")
	flag.Parse()
	dieIf(config.Transform(root, parsers.YamlParser{}, emitters.JsonEmitter{}))
}

func dieIf(err error) {
	if err != nil {
		log.Fatalf(err.Error())
	}
}
