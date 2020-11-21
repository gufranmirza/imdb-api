package health

import (
	"log"
	"net/http"
	"os"
	"runtime"
	"time"

	"github.com/gufranmirza/imdb-api/config"
	"github.com/gufranmirza/imdb-api/database/connection"
	"github.com/gufranmirza/imdb-api/models"
	"github.com/gufranmirza/imdb-api/web/interfaces/v1/healthinterface"
	"github.com/gufranmirza/imdb-api/web/renderers"
)

type health struct {
	logger *log.Logger
	config *models.AppConfig
	db     connection.MongoStore
}

// NewHealth returns new health object
func NewHealth() Health {
	return &health{
		logger: log.New(os.Stdout, "health :=> ", log.LstdFlags),
		config: config.Config,
		db:     connection.NewMongoStore(),
	}
}

// @Summary Get health of the service
// @Description It returns the health of the service
// @Tags health
// @Accept  json
// @Produce  json
// @Success 200 {object} healthinterface.Health{}
// @Failure 400 {object} errorinterface.ErrorResponse{}
// @Failure 404 {object} errorinterface.ErrorResponse{}
// @Failure 500 {object} errorinterface.ErrorResponse{}
// @Router /health [GET]
// GetHealth returns heath of service, can be extended if
// service is running on multile instances
// GetHealth returns heath of service
func (h *health) GetHealth(w http.ResponseWriter, r *http.Request) {

	txID, _ := r.Context().Value(models.HdrRequestID).(string)

	healthStatus := healthinterface.Health{}
	healthStatus.ServiceName = h.config.ServiceName
	healthStatus.ServiceProvider = h.config.ServiceProvider
	healthStatus.ServiceVersion = h.config.ServiceVersion
	healthStatus.TimeStampUTC = time.Now().UTC()
	healthStatus.ServiceStatus = healthinterface.ServiceRunning

	inbound := []healthinterface.InboundInterface{}
	outbound := []healthinterface.OutboundInterface{}

	// add mongo connection status
	mongo := h.db.Health()
	outbound = append(outbound, *mongo)

	// add internal server details
	name, _ := os.Hostname()

	server := healthinterface.InboundInterface{}
	server.Hostname = name
	server.OS = runtime.GOOS
	server.TimeStampUTC = time.Now().UTC()
	server.ApplicationName = h.config.ServiceName
	server.ConnectionStatus = healthinterface.ConnectionActive

	exIP, err := externalIP()
	if err != nil {
		h.logger.Printf("%s:%s Failed to obtain inbound ip address with error %v\n", models.HdrRequestID, txID, err)
		server.ConnectionStatus = healthinterface.ConnectionDisconnected
	}
	server.Address = exIP
	inbound = append(inbound, server)

	healthStatus.InboundInterfaces = inbound
	healthStatus.OutboundInterfaces = outbound
	renderers.RenderJSON(w, r, txID, healthStatus)
	return

}
