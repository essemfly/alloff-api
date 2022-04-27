package dto

type DocumentCountDTO struct {
	Aggregations struct {
		GroupByState struct {
			Buckets []struct {
				Key      string `json:"key"`
				DocCount int    `json:"doc_count"`
			} `json:"buckets"`
		} `json:"group_by_state"`
	} `json:"aggregations"`
}
