package engine

import (
	"fmt"
	"time"

	"github.com/kujourinka/ocr-translator/config"
	"github.com/kujourinka/ocr-translator/ocr"
	"github.com/kujourinka/ocr-translator/scanner"
	"github.com/kujourinka/ocr-translator/translator"
)

type Engine struct {
	config *config.Config

	scanner scanner.Scanner
}

func NewEngine(cfg *config.RawConfig) (*Engine, error) {
	c := &config.Config{}

	var err error

	c.OCR, err = ocr.NewGosseractOcr(cfg.OCR.Lang...)
	if err != nil {
		return nil, err
	}

	for _, t := range cfg.Translators {
		switch t.Type {
		case "google":
			tl, err := translator.NewGoogleTranslator(
				config.StrToLang(t.TargetLang),
				config.StrToLang(t.SourceLang),
				t.APIKey,
			)
			if err != nil {
				return nil, err
			}
			c.Translators = append(c.Translators, tl)
		}
	}

	s := scanner.NewDefaultScanner(-1, -1, -1, -1)

	return &Engine{
		config:  c,
		scanner: s,
	}, nil
}

func (e *Engine) Run() error {
	for {
		// TODO: lots of things to do
		img, err := e.scanner.Scan()
		if err != nil {
			return err
		}

		text, err := e.config.OCR.ImgToText(img)
		if err != nil {
			return err
		}

		for _, tl := range e.config.Translators {
			transText, err := tl.Translate(text)
			if err != nil {
				return err
			}
			fmt.Println(text, "->", transText)
		}

		time.Sleep(1 * time.Second)
	}
	return nil
}

func (e *Engine) Config() *config.Config {
	return e.config
}
