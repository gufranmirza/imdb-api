package renderers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gufranmirza/imdb-api/web/interfaces/v1/errorinterface"
)

// ErrorInvalidRequest returns status 422 Unprocessable Entity including error message.
func ErrorInvalidRequest(w http.ResponseWriter, r *http.Request, err error) {
	resp := &errorinterface.ErrorResponse{
		Err:            err,
		HTTPStatusCode: http.StatusUnprocessableEntity,
		Status:         http.StatusText(http.StatusUnprocessableEntity),
		Error:          err.Error(),
	}
	SendError(w, r, resp)
}

// ErrorUnauthorized renders status 401 Unauthorized with custom error message.
func ErrorUnauthorized(w http.ResponseWriter, r *http.Request, err error) {
	resp := &errorinterface.ErrorResponse{
		Err:            err,
		HTTPStatusCode: http.StatusUnauthorized,
		Status:         http.StatusText(http.StatusUnauthorized),
		Error:          err.Error(),
	}
	SendError(w, r, resp)
}

// ErrorForbidden renders status 403 forbidden with custom error message.
func ErrorForbidden(w http.ResponseWriter, r *http.Request, err error) {
	resp := &errorinterface.ErrorResponse{
		Err:            err,
		Error:          err.Error(),
		HTTPStatusCode: http.StatusForbidden,
		Status:         http.StatusText(http.StatusForbidden),
	}
	SendError(w, r, resp)
}

// ErrorBadRequest return status 400 Bad Request for malformed request body.
func ErrorBadRequest(w http.ResponseWriter, r *http.Request, err error) {
	resp := &errorinterface.ErrorResponse{
		HTTPStatusCode: http.StatusBadRequest,
		Status:         http.StatusText(http.StatusBadRequest),
		Err:            err,
		Error:          err.Error(),
	}
	SendError(w, r, resp)

}

// ErrorNotFound returns status 404 Not Found for invalid resource request.
func ErrorNotFound(w http.ResponseWriter, r *http.Request, err error) {
	resp := &errorinterface.ErrorResponse{
		HTTPStatusCode: http.StatusNotFound,
		Status:         http.StatusText(http.StatusNotFound),
		Err:            err,
		Error:          err.Error(),
	}
	SendError(w, r, resp)

}

// ErrorInternalServerError returns status 500 Internal Server Error.
func ErrorInternalServerError(w http.ResponseWriter, r *http.Request, err error) {
	resp := &errorinterface.ErrorResponse{
		HTTPStatusCode: http.StatusInternalServerError,
		Status:         http.StatusText(http.StatusInternalServerError),
		Err:            err,
		Error:          err.Error(),
	}
	SendError(w, r, resp)
}

// ErrorNotFoundError returns status 404  Error.
func ErrorNotFoundError(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(http.StatusNotFound)
	return
}

// SendError constructs and sends the response
func SendError(w http.ResponseWriter, r *http.Request, data *errorinterface.ErrorResponse) {
	buff, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(data.HTTPStatusCode)
	fmt.Fprintln(w, string(buff))
	return
}
