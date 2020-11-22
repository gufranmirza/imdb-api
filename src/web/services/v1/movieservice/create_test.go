package movieservice

import (
	"errors"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gufranmirza/imdb-api/dal/moviedal"
	"github.com/gufranmirza/imdb-api/mocks/moviedalmock"
)

func Test_movieservice_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer func() {
		ctrl.Finish()
	}()

	mvdalmock := moviedalmock.NewMockMovieDal(ctrl)

	body := strings.NewReader("my request")
	req, _ := http.NewRequest("GET", "/movie", body)
	w := httptest.NewRecorder()

	body1 := strings.NewReader(`{
			"99popularity": 88.0,
			"director": "Lucas",
			"genre": [
			  "Sci-Fi"
			],
			"imdb_score": 8.8,
			"name": "Star Wars"
		}`)

	body2 := strings.NewReader(`{
			"99popularity": 88.0,
			"director": "Lucas",
			"genre": [
			  "Sci-Fi"
			],
			"imdb_score": 8.8,
			"name": ""
		}`)

	req1, _ := http.NewRequest("GET", "/movie", body1)
	w1 := httptest.NewRecorder()

	req2, _ := http.NewRequest("GET", "/movie", body1)
	w2 := httptest.NewRecorder()

	req3, _ := http.NewRequest("GET", "/movie", body2)
	w3 := httptest.NewRecorder()

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
			name: "Invalid request body",
			args: args{
				r: req,
				w: w,
			},
		},
		{
			name: "Incomplete request body",
			args: args{
				r: req3,
				w: w3,
			},
		},
		{
			name: "Database failed",
			args: args{
				r: req1,
				w: w1,
			},
			fields: fields{
				logger: log.New(os.Stdout, "movieservice :=> ", log.LstdFlags),
				moviedal: func() *moviedalmock.MockMovieDal {
					mvdalmock.EXPECT().Create(gomock.Any(), gomock.Any()).Return(errors.New("database error"))
					return mvdalmock
				}(),
			},
		},
		{
			name: "Create Success",
			args: args{
				r: req2,
				w: w2,
			},
			fields: fields{
				logger: log.New(os.Stdout, "movieservice :=> ", log.LstdFlags),
				moviedal: func() *moviedalmock.MockMovieDal {
					mvdalmock.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil).AnyTimes()
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
			m.Create(tt.args.w, tt.args.r)
		})
	}
}
