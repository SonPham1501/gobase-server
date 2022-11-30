package app

import (
	"net"
	"net/http"
	"sync"

	"com.son.server.base/services/timezones"
	"github.com/gorilla/mux"
)

type ServiceKey string

const (
	RouterKey ServiceKey = "router"
)

type Server struct {
	// RootRouter is the starting point for all HTTP requests to the server.
	RootRouter *mux.Router

	// LocalRouter is the starting point for all the local UNIX socket
	// requests to the server
	LocalRouter *mux.Router

	// Router is the starting point for all web, api4 and ws requests to the server. It differs
	// from RootRouter only if the SiteURL contains a /subpath.
	Router *mux.Router

	Server     *http.Server
	ListenAddr *net.TCPAddr

	timezones *timezones.Timezones

	serverMux sync.RWMutex

	products map[string]Product
}

func NewServer(options ...Option) (*Server, error)  {
	rootRouter := mux.NewRouter()
	localRouter := mux.NewRouter()

	s := &Server{
		RootRouter: rootRouter,
		LocalRouter: localRouter,
		timezones: timezones.New(),
		products: make(map[string]Product),
	}

	return s, nil
}
