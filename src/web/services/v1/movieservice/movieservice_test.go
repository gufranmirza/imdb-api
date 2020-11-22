package movieservice

import (
	"log"
	"os"
	"testing"

	"github.com/gufranmirza/imdb-api/dal/moviedal"
)

func TestNewMovieServiceService(t *testing.T) {
	tests := []struct {
		name string
		want MovieService
	}{
		{
			name: "Happy Path",
			want: &movieservice{
				logger:   log.New(os.Stdout, "movieservice :=> ", log.LstdFlags),
				moviedal: moviedal.NewMovieDal(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMovieServiceService(); got == nil && tt.want != nil {
				t.Errorf("NewMovieServiceService() = %v, want %v", got, tt.want)
			}
		})
	}
}
