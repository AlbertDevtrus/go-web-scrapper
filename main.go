package main

import (
	"fmt"
	"net/http"

	"github.com/AlbertDevtrus/go-web-scrapper/set"
	"golang.org/x/net/html"
)

const base_url = "https://scrape-me.dreamsofcode.io"

func main() {

	visited := set.NewSet()

	response, err := Get(base_url)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	visited.Add(base_url)

	urlList := GetUrlList(response)

	CrawlList(urlList, visited)
}

func Get(url string) (response *http.Response, err error) {
	fmt.Printf("%s\n", url)

	response, err = http.Get(url)

	if err != nil {
		return nil, err
	}

	return response, nil
}

func GetUrlList(response *http.Response) (urlList []string) {

	tokenizer := html.NewTokenizer(response.Body)

	hostUrl := fmt.Sprintf("%s://%s", response.Request.URL.Scheme, response.Request.URL.Host)

	if hostUrl != base_url {
		return urlList
	}

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

func CrawlList(urlList []string, visited *set.Set) {
	for i := 0; i < len(urlList); i++ {

		if visited.Has(urlList[i]) {
			continue
		}

		visited.Add(urlList[i])

		response, err := Get(urlList[i])

		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		newUrlList := GetUrlList(response)

		CrawlList(newUrlList, visited)
	}
}
