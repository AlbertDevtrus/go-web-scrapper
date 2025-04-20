package main

import (
	"fmt"
	"net/http"

	"golang.org/x/net/html"
)

const base_url = "https://scrape-me.dreamsofcode.io"

func main() {

	response, err := Get(base_url)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	urlList := GetUrlList(response)

	fmt.Print(urlList)

}

func Get(url string) (response *http.Response, err error) {
	fmt.Print(url)

	response, err = http.Get(url)

	if err != nil {
		return nil, err
	}

	return response, nil
}

func GetUrlList(response *http.Response) (urlList []string) {

	tokenizer := html.NewTokenizer(response.Body)

	for {
		tokenType := tokenizer.Next()
		if tokenType == html.ErrorToken {
			break
		}

		token := tokenizer.Token()

		if tokenType == html.StartTagToken && token.Data == "a" {

			for _, attr := range token.Attr {
				if attr.Key != "href" {
					continue
				}

				url := attr.Val

				if url[0] == '/' {
					url = base_url + attr.Val
				}

				urlList = append(urlList, url)
			}

		}

	}

	return urlList
}
