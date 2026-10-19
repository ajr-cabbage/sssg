package main

import (
	"log"

	"github.com/ajr-cabbage/sssg/src/internal/build"
)

func main() {
	conf := build.Config{
		SourcePath:   "./content/docs/",
		TargetPath:   "./target/",
		AssetPath:    "./content/assets/",
		TemplatePath: "./templates/default.html",
	}
	err := conf.Build()
	if err != nil {
		log.Fatal(err)
	}
}
