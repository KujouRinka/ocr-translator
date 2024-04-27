package config

import (
	"os"

	"golang.org/x/text/language"
	"gopkg.in/yaml.v3"

	"github.com/kujourinka/ocr-translator/ocr"
	"github.com/kujourinka/ocr-translator/translator"
)

type RawConfig struct {
	OCR         OCRConfig          `yaml:"ocr"`
	Translators []TranslatorConfig `yaml:"translators"`
}

type OCRConfig struct {
	Type string `yaml:"type"`
}

type TranslatorConfig struct {
	Type       string   `yaml:"type"`
	APIKey     string   `yaml:"api"`
	TargetLang string   `yaml:"target"`
	SourceLang []string `yaml:"source"`
}

type Config struct {
	OCR         ocr.Engine
	Translators []translator.Engine
}

func LoadRawConfig(path string) (*RawConfig, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var rawConfig RawConfig
	return &rawConfig, yaml.NewDecoder(file).Decode(&rawConfig)
}

func StrToLang(str string) language.Tag {
	if lang, ok := strToLangMap[str]; ok {
		return lang
	}
	return language.English
}

var (
	strToLangMap = map[string]language.Tag{
		"en":     language.English,
		"ja":     language.Japanese,
		"zh-chs": language.SimplifiedChinese,
		"zh-cht": language.TraditionalChinese,
	}
)
