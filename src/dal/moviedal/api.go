package moviedal

import "github.com/gufranmirza/imdb-api/database/dbmodels"

// MovieDal interface
type MovieDal interface {
	Create(txID string, movie *dbmodels.Movie) error
	Search(query string) ([]dbmodels.Movie, error)
}
