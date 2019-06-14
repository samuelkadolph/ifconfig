package main

import (
	"fmt"
	"net"
	"net/http"
	"strings"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Wait() {
	select {}
}

func NewServer(listens Listens) (*Server, error) {
	handler := http.NewServeMux()
	handler.HandleFunc("/", Root)

	httpServer := &http.Server{
		Handler: handler,
	}

	for _, listen := range listens {
		l, err := net.Listen("tcp", listen)
		if err != nil {
			fmt.Printf("Failed to listen on '%s': %s\n", listen, err)
			return nil, err
		}
		go httpServer.Serve(l)
		fmt.Printf("Listening on %s\n", listen)
	}

	return &Server{httpServer: httpServer}, nil
}

func DetermineIP(req *http.Request) string {
	if forwardedFor := req.Header.Get("X-Forwarded-For"); forwardedFor != "" {
		if ipChain := strings.Split(forwardedFor, ", "); len(ipChain) > 0 {
			return ipChain[0]
		}
	}

	if realIP := req.Header.Get("X-Real-IP"); realIP != "" {
		return realIP
	}

	if parts := strings.Split(req.RemoteAddr, ":"); len(parts) > 0 {
		return parts[0]
	}

	return ""
}

func Root(writer http.ResponseWriter, req *http.Request) {
	remoteIP := DetermineIP(req)

	fmt.Printf("%s %s %s '%s' '%s'\n", req.Method, req.Proto, req.RequestURI, remoteIP, req.Header.Get("User-Agent"))

	writer.WriteHeader(200)
	fmt.Fprintf(writer, "%s", remoteIP)
}
