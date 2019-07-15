package yelpscraper

import (
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

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
	return ""
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
