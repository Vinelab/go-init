package throw

import (
	"net/http"

	"github.com/vinelab/go-init/response"
)

func HandleHttpRequestErrors(w http.ResponseWriter) {

	//check if panic happened
	if err := recover(); err != nil {
		var status uint16
		var code uint16

		switch err.(type) {
		case InvalidInput:
			status = 400

			err := err.(InvalidInput)
			code = err.Code
		default:
			status = 500
			code = 500
		}

		response.JSONError(w, status, code, err.(error).Error())
	}
}
