package main

import (
	"log"
	"net/http"

	"golang.org/x/net/html"
)

type Dashboard struct {
	URL string
}

func NewDashboard(url string) Dashboard {
	return Dashboard{
		URL: url,
	}
}
func (d Dashboard) Open() (*html.Node, error) {
	resp, err := http.Get(d.URL)
	if err != nil {
		return &html.Node{}, err
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Printf("unable to close response body: %v", err)
		}
	}()

	node, err := html.Parse(resp.Body)
	return node, err
}
