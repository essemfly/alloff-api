package domain

import (
	"bytes"
	"encoding/json"
	"log"
)

type LogType string

const (
	PRODUCT_VIEW = LogType("PRODUCT_VIEW")
	ORDERED_ITEM = LogType("ORDERED_ITEM")
)

type GraphQLRequest struct {
	OperationName interface{}            `json:"operation_name"`
	Query         interface{}            `json:"query"`
	Variables     map[string]interface{} `json:"variables"`
}

type AccessLogDAO struct {
	Body          GraphQLRequest `json:"body,omitempty"`
	Type          string         `json:"type,omitempty"`
	Timestamp     string         `json:"ts,omitempty"`
	RemoteAddr    string         `json:"remote_addr,omitempty"`
	UserAgent     string         `json:"user_agent,omitempty"`
	Uri           string         `json:"uri,omitempty"`
	RespStatus    int            `json:"resp_status,omitempty"`
	RespElapsedMs float64        `json:"resp_elapsed_ms,omitempty"`
	HttpMethod    string         `json:"http_method"`
	HttpProto     string         `json:"http_proto"`
}

type ProductLogDAO struct {
	Product *ProductDAO `json:"product"`
	Ts      string      `json:"ts"`
	Type    LogType     `json:"type"`
}

type BrandLogDAO struct {
	Brand *BrandDAO `json:"brand"`
	Ts    string    `json:"ts"`
	Type  LogType   `json:"type"`
}

type SearchLogDAO struct {
	Keyword string `json:"keyword"`
	Ts      string `json:"ts"`
}

func (f *AccessLogDAO) JsonEncoder() string {
	bytesBuffer := new(bytes.Buffer)
	res := "{}"
	if f != nil {
		encoder := json.NewEncoder(bytesBuffer)
		encoder.SetEscapeHTML(false)
		encoder.SetIndent("", " ")
		encoder.Encode(f)
		res = bytesBuffer.String()
	} else {
		log.Println("nil pointer of AccessDAO passed")
	}
	return res
}
