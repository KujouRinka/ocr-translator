package main

import (
	"github.com/kujourinka/ocr-translator/config"
	"github.com/kujourinka/ocr-translator/engine"
)

func main() {
	cfg, err := config.LoadRawConfig("config/sample_config.yaml")
	if err != nil {
		panic(err)
	}

	e, err := engine.NewEngine(cfg)
	if err != nil {
		panic(err)
	}

	err = e.Run()
	if err != nil {
		panic(err)
	}
}
