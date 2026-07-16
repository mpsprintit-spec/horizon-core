package wikipedia

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type WikiResponse struct {
	Title   string `json:"title"`
	Extract string `json:"extract"`
}

func SearchWikipedia(keyword string) (*WikiResponse, error) {

	endpoint := WikipediaAPI + url.PathEscape(keyword)

	resp, err := http.Get(endpoint)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var result WikiResponse

	err = json.Unmarshal(body, &result)

	if err != nil {
		return nil, err
	}

	if result.Extract == "" {
		return nil, fmt.Errorf("artikel tidak ditemukan")
	}

	return &result, nil
}
