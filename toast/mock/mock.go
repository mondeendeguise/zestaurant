package mock

import (
	"bytes"
	"embed"
	"io"
	"net/http"
)

//go:embed json
var jsonFS embed.FS

const METADATA_FILE string = "json/metadata.json"
const MENUS_FILE string = "json/menus.json"

func MockResponse(filename string, status string, statusCode int) (response *http.Response, err error) {
	jsonFile, err := jsonFS.ReadFile(filename)

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
