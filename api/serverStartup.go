package api

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/hlog"
)

/*
func hello(w http.ResponseWriter, req *http.Request) {

		fmt.Fprintf(w, "hello\n")
	}

func headers(w http.ResponseWriter, req *http.Request) {

		for name, headers := range req.Header {
			for _, h := range headers {
				fmt.Fprintf(w, "%v: %v\n", name, h)
			}
		}
	}
*/
type ServerOptions struct {
	Logger            zerolog.Logger
	ServerHostName    string
	ServerHostPort    int
	ApiServerHostName string
	LocalEvaluate     bool
	LogGlobalLevel    zerolog.Level
}
type ServerEnvironment struct {
	Logger            zerolog.Logger
	ServerHostName    string
	ServerHostPort    int
	ApiServerHostName string
	ServerStatus      string
	ServerName        string
	LocalEvaluate     bool
	LogGlobalLevel    zerolog.Level
}

func Start(ctx context.Context, options ServerOptions) {

	serverEnvironment := ServerEnvironment{
		Logger:            options.Logger,
		ServerHostName:    options.ServerHostName,
		ServerHostPort:    options.ServerHostPort,
		ApiServerHostName: options.ApiServerHostName,
		LocalEvaluate:     options.LocalEvaluate,
		LogGlobalLevel:    options.LogGlobalLevel,
	}
	//specify endpoints, handler functions and HTTP method
	hostName := serverEnvironment.ServerHostName
	port := serverEnvironment.ServerHostPort
	listener := NewHttpServer(hostName, port, serverEnvironment)

	//start and listen to requests
	options.Logger.Info().Msg("Server running - listening on port 8080")
	listener.ListenAndServe()
}
func setupHttpRouter(serverEnvironment ServerEnvironment) http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/health-check", HealthCheck).Methods("GET")
	router.HandleFunc("/api/evaluate", serverEnvironment.apiEvaluate_POST).Methods("POST")
	router.HandleFunc("/calculator", serverEnvironment.calculatorView).Methods("GET")
	router.HandleFunc("/evaluate", serverEnvironment.webEvaluate).Methods("POST")
	return router
}

func NewHttpServer(address string, port int, serverEnvironment ServerEnvironment) *http.Server {
	router := setupHttpRouter(serverEnvironment)

	//server := alice.New(Logger).Then(router)
	c := alice.New()
	//zeroLog := serverEnvironment.Log
	// Install the logger handler with default output on the console
	c = c.Append(hlog.NewHandler(serverEnvironment.Logger))

	// Install some provided extra handler to set some request's context fields.
	// Thanks to that handler, all our logs will come with some prepopulated fields.
	c = c.Append(hlog.AccessHandler(func(r *http.Request, status, size int, duration time.Duration) {
		hlog.FromRequest(r).Info().
			Str("method", r.Method).
			Stringer("url", r.URL).
			Int("status", status).
			Int("size", size).
			Dur("duration", duration).
			Msg("server.NewHttpServer")
	}))
	c = c.Append(hlog.RemoteAddrHandler("ip"))
	c = c.Append(hlog.UserAgentHandler("user_agent"))
	c = c.Append(hlog.RefererHandler("referer"))
	c = c.Append(hlog.RequestIDHandler("req_id", "Request-Id"))
	server := c.Then(router)
	listener := &http.Server{
		Addr:    "localhost:8080",
		Handler: server,
	}

	return listener
}
func Logger(h http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		startTime := time.Now()
		h.ServeHTTP(writer, request)
		log.Printf("%s - %s (%v)\n", request.Method, request.URL.Path, time.Since(startTime))
	})
}
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	//specify status code
	w.WriteHeader(http.StatusOK)

	//update response writer
	fmt.Fprintf(w, "API is up and running")
}
