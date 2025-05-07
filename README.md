# Go Web Scrapper

A high-performance web scraper that uses goroutines and chromedp to recursively extract URLs with concurrent processing.

## Features

- Recursive URL scraping from any website

- Goroutine-powered concurrency for blazing fast scraping

- Error tracking with dedicated error URL listing

- Randomized delays between requests to be server-friendly

## Installation

1. Clone repository

``` git clone https://github.com/your-username/go-web-scraper.git ```

2. Install dependencies

``` go mod download ```

## Usage

Run the scrapper  ` go run main.go ` and then enter the target url 

## Performance 

Memory-efficient (each goroutine uses ~2KB)

## Additional information

This is an educational project, there are still some things to add, like better error handling, testing to main functions and adding workers to the goroutines, I also want to add a return file and some modifications to the user experience!