package translator

// Use google cloud translate API v2

import (
	"context"
	"fmt"
	"net/http"

	"cloud.google.com/go/translate"
	"golang.org/x/net/proxy"
	"golang.org/x/text/language"
	"google.golang.org/api/googleapi/transport"
	"google.golang.org/api/option"
)

var _ Engine = (*GoogleTranslator)(nil)

type GoogleTranslator struct {
	client *translate.Client
	apiKey string
	ctx    context.Context

	targetLang language.Tag
	sourceLang language.Tag
	proxy      string
}

func NewGoogleTranslator(target, source language.Tag, apiKey, p string) (*GoogleTranslator, error) {
	ctx := context.Background()
	opt := make([]option.ClientOption, 0)

	if p != "" {
		socksProxyAddr := p

		dialer, err := proxy.SOCKS5("tcp", socksProxyAddr, nil, proxy.Direct)
		if err != nil {
			return nil, err
		}

		httpClient := &http.Client{
			Transport: &transport.APIKey{
				Key: apiKey,
				Transport: &http.Transport{
					Dial: dialer.Dial,
				},
			},
		}

		opt = append(opt, option.WithHTTPClient(httpClient))
	} else {
		opt = append(opt,
			option.WithAPIKey(apiKey),
			// option.WithoutAuthentication(),
		)
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
