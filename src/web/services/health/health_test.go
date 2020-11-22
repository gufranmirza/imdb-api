package health

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gufranmirza/imdb-api/config"
	"github.com/gufranmirza/imdb-api/database/connection"
	"github.com/gufranmirza/imdb-api/models"
)

func TestNewHealth(t *testing.T) {
	config.Config = &models.AppConfig{
		JWT: &models.JWT{
			Secret: "1456465",
		},
	}

	tests := []struct {
		name string
		want Health
	}{
		{
			name: "Happy Path",
			want: &health{
				logger: log.New(os.Stdout, "health :=> ", log.LstdFlags),
				config: config.Config,
				db:     connection.NewMongoStore(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHealth(); got == nil && tt.want != nil {
				t.Errorf("NewHealth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_health_GetHealth(t *testing.T) {
	config.Config = &models.AppConfig{
		JWT: &models.JWT{
			Secret: "1456465",
		},
		Database: &models.DB{
			Host:     "mongodb://mongo:27017",
			Database: "imdb-api",
		},
	}
	health := NewHealth()
	req, err := http.NewRequest("GET", "/health/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(health.GetHealth)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}
