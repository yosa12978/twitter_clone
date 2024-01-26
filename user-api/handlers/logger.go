package handlers

import (
	"net/http"
	"time"

	"github.com/yosa12978/twitter/user-api/logging"
)

var (
	logger = logging.New("logging middleware")
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		checkpoint := time.Now().UnixMicro()
		next.ServeHTTP(w, r)
		logger.Fields(map[string]interface{}{"latency(μs)": time.Now().UnixMicro() - checkpoint})
	})
}
