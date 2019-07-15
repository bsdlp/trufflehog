package yelpscraper

import (
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func TestGetYelpBizID(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		node := html.Token{
			DataAtom: atom.Meta,
			Data:     "meta",
			Attr: []html.Attribute{
				{
					Key: "name",
					Val: "yelp-biz-id",
				},
				{
					Key: "content",
					Val: "4yPqqJDJOQX69gC66YUDkA",
				},
			},
		}

		assert.Equal(t, "4yPqqJDJOQX69gC66YUDkA", getYelpBizID(node))
	})

	t.Run("biz id but no id", func(t *testing.T) {
		node := html.Token{
			DataAtom: atom.Meta,
			Data:     "meta",
			Attr: []html.Attribute{
				{
					Key: "name",
					Val: "yelp-biz-id",
				},
			},
		}

		assert.Empty(t, getYelpBizID(node))
	})

	t.Run("no biz id", func(t *testing.T) {
		node := html.Token{
			DataAtom: atom.Meta,
			Data:     "meta",
			Attr: []html.Attribute{
				{
					Key: "name",
					Val: "ome-othr-id",
				},
				{
					Key: "content",
					Val: "4yPqqJDJOQX69gC66YUDkA",
				},
			},
		}

		assert.Empty(t, getYelpBizID(node))
	})
}

func TestFindYelpBizID(t *testing.T) {
	assert.Equal(t, "4yPqqJDJOQX69gC66YUDkA", findYelpBizID(html.NewTokenizer(strings.NewReader(peterLugerYelpPage))))
}

func TestIntegrationFindYelpBizID(t *testing.T) {
	resp, err := http.Get("https://www.yelp.com/biz/peter-luger-brooklyn-2")
	require.NoError(t, err)

	defer resp.Body.Close()

	assert.Equal(t, "4yPqqJDJOQX69gC66YUDkA", findYelpBizID(html.NewTokenizer(resp.Body)))
}