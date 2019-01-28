package main

import (
	"fmt"
	"io"
	"log"
	myutils "nerdcast-downloader/my-utils"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	resp, err := http.Get("https://jovemnerd.com.br/feed-nerdcast/")
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range getLinks(resp.Body) {
		fmt.Println(v)
	}
}

//Collect all links from response body and return it as an array of strings
func getLinks(body io.Reader) []string {
	var links []string
	z := html.NewTokenizer(body)
	for {
		tt := z.Next()

		switch tt {
		case html.ErrorToken:
			//todo: links list shoudn't contain duplicates
			return links
		case html.SelfClosingTagToken:
			token := z.Token()
			if "enclosure" == token.Data {
				for _, attr := range token.Attr {
					if attr.Key == "url" {
						if strings.HasPrefix(attr.Val, "https://nerdcast.jovemnerd.com.br/nerd") {
							if !myutils.Contains(links, attr.Val) {
								links = append(links, attr.Val)
							}
						}
					}

				}
			}

		}
	}
}
