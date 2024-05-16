package middleware

import (
	"log"
	"net/http"
	"time"
)

type wrappedWriter struct {
	http.ResponseWriter // type embedding or composition - kind of inheritance 
	statusCode int
}

// Method Overriding
func (w *wrappedWriter) WriteHeader(statusCode int) {
	w.ResponseWriter.WriteHeader(statusCode)
	w.statusCode = statusCode
}

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(rw http.ResponseWriter, r *http.Request){
			start := time.Now()
			wrw := &wrappedWriter{
				ResponseWriter: rw,
				statusCode: http.StatusOK,
			}
			next.ServeHTTP(wrw,r)
			log.Println(r.Method, r.URL.Path, wrw.statusCode, time.Since(start))
		},
	)
}