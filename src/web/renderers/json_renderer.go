package renderers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

// RenderJSON constructs and sends the JSON response
func RenderJSON(w http.ResponseWriter, r *http.Request, txID string, data interface{}) {
	logger := log.New(os.Stdout, "rendereres :=> ", log.LstdFlags)

	buff, err := json.Marshal(data)
	if err != nil {
		logger.Printf("%s Failed unmarshal health object error %v\n", txID, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, string(buff))
	return
}
