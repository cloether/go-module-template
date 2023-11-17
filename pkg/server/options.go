package server

type Option func(*Server)

type OptionFunc func(addr string) Option

func WithAddr(addr string) Option { return func(server *Server) { server.SetAddr(addr) } }

func WithPort(port uint16) Option { return func(server *Server) { server.SetPort(port) } }
