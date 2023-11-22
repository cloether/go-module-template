package server

import (
	"context"
	"encoding/json"
	"errors"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/cloether/go-module-template/internal/env"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

type Server struct {
	addr   string
	port   uint16
	router *mux.Router
	log    *zap.SugaredLogger
}

func New(ctx context.Context, options ...Option) *Server {
	// create default server
	logger := zap.SugaredLogger{}
	server := &Server{
		router: &mux.Router{},
		log:    &logger,
	}

	server.log.With("starting server",
		"environment", env.FromContextWithDefault(ctx, "development"))

	// loop through each option
	for _, option := range options {
		// call the option giving the instantiated *Server as the argument
		option(server)
	}

	environment := env.FromContextWithDefault(ctx, "development")

	server.log.With("starting server", "addr", server.addr, "port", server.port, "environment", environment)

	return server
}

func (s *Server) Addr() string {
	return s.addr
}

func (s *Server) SetAddr(addr string) {
	s.addr = addr
}

func (s *Server) Port() uint16 {
	return s.port
}

func (s *Server) SetPort(port uint16) {
	s.port = port
}

func (s *Server) Decode(_ http.ResponseWriter, r *http.Request, v interface{}) error {
	return json.NewDecoder(r.Body).Decode(&v)
}

func (s *Server) HandleTemplate(files ...string) http.HandlerFunc {
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

func (s *Server) HandleIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	}
}

func (s *Server) Respond(w http.ResponseWriter, _ *http.Request, data interface{}, status int) {
	w.WriteHeader(status)

	if data != nil {
		if err := json.NewEncoder(w).Encode(data); err != nil {
			s.log.Error(err)
		}
	}

	log.Println(http.StatusText(status))
}

// Serve is the Server Entry Point
func (s *Server) Serve(ctx context.Context, version string) {
	s.log.With("server", "version", version)

	ctx, cancel := context.WithCancel(ctx)

	// trap ctrl+c and call cancel on the context
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	defer func() {
		signal.Stop(c)
		cancel()
	}()

	go func() {
		select {
		case <-c:
			cancel()
		case <-ctx.Done():
		}
	}()

	// run the command
	s.Run(ctx, version)
}

func (s *Server) Run(ctx context.Context, addr string) {
	//goland:noinspection HttpUrlsUsage
	s.log.With("starting server", "addr", "http://%s")

	srv := s.server(addr) // initialize server

	go func() { // run our server in a goroutine so that it does not block.
		if err := srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			log.Fatalln(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c // Block until we receive our signal.

	// Accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(
		context.Background(),
		100*time.Millisecond,
	)
	defer cancel()

	// does not block if there are no connections,
	// but will otherwise wait until the timeout deadline.
	_ = srv.Shutdown(ctx)

	// optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services to
	// finalize based on context cancellation.
	_, _ = os.Stderr.Write([]byte("\b\b")) // Remove "^C" from output

	s.log.Debug("shutting down")
	os.Exit(0) // exit with status code 0 for successful shutdown
}

func (s *Server) routes() {
	s.router.HandleFunc("/", s.HandleIndex()).Methods(http.MethodGet)
	s.router.HandleFunc("/index.html", s.HandleTemplate("index.html"))
}

func (s *Server) server(addr string) *http.Server {
	s.routes() // registers routes
	return &http.Server{
		Addr:         addr,
		Handler:      s.router, // Pass our instance of gorilla/mux in.
		ReadTimeout:  time.Second * 15,
		WriteTimeout: time.Second * 15,
		IdleTimeout:  time.Second * 60,
	}
}
