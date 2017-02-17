package response

import "net/http"
import "encoding/json"

// Success is the response structure for a successful request
type Success struct {
	Status uint16      `json:"status"`
	Data   interface{} `json:"data"`
}

// Failure is the response structure for a failed request
type Failure struct {
	Status uint16 `json:"status"`
	Error  `json:"error"`
}

// Error is the error structure
type Error struct {
	Code    uint16 `json:"code"`
	Message string `json:"message"`
}

// JSON responds with a pre-defined json structure for success
func JSON(w http.ResponseWriter, response interface{}) {
	w.Header().Set("Content-Type", "application/json")

	s := &Success{
		Status: 200,
		Data:   response,
	}

	j, err := json.Marshal(s)
	if err != nil {
		panic(err)
	}

	w.Write(j)
}

// JSONError responds with a pre-defined json structure for errors.
func JSONError(w http.ResponseWriter, s uint16, c uint16, m string) {
	w.Header().Set("Content-Type", "application/json")

	e := &Failure{
		Status: s,
		Error: Error{
			Code:    c,
			Message: m,
		},
	}

	j, err := json.Marshal(e)
	if err != nil {
		panic(err)
	}

	switch {
	case s >= 400 && s < 500:
		w.WriteHeader(http.StatusBadRequest)
	case s >= 500 && s < 600:
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Write(j)
}
