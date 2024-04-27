package translator

// Use google cloud translate API v2

import (
	"context"
	"fmt"

	"cloud.google.com/go/translate"
	"golang.org/x/text/language"
	"google.golang.org/api/option"
)

type GoogleTranslator struct {
	client *translate.Client
	apiKey string

	targetLang language.Tag
	sourceLang language.Tag
	ctx        context.Context
}

func NewGoogleTranslator(target, source language.Tag, apiKey string) (*GoogleTranslator, error) {
	ctx := context.Background()
	opt := []option.ClientOption{
		option.WithAPIKey(apiKey),
		// option.WithoutAuthentication(),
	}

	client, err := translate.NewClient(ctx, opt...)
	if err != nil {
		return nil, err
	}

	return &GoogleTranslator{
		client:     client,
		apiKey:     apiKey,
		targetLang: target,
		sourceLang: source,
		ctx:        ctx,
	}, nil
}

func (g *GoogleTranslator) Translate(text string) (string, error) {
	resp, err := g.client.Translate(g.ctx, []string{text}, g.targetLang, &translate.Options{
		Source: g.sourceLang,
		Format: translate.Text,
	})
	if err != nil {
		return "", err
	}
	if len(resp) == 0 {
		return "", fmt.Errorf("translate returned empty response to text: %s", text)
	}
	return resp[0].Text, nil
}

func (g *GoogleTranslator) Close() error {
	return g.client.Close()
}
