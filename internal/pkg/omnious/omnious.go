package omnious

type OmniousResult struct {
	Category struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"category"`
	domain.TaggingResult
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
