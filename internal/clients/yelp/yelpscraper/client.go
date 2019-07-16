package yelpscraper

import (
	"net/http"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

// ScrapeForYelpBizID parses a yelp business page for its business ID
func ScrapeForYelpBizID(yelpURL string) (bizID string, err error) {
	resp, err := http.Get(yelpURL)
	if err != nil {
		return "", err
	}
	bizID = findYelpBizID(html.NewTokenizer(resp.Body))
	err = resp.Body.Close()
	return bizID, err
}

func findYelpBizID(n *html.Tokenizer) string {
	for {
		tt := n.Next()
		switch {
		case tt == html.ErrorToken:
			return ""
		case tt == html.StartTagToken:
			bizID := getYelpBizID(n.Token())
			if bizID != "" {
				return bizID
			}
		}
	}
}

func getYelpBizID(n html.Token) string {
	if n.DataAtom == atom.Meta {
		var bizID string
		var elementIsBizID bool
		for _, attribute := range n.Attr {
			if attribute.Key == "content" {
				bizID = attribute.Val
			}
			if attribute.Key == "name" && attribute.Val == "yelp-biz-id" {
				elementIsBizID = true
			}
		}
		if elementIsBizID {
			return bizID
		}
	}
	return ""
}
