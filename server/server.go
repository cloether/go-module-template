package main

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"template/internal/logging"
	"time"
)

type server struct {
	router *mux.Router
	log    *zap.SugaredLogger
}

func NewServer(ctx context.Context) *server {
	logger := logging.FromContext(ctx)
	return &server{
		router: &mux.Router{},
		log:    logger,
	}
}

func (s *server) Decode(_ http.ResponseWriter, r *http.Request, v interface{}) error {
	return json.NewDecoder(r.Body).Decode(v)
}

func (s *server) HandleTemplate(files ...string) http.HandlerFunc {
	// expensive setup when the handler is first hit to improve startup
	// if handler isn't called, the work here is never done
	var (
		init   sync.Once
		tpl    *template.Template
		tplErr error
	)
	return func(w http.ResponseWriter, r *http.Request) {
		init.Do(func() {
			tpl, tplErr = template.ParseFiles(files...)
		})
		if tplErr != nil {
			http.Error(w, tplErr.Error(), http.StatusInternalServerError)
			return
		}
		// use template
		http.Error(w, tpl.Name(), http.StatusInternalServerError)
	}
}

func (s *server) HandleIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	}
}

func (s *server) Respond(w http.ResponseWriter, _ *http.Request, data interface{}, status int) {
	w.WriteHeader(status)
	if data != nil {
		if err := json.NewEncoder(w).Encode(data); err != nil {
			s.log.Error(err)
		}
	}
	log.Println(http.StatusText(status))
}

func (s *server) Run(ctx context.Context, addr string) {
	logger := logging.FromContext(ctx)

	logger.Debug("starting server on http://%s", addr)
	srv := s.server(addr) // initialize server

	go func() { // run our server in a goroutine so that it does not block.
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalln(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c // Block until we receive our signal.

	// Accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	// Does not block if there no connections, but will otherwise wait until
	// the timeout deadline.
	_ = srv.Shutdown(ctx)

	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services to
	// finalize based on context cancellation.
	_, _ = os.Stderr.Write([]byte("\b\b")) // Remove "^C" from output
	logger.Debug("shutting down")
	os.Exit(0) // exit with status code 0 for successful shutdown
}

func (s *server) routes() {
	s.router.HandleFunc("/", s.HandleIndex()).Methods(http.MethodGet)
	s.router.HandleFunc("/index.html", s.HandleTemplate("index.html"))
}

func (s *server) server(addr string) *http.Server {
	s.routes() // registers routes
	return &http.Server{
		Addr:              addr,
		Handler:           s.router, // Pass our instance of gorilla/mux in.
		TLSConfig:         nil,
		ReadTimeout:       time.Second * 15,
		ReadHeaderTimeout: 0,
		// Good practice to set timeouts to avoid Slowloris attacks.
		// https://en.wikipedia.org/wiki/Slowloris_(computer_security)
		WriteTimeout:   time.Second * 15,
		IdleTimeout:    time.Second * 60,
		MaxHeaderBytes: 0,
		TLSNextProto:   nil,
		ConnState:      nil,
		ErrorLog:       nil,
		BaseContext:    nil,
		ConnContext:    nil,
	}
}
