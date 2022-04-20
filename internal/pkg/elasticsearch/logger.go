package elasticsearch

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/middleware"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

type LogFields struct {
	Body          GqlReq  `json:"body,omitempty"`
	Type          string  `json:"type,omitempty"`
	Timestamp     string  `json:"ts,omitempty"`
	RemoteAddr    string  `json:"remote_addr,omitempty"`
	UserAgent     string  `json:"user_agent,omitempty"`
	Uri           string  `json:"uri,omitempty"`
	RespStatus    int     `json:"resp_status,omitempty"`
	RespElapsedMs float64 `json:"resp_elapsed_ms,omitempty"`
	HttpMethod    string  `json:"http_method"`
	HttpProto     string  `json:"http_proto"`
}

type GqlReq struct {
	OperationName interface{}            `json:"operation_name"`
	Query         interface{}            `json:"query"`
	Variables     map[string]interface{} `json:"variables"`
}

// StructuredLogger : RequestLogger 미들웨어를 등록하기 위한 인자
type StructuredLogger struct {
	Logger *log.Logger
	Fields LogFields
}

// NewLogEntry : LogFormatter 인터페이스를 구현하는 메서드로, 커스터마이징한 StructuredLogger 를 이용해 리퀘스트에 대한 로깅을 할 수 있도록 한다.
func (l *StructuredLogger) NewLogEntry(r *http.Request) middleware.LogEntry {
	// request body를 읽어서 body 변수에 unmarshal 한다. body 변수는 map[string]interface{] 타입이다.
	var body map[string]interface{}
	b, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(b, &body)
	// 이미 읽은 r.Body에 다시 body 넣어줌
	r.Body = io.NopCloser(bytes.NewReader(b))

	query, ok := body["query"].(string)
	if !ok {
		query = ""
	}

	variables, ok := body["variables"].(map[string]interface{})
	if !ok {
		variables = nil
	}

	// gql 서버에 오는 request body를 저장
	queryPrettier := strings.Replace(query, "\n", "", -1)
	gqlBody := GqlReq{
		OperationName: body["operationName"],
		Query:         queryPrettier,
		Variables:     variables,
	}
	l.Fields.Body = gqlBody

	// request가 서버에 전송된 시간을 / request의 원격주소 / request의 user_agent 를 기록
	l.Fields.Timestamp = time.Now().Format("2006-01-02 15:04:05")
	l.Fields.RemoteAddr = r.RemoteAddr
	l.Fields.UserAgent = r.UserAgent()
	l.Fields.HttpProto = r.Proto
	l.Fields.HttpMethod = r.Method

	// request의 타겟 uri를 기록
	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}
	l.Fields.Uri = fmt.Sprintf("%s://%s%s", scheme, r.Host, r.RequestURI)
	return l
}

// Write : LogEntry 인터페이스를 구현하는 메서드로, 리퀘스트가 종료될때 로그를 받아오는 역할을 함.
// ElasticSearch에 기록하기위해 리퀘스트 종료 후 응답에 대한 내용을 기록하고 이를 별개의 고루틴에서 ElasticSearch에 인덱싱한다.
func (l *StructuredLogger) Write(status, bytes int, header http.Header, elapsed time.Duration, extra interface{}) {
	l.Fields.RespStatus = status
	l.Fields.Type = "request completed"
	l.Fields.RespElapsedMs = float64(elapsed.Nanoseconds()) / 1000000.0
	go func() {
		_, err := logRequest(&l.Fields)
		if err != nil {
			log.Println("Error on request to elastic search : ", err)
		}
	}()
}

// Panic : LogEntry 인터페이스를 구현하는 메서드로, 패닉이 일어날 때 로그를 남김
func (l *StructuredLogger) Panic(v interface{}, stack []byte) {
	l.Logger.Printf(`
		"stack": "%s",
		"panic": "%+v",
	`, stack, v)
}

func (f *LogFields) jsonEncoder() string {
	bytesBuffer := new(bytes.Buffer)
	encoder := json.NewEncoder(bytesBuffer)
	encoder.SetEscapeHTML(false)
	encoder.SetIndent("", " ")
	encoder.Encode(f)
	res := bytesBuffer.String()
	return res
}
