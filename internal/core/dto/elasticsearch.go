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

type AccessLogDTO struct {
	Hits struct {
		Hits []struct {
			Index  string `json:"_index"`
			ID     string `json:"_id"`
			Source struct {
				Body struct {
					OperationName interface{} `json:"operation_name"`
					Query         string      `json:"query"`
					Variables     interface{} `json:"variables"`
				} `json:"body"`
				Type          string  `json:"type"`
				Ts            string  `json:"ts"`
				RemoteAddr    string  `json:"remote_addr"`
				UserAgent     string  `json:"user_agent"`
				URI           string  `json:"uri"`
				RespStatus    int     `json:"resp_status"`
				RespElapsedMs float64 `json:"resp_elapsed_ms"`
				HTTPMethod    string  `json:"http_method"`
				HTTPProto     string  `json:"http_proto"`
			} `json:"_source"`
		} `json:"hits"`
	} `json:"hits"`
}
