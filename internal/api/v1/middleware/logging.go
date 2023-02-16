package middleware

import (
	"net/http"
	"social_network/utils/logger"
	"time"
)

// logging for http handlers

type ResponseRecorder struct {
	http.ResponseWriter
	StatusCode int
	Body       []byte

	Response interface {
		WriteHeader(statusCode int)
		Write(body []byte) (int, error)
	}
}

func (rec *ResponseRecorder) Write(body []byte) (int, error) {
	rec.Body = body
	return rec.ResponseWriter.Write(body)
}

func (res *ResponseRecorder) WriteHeader(statusCode int) {
	res.StatusCode = statusCode
	res.ResponseWriter.WriteHeader(statusCode)
}

func Logging(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(wrt http.ResponseWriter, req *http.Request) {
		start_time := time.Now()
		rec := &ResponseRecorder{
			ResponseWriter: wrt,
			StatusCode:     http.StatusOK,
		}

		handler.ServeHTTP(wrt, req)
		duration := time.Since(start_time) // to calculate the duretion of work

		if rec.StatusCode != http.StatusOK {
			logger.Error("Body", rec.Body)
		}

		logger.Str("protocol", "http",
			"method", req.Method,
			"path", req.RequestURI,
			"Duration ", duration,
			"Status Code ", rec.StatusCode,
			"Status text", http.StatusText(rec.StatusCode))
	})
}
