package websearch

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
)

// ============================================================================
// Search Knowledge
// ============================================================================

func CheckKnowledge(word string) bool {

	dictionaryPath := filepath.Join(
		"knowledge",
		"language",
		"Indonesian",
		"dictionary",
		word+".yaml",
	)

	_, err := os.Stat(dictionaryPath)

	return err == nil
}

// ============================================================================
// Download Wikipedia
// ============================================================================

func SearchWikipedia(word string) ([]byte, error) {

	api := WikipediaAPI + url.PathEscape(word)

	request, err := http.NewRequest(
		http.MethodGet,
		api,
		nil,
	)

	if err != nil {
		return nil, err
	}

	request.Header.Set("User-Agent", UserAgent)

	client := &http.Client{
		Timeout: HTTPTimeout,
	}

	response, err := client.Do(request)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {

		return nil, fmt.Errorf(
			"wikipedia response : %d",
			response.StatusCode,
		)

	}

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	return body, nil

}

// ============================================================================
// Check Internet
// ============================================================================

func CheckInternet() bool {

	request, err := http.NewRequest(
		http.MethodHead,
		"https://id.wikipedia.org",
		nil,
	)

	if err != nil {
		return false
	}

	client := &http.Client{
		Timeout: HTTPTimeout,
	}

	_, err = client.Do(request)

	return err == nil
}
