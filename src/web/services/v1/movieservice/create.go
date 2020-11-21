package movieservice

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/render"
	"github.com/gufranmirza/imdb-api/database/dbmodels"
	"github.com/gufranmirza/imdb-api/models"
	"github.com/gufranmirza/imdb-api/models/authmodel"
	"github.com/gufranmirza/imdb-api/web/renderers"
)

// @Summary Create Movie
// @Description It allows to Create a new movie details
// @Tags movie
// @Accept json
// @Produce json
// @Param * body &dbmodels.Movie{} true "create movie"
// @Success 200
// @Failure 400 {object} errorinterface.ErrorResponse{}
// @Failure 404 {object} errorinterface.ErrorResponse{}
// @Failure 500 {object} errorinterface.ErrorResponse{}
// @Router /movie [POST]
func (m *movieservice) Create(w http.ResponseWriter, r *http.Request) {
	txID := r.Header.Get(models.HdrRequestID)

	data := &dbmodels.Movie{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		renderers.ErrorBadRequest(w, r, authmodel.ErrIncorrectDetails)
		return
	}

	if len(data.Name) == 0 || len(data.Director) == 0 {
		renderers.ErrorBadRequest(w, r, authmodel.ErrIncorrectDetails)
		return
	}

	err = m.moviedal.Create(txID, data)
	if err != nil {
		m.logger.Printf("%s %s Failed to create movie record with error %v", txID, authmodel.FailedToSignUp, err)
		renderers.ErrorInternalServerError(w, r, authmodel.ErrServerError)
		return
	}

	render.Respond(w, r, http.NoBody)
	return
}
