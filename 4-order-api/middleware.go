package orderapi

import (
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

func MidLog(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next.ServeHTTP(w, r)

		logrus.WithFields(logrus.Fields{
			"method":           r.Method,
			"path":             r.URL.Path,
			"serve delay (us)": time.Since(start).Microseconds(),
		}).Info("new request is received")
	})
}
