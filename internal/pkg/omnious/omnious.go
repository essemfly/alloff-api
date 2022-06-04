package omnious

import "github.com/lessbutter/alloff-api/internal/core/domain"

type OmniousResult struct {
	Category struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"category"`
	domain.TaggingResultDAO
}

type PostResponse struct {
	Data struct {
		Objects []struct {
			Type string          `json:"type"`
			Tags []OmniousResult `json:"tags"`
		} `json:"objects"`
		NotMatchedObject []struct {
			Type string          `json:"type"`
			Tags []OmniousResult `json:"tags"`
		} `json:"notMatchedObject"`
	} `json:"data"`
	Error  interface{} `json:"error"`
	Status string      `json:"status"`
}
