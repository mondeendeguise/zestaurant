package mock

import (
	"bytes"
	"embed"
	"io"
	"net/http"
)

//go:embed json
var jsonFS embed.FS

type ResponseType int

const (
	Metadata ResponseType = iota
	Menus
)

var responseFiles = map[ResponseType]string{
	Metadata: "json/metadata.json",
	Menus: "json/menus.json",
}

func MockResponse(responseType ResponseType, status string, statusCode int) (response *http.Response, err error) {
	jsonFile, err := jsonFS.ReadFile(responseFiles[responseType])
	var buf *bytes.Buffer = bytes.NewBuffer(jsonFile)

	response = &http.Response {
		Status: status,
		StatusCode: statusCode,
		Proto: "HTTP/1.0",
		ProtoMajor: 1,
		ProtoMinor: 0,

		Header: http.Header {},
		Body: io.NopCloser(buf),
		ContentLength: int64(len(jsonFile)),
		TransferEncoding: nil,
		Close: false,
		Uncompressed: true,
		Trailer: nil,
		Request: nil,
		TLS: nil,
	}

	return response, err
}
