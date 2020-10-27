package helpers

import "encoding/json"

type Stdout struct {
	Request    interface{} `json:"request"`
	Header     interface{} `json:"header"`
	StatusCode interface{} `json:"status_code"`
	Response   interface{} `json:"response"`
}

type StatusCode struct {
	StatusCode interface{} `json:"status_code"`
}

type Header struct {
	ContentType interface{} `json:"Content-Type"`
}

func GenerateStdout(request, header, statusCode, response interface{}) string {
	statusCodeStruct := &StatusCode{
		StatusCode: statusCode,
	}
	headerStruct := &Header{
		ContentType: header,
	}

	requestByte, _ := json.Marshal(request)
	headerByte, _ := json.Marshal(headerStruct)
	statusCodeByte, _ := json.Marshal(statusCodeStruct)
	responseByte, _ := json.Marshal(response)

	return string(requestByte) + " | " + string(headerByte) + " | " + string(statusCodeByte) + " | " + string(responseByte)
}
