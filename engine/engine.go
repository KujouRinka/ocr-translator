package engine

import (
	"fmt"
	"math"
	"time"

	"golang.design/x/hotkey"

	"github.com/kujourinka/ocr-translator/config"
	"github.com/kujourinka/ocr-translator/ocr"
	"github.com/kujourinka/ocr-translator/scanner"
	"github.com/kujourinka/ocr-translator/scanner/mouse"
	"github.com/kujourinka/ocr-translator/translator"
)

type Engine struct {
	config *config.Config

	scanner   scanner.Scanner
	scanDelay int
	startHk   *hotkey.Hotkey
	endHk     *hotkey.Hotkey
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
				t.Socks5,
			)
			if err != nil {
				return nil, err
			}
			c.Translators = append(c.Translators, tl)
		}
	}

	s := scanner.NewDefaultScanner(0, 0, 0, 0)

	startHotkey := hotkey.New([]hotkey.Modifier{hotkey.ModCtrl, hotkey.ModShift}, hotkey.KeyO)
	err = startHotkey.Register()
	if err != nil {
		return nil, err
	}
	endHotkey := hotkey.New([]hotkey.Modifier{hotkey.ModCtrl, hotkey.ModShift}, hotkey.KeyM)
	err = endHotkey.Register()
	if err != nil {
		return nil, err
	}

	if cfg.OCR.ScanDelay == 0 {
		cfg.OCR.ScanDelay = 1000
	}

	return &Engine{
		config:    c,
		scanner:   s,
		scanDelay: cfg.OCR.ScanDelay,
		startHk:   startHotkey,
		endHk:     endHotkey,
	}, nil
}

func (e *Engine) Run() error {
	var lastText string

	// update scanner window
	go func() {
		for {
			select {
			case <-e.startHk.Keydown():
				x, y := mouse.Location()
				fmt.Println("start point:", x, y)
				err := e.scanner.SetX(x)
				if err != nil {
					continue
				}
				err = e.scanner.SetY(y)
				if err != nil {
					continue
				}
			case <-e.endHk.Keydown():
				x, y := mouse.Location()
				fmt.Println("end point:", x, y)
				startX, startY := e.scanner.GetX(), e.scanner.GetY()
				err := e.scanner.SetWidth(x - startX)
				if err != nil {
					continue
				}
				err = e.scanner.SetHeight(y - startY)
				if err != nil {
					continue
				}
			}
		}
	}()

	for {
		time.Sleep(time.Millisecond * time.Duration(e.scanDelay))
		// TODO: lots of things to do
		img, err := e.scanner.Scan()
		if err != nil {
			fmt.Println(err)
			continue
		}

		text, err := e.config.OCR.ImgToText(img)
		if err != nil {
			fmt.Println(err)
			continue
		}

		text = ocr.Format(text)
		if !needTranslate(text, lastText) {
			continue
		}
		lastText = text

		fmt.Println(text, "->")
		for _, tl := range e.config.Translators {
			transText, err := tl.Translate(text)
			if err != nil {
				return err
			}
			fmt.Println(transText)
		}

	}
	return nil
}

func (e *Engine) Config() *config.Config {
	return e.config
}

func needTranslate(s1, s2 string) bool {
	s1Map := make(map[rune]int)
	s2Map := make(map[rune]int)

	for _, r := range []rune(s1) {
		s1Map[r]++
	}
	for _, r := range []rune(s2) {
		s2Map[r]++
	}

	var diff int
	for r, n1 := range s1Map {
		n2, ok := s2Map[r]
		if !ok {
			diff += n1
			continue
		}
		diff += int(math.Abs(float64(n1 - n2)))
	}

	return float64(diff)/float64(len(s1)) > 0.1
}
