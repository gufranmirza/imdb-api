package movieservice

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gufranmirza/imdb-api/models"
	"github.com/gufranmirza/imdb-api/web/renderers"
)

// @Summary Search Movies
// @Description It allows to search Movies
// @Tags Repository
// @Accept json
// @Produce json
// @Param q query string true "search query"
// @Param limit query string true "response limit default is 10"
// @Success 200 {object} []dbmodels.Movie{}
// @Failure 400 {object} errorinterface.ErrorResponse{}
// @Failure 404 {object} errorinterface.ErrorResponse{}
// @Failure 500 {object} errorinterface.ErrorResponse{}
// @Router /repository/search [GET]
func (m *movieservice) Search(w http.ResponseWriter, r *http.Request) {
	txID := r.Header.Get(models.HdrRequestID)
	query := r.URL.Query().Get("q")
	limitQuery := r.URL.Query().Get("limit")
	limit := 10

	if len(limitQuery) > 0 {
		i, err := strconv.Atoi(limitQuery)
		if err == nil {
			limit = i
		}
	}

	if len(query) == 0 {
		m.logger.Printf("%s %sInvalid query %s", txID, FailedToSearchRepo, query)
		renderers.ErrorInvalidRequest(w, r, errors.New("Query can not be empty"))
		return
	}

	resp, err := m.moviedal.Search(query)
	if err != nil {
		m.logger.Printf("%s %s error %v", txID, FailedToSearchRepo, err)
		renderers.ErrorNotFound(w, r, err)
		return
	}

	if len(resp) > limit {
		resp = resp[:limit]
	}

	renderers.RenderJSON(w, r, txID, resp)
	return
}
