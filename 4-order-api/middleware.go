package orderapi

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
)

type key string

const CtxPhoneKey key = "phonekey"

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

func IsAuthed(next http.Handler, userRepo *UserRepository) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHdr := r.Header.Get("Authorization")
		if !strings.HasPrefix(authHdr, "Bearer ") {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(http.StatusText(http.StatusUnauthorized)))
			return
		}
		token := strings.TrimPrefix(authHdr, "Bearer ")
		payloadFromToken := jwt.MapClaims{}
		jwt.ParseWithClaims(token, payloadFromToken, func(t *jwt.Token) (any, error) {
			return nil, nil
		})

		user, err := userRepo.GetUserByPhone(payloadFromToken["phone"].(string))
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(http.StatusText(http.StatusUnauthorized)))
			return
		}
		tkn, err := jwt.Parse(token, func(t *jwt.Token) (any, error) {
			return []byte(user.Secret), nil
		})
		if err != nil || !tkn.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(http.StatusText(http.StatusUnauthorized)))
			return
		}
		ctx := context.WithValue(r.Context(), CtxPhoneKey, user.Phone)
		req := r.WithContext(ctx)
		next.ServeHTTP(w, req)
	})
}
