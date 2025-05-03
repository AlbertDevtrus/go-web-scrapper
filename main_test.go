package main

import (
	"testing"

	"github.com/AlbertDevtrus/go-web-scrapper/set"
)

// type MockChromedp struct {
// 	mock.Mock
// }

// func (m *MockChromedp) Run(ctx context.Context, tasks ...chromedp.Action) error {
// 	args := m.Called(ctx, tasks)
// 	return args.Error(0)
// }

// func TestGetHTMLMock(t *testing.T) {
// 	ctx, cancel := chromedp.NewContext(context.Background())
// 	defer cancel()

// 	mockChromedp := new(MockChromedp)

// 	mockChromedp.On("Run", ctx, mock.Anything).Return(nil)

// 	tokenizer, err := GetHTML(ctx, "https://example.com")

// 	if err != nil {
// 		t.Fatalf("GetHTML() error = %v", err)
// 	}

// 	if tokenizer == nil {
// 		t.Fatal("GetHTML() tokenizer is nil")
// 	}

// 	mockChromedp.AssertExpectations(t)
// }

func TestGetHostUrl(t *testing.T) {
	tests := []struct {
		url      string
		expected string
	}{
		{"https://scrape-me.dreamsofcode.io", "https://scrape-me.dreamsofcode.io"},
		{"https://example.com/path/to/resource", "https://example.com"},
		{"http://localhost:8080", "http://localhost:8080"},
		{"ftp://example.com/file.txt", "ftp://example.com"},
	}

	for _, test := range tests {
		result, _ := GetHostUrl(test.url)
		if result != test.expected {
			t.Errorf("GetHostUrl(%s) = %s, expected = %s", test.url, result, test.expected)
		}
	}
}

func TestIsValidURL(t *testing.T) {
	visited := set.NewSet()

	tests := []struct {
		link     string
		expected bool
	}{
		{"https://valid.url", true},
		{"mailto:test@example.com", false},
		{"tel:+123456789", false},
		{"", false},
		{"https://scrape-me.dreamsofcode.io", true},
		{"#", false},
	}

	for _, test := range tests {
		result := isValidURL(visited, test.link)
		if result != test.expected {
			t.Errorf("isValidURL(%s) = %v, expected = %v", test.link, result, test.expected)
		}
	}
}

func TestGetStatusCode(t *testing.T) {
	tests := []struct {
		url      string
		expected int
	}{
		{"https://scrape-me.dreamsofcode.io", 200},
		{"https://scrape-me.dreamsofcode.io/404", 404},
	}

	for _, test := range tests {
		statusCode, err := GetStatusCode(test.url)
		if err != nil {
			t.Errorf("GetStatusCode(%s) error = %v", test.url, err)
		}
		if statusCode != test.expected {
			t.Errorf("GetStatusCode(%s) = %d, expected = %d", test.url, statusCode, test.expected)
		}
	}
}
