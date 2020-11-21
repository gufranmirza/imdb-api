package movieservice

import (
	"errors"
	"net/http"
)

// MovieService interface ...
type MovieService interface {
	Search(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
}

// The list of error types presented to the end user as error message.
var (
	ErrIncompleteDetails = errors.New("Incorrect details provided, please provice correct details")
	FailedToSearchRepo   = "Failed-To-Search-Repos"
)
