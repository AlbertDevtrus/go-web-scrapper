package main

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/AlbertDevtrus/go-web-scrapper/set"
	"github.com/chromedp/chromedp"
	"golang.org/x/net/html"
)

/*
	TODO:  Ideas

* Guardar los resultados en un archivo (json, csv, etc.)
* Añadir tests unitarios
* Hacerlo concurrente con goroutines y canales
*/
var base_url = "https://scrape-me.dreamsofcode.io"

func main() {

	fmt.Println("Welcome to the web crawler!")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("This is a simple web crawler that crawls a website and prints the URLs found recursively.")
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("")
	fmt.Println("Please, enter the URL to crawl.")

	var user_url string
	fmt.Scanln(&user_url)

	base_url = ParseURL(user_url)

	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	visited := set.NewSet()
	errorUrls := set.NewSet()

	visited.Add(base_url)

	fmt.Println("Crawling:", base_url)
	urlList := GetUrlList(ctx, base_url)

	CrawlList(ctx, urlList, visited, errorUrls)

	PrintErrorUrls(errorUrls.List())
}

func GetHTML(ctx context.Context, url string) (tokenizer *html.Tokenizer, err error) {
	var htmlContent string
	err = chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.WaitReady("body", chromedp.ByQuery),
		chromedp.OuterHTML("html", &htmlContent, chromedp.ByQuery),
	)

	if err != nil {
		return nil, err
	}

	reader := strings.NewReader(htmlContent)
	tokenizer = html.NewTokenizer(reader)

	return tokenizer, nil
}

func GetUrlList(ctx context.Context, url string) (urlList []string) {

	tokenizer, err := GetHTML(ctx, url)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	hostUrl, err := GetHostUrl(url)

	if err != nil {
		fmt.Println("Error:", err)
		return
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

				if len(url) > 0 && url[0] == '/' {
					url = hostUrl + attr.Val
				}

				urlList = append(urlList, url)
			}

		}

	}

	return urlList
}

func CrawlList(ctx context.Context, urlList []string, visited *set.Set, errorUrls *set.Set) {
	for _, currentUrl := range urlList {
		if !isValidURL(visited, currentUrl) {
			continue
		}

		fmt.Println("Crawling:", currentUrl)

		seconds := rand.Intn(3) + 1
		time.Sleep(time.Duration(seconds) * time.Second)

		visited.Add(currentUrl)
		statusCode, err := GetStatusCode(currentUrl)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		if statusCode >= 400 {
			errorUrls.Add(currentUrl)
		}

		hostUrl, err := GetHostUrl(currentUrl)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		if hostUrl == base_url {
			newUrlList := GetUrlList(ctx, currentUrl)
			CrawlList(ctx, newUrlList, visited, errorUrls)
		}
	}
}

func PrintErrorUrls(urlList []string) {
	fmt.Print("\n=======================\n")
	fmt.Print("\n----- ERROR URLS ------\n")
	fmt.Print("\n=======================\n")

	for i := range urlList {
		fmt.Printf("\033[31m%s\033[0m \n", urlList[i])
	}
}

func GetStatusCode(url string) (statusCode int, err error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return 0, err
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) "+
		"AppleWebKit/537.36 (KHTML, like Gecko) "+
		"Chrome/113.0.0.0 Safari/537.36")

	response, err := client.Do(req)

	if err != nil {
		return 0, err
	}
	defer response.Body.Close()

	statusCode = response.StatusCode

	return statusCode, nil
}

func GetHostUrl(rawUrl string) (hostUrl string, err error) {
	parsedUrl, err := url.Parse(rawUrl)
	if err != nil {
		return "", err
	}

	hostUrl = fmt.Sprintf("%s://%s", parsedUrl.Scheme, parsedUrl.Host)
	return hostUrl, nil
}

func isValidURL(visited *set.Set, link string) bool {
	if visited.Has(link) {
		return false
	}

	if link == "" || strings.HasPrefix(link, "#") || strings.HasPrefix(link, ".") {
		return false
	}

	if strings.HasPrefix(link, "mailto:") || strings.HasPrefix(link, "tel:") {
		return false
	}

	_, err := url.ParseRequestURI(link)

	return err == nil
}

func ParseURL(url string) string {
	if strings.HasSuffix(url, "/") {
		return strings.TrimSuffix(url, "/")
	}
	return url
}
