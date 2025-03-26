package feeds

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
	"time"
)

type FeedItem struct {
	ID        string
	Title     string
	Link      string
	Timestamp time.Time
	ImageData []byte
}

type Feed interface {
	Name() string
	FetchItems() ([]FeedItem, error)
}

func fetchImageBytes(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 || !strings.HasPrefix(resp.Header.Get("Content-Type"), "image/") {
		return nil, fmt.Errorf("not an image or bad status: %s", resp.Status)
	}

	return io.ReadAll(resp.Body)
}

var (
	markdownImageRegex = regexp.MustCompile(`!\[[^\]]*\]\(([^)]+)\)`)
	htmlImageRegex     = regexp.MustCompile(`<img[^>]+src="([^">]+)"`)
)

func extractFirstImageURL(content string) string {
	if match := markdownImageRegex.FindStringSubmatch(content); len(match) > 1 {
		return match[1]
	}
	if match := htmlImageRegex.FindStringSubmatch(content); len(match) > 1 {
		return match[1]
	}
	return ""
}
