package dto

type EstimateModelType struct {
	Id         string
	Name       string
	Confidence float64
}

type OmniousResult struct {
	Category struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"category"`
	TaggingResult
}

type TaggingResult struct {
	Item         EstimateModelType   `json:"item"`
	Colors       []EstimateModelType `json:"colors"`
	ColorDetails []EstimateModelType `json:"colorDetails"`
	Prints       []EstimateModelType `json:"prints"`
	Looks        []EstimateModelType `json:"looks"`
	Textures     []EstimateModelType `json:"textures"`
	Details      []EstimateModelType `json:"details"`
	Length       EstimateModelType   `json:"length"`
	SleeveLength EstimateModelType   `json:"sleeveLength"`
	NeckLine     EstimateModelType   `json:"neckLine"`
	Fit          EstimateModelType   `json:"fit"`
	Shape        EstimateModelType   `json:"shape"`
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
