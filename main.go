package main

import (
	"flag"
	"log"
	"os"
)

func main() {
	pwd, err := os.Getwd()
	dieIf(err)
	var (
		root string
	)
	flag.StringVar(&root, "root", pwd, "location to search for yaml files")
	flag.Parse()
	resolver := Resolver{
		YamlParser{},
		JsonEmitter{},
	}
	dieIf(resolver.Emit(root))
}

func dieIf(err error) {
	if err != nil {
		log.Fatalf(err.Error())
	}
}
