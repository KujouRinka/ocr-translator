package main

import (
	"fmt"

	"github.com/otiai10/gosseract/v2"

	"github.com/kujourinka/ocr-translator/config"
	"github.com/kujourinka/ocr-translator/engine"
)

func main() {
	fmt.Println("Tesseract version:", gosseract.Version())

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
