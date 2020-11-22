package movieservice

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gufranmirza/imdb-api/dal/moviedal"
	"github.com/gufranmirza/imdb-api/database/dbmodels"
	_ "github.com/gufranmirza/imdb-api/database/dbmodels"
	"github.com/gufranmirza/imdb-api/mocks/moviedalmock"
)

func Test_movieservice_Search(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer func() {
		ctrl.Finish()
	}()

	mvdalmock := moviedalmock.NewMockMovieDal(ctrl)

	req, _ := http.NewRequest("GET", "/movies/search", nil)
	w := httptest.NewRecorder()

	req1, _ := http.NewRequest("GET", "/movies/search?q=5&limit=10", nil)
	w1 := httptest.NewRecorder()

	fmt.Println(req1.URL.String())

	type fields struct {
		moviedal moviedal.MovieDal
		logger   *log.Logger
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Query string and limit empty",
			args: args{
				r: req,
				w: w,
			},
			fields: fields{
				logger: log.New(os.Stdout, "movieservice :=> ", log.LstdFlags),
			},
		},
		{
			name: "Database Error",
			args: args{
				r: req1,
				w: w1,
			},
			fields: fields{
				logger: log.New(os.Stdout, "movieservice :=> ", log.LstdFlags),
				moviedal: func() *moviedalmock.MockMovieDal {
					mvdalmock.EXPECT().Search(gomock.Any()).Return(nil, errors.New("database error"))
					return mvdalmock
				}(),
			},
		},
		{
			name: "Search Success",
			args: args{
				r: req1,
				w: w1,
			},
			fields: fields{
				logger: log.New(os.Stdout, "movieservice :=> ", log.LstdFlags),
				moviedal: func() *moviedalmock.MockMovieDal {
					mvdalmock.EXPECT().Search(gomock.Any()).Return([]dbmodels.Movie{
						{
							Name:     "Star Wars",
							Director: "Lucas",
						},
					}, nil)
					return mvdalmock
				}(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &movieservice{
				moviedal: tt.fields.moviedal,
				logger:   tt.fields.logger,
			}
			m.Search(tt.args.w, tt.args.r)
		})
	}
}
