package helper

import (
	"encoding/json"
	"net/http"
)

func RaedFromRequestBody(request *http.Request, result interface{}) {
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(result)
	PanicIfError(err)

}

func WriteToResponseBody(writer http.ResponseWriter, response interface{}) {
	writer.Header().Add("Content-type", "application/json")
	encode := json.NewEncoder(writer)
	err := encode.Encode(response)
	PanicIfError(err)
}
