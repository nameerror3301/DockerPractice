package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"sync"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type Server struct {
	wg       *sync.WaitGroup
	server   *http.Server
	Callback func()
	use      bool
}

func using(s *Server) {
	if s.use {
		s.use = false
	} else {
		s.use = true
	}
}

func NewServer(errChan chan<- error, address string, port int, callback func(), routs *mux.Router) (*Server, error) {
	if errChan == nil {
		errChan = make(chan<- error, 10)
	}

	srv := &Server{Callback: callback, use: false, wg: &sync.WaitGroup{}}

	handler := handlers.CORS(
		handlers.AllowCredentials(),
		handlers.AllowedHeaders([]string{"X-Requested-With"}),
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "OPTIONS"}),
	)(routs)

	if err := srv.RunServer(errChan, address, port, &handler); err != nil {
		return nil, err
	}

	return srv, nil
}

func (s *Server) Stop() error {
	s.server.Shutdown(context.Background())
	s.wg.Wait()
	return nil
}

func (s *Server) RunServer(errChan chan<- error, address string, port int, handler *http.Handler) error {
	if s.use {
		return fmt.Errorf("service started")
	}

	// Создание сервера TCP.
	listen, err := net.Listen("tcp4", fmt.Sprintf("%s:%d", address, port))
	if err != nil {
		return fmt.Errorf("filed listen tcp server %s", err.Error())
	}

	s.wg.Add(1)

	srvChan := make(chan error, 10)

	go func(s *Server, errChan chan error) {
		using(s)
		defer using(s)
		defer s.Callback()

		//Set server and use logrus
		s.server = &http.Server{
			ErrorLog: log.New(
				logrus.StandardLogger().Writer(),
				"[HTTP]\t",
				log.Ldate),
			Handler: *handler,
		}

		errChan <- s.server.Serve(listen)

		s.wg.Done()
	}(s, srvChan)

	return nil
}
