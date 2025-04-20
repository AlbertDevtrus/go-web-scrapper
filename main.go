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
	errorUrls := set.NewSet()

	response, err := Get(base_url, errorUrls)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	visited.Add(base_url)

	urlList := GetUrlList(response)
	response.Body.Close()

	CrawlList(urlList, visited, errorUrls)

	PrintErrorUrls(errorUrls.List())
}

func Get(url string, errorUrls *set.Set) (response *http.Response, err error) {
	fmt.Printf("%s\n", url)

	response, err = http.Get(url)

	if err != nil {
		return nil, err
	}

	if response.StatusCode >= 400 {
		errorUrls.Add(url)
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

func CrawlList(urlList []string, visited *set.Set, errorUrls *set.Set) {
	for i := 0; i < len(urlList); i++ {

		if visited.Has(urlList[i]) {
			continue
		}

		visited.Add(urlList[i])

		response, err := Get(urlList[i], errorUrls)

		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		newUrlList := GetUrlList(response)
		response.Body.Close()

		CrawlList(newUrlList, visited, errorUrls)
	}
}

func PrintErrorUrls(urlList []string) {
	fmt.Print("\n=======================\n")
	fmt.Print("\n ERROR URLS \n")
	fmt.Print("\n=======================\n")

	for i := 0; i < len(urlList); i++ {
		fmt.Printf("\033[31m%s\033[0m \n", urlList[i])
	}
}
