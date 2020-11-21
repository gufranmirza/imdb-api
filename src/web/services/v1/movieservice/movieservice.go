package movieservice

import (
	"log"
	"os"

	"github.com/gufranmirza/imdb-api/dal/moviedal"
)

type movieservice struct {
	moviedal moviedal.MovieDal
	logger   *log.Logger
}

// NewMovieServiceService returns service impl
func NewMovieServiceService() MovieService {
	return &movieservice{
		moviedal: moviedal.NewMovieDal(),
		logger:   log.New(os.Stdout, "movieservice :=> ", log.LstdFlags),
	}
}
