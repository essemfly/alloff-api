package omnious

type EstimateModelType struct {
	Id         string
	Name       string
	Confidence float64
}

type Tags struct {
	Category struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"category"`
	Position struct {
		X float64 `json:"x"`
		Y float64 `json:"y"`
	} `json:"position"`
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
			Type string `json:"type"`
			Tags []Tags `json:"tags"`
		} `json:"objects"`
		NotMatchedObject []struct {
			Type string `json:"type"`
			Tags []Tags `json:"tags"`
		} `json:"notMatchedObject"`
	} `json:"data"`
	Error  interface{} `json:"error"`
	Status string      `json:"status"`
}
