package http

import (
	"context"
	"fmt"
	"log"

	"net/http"
	"sync"
	"time"

	"github.com/gorilla/mux"
)

type Server struct {
	port   int
	server http.Server
	router *mux.Router
	wg     sync.WaitGroup
}

func NewServer(port int, router *mux.Router) *Server {
	return &Server{
		port:   port,
		router: router,
	}
}

// Start will start the server and if it cannot bind to the port
// it will exit with a fatal log message
func (c *Server) Start() {
	_, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Create the HTML Server
	c.server = http.Server{
		Addr:           fmt.Sprintf(":%d", c.port),
		Handler:        c.router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: http.DefaultMaxHeaderBytes,
	}

	// Add to the WaitGroup for the listener goroutine
	c.wg.Add(1)

	// Start the listener
	go func() {
		if err := c.server.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("Server listen and server error: %v", err) // TODO: should try handle this error
		}
		c.wg.Done()
	}()
}

// Stop stops the API Server
func (c *Server) Stop() error {
	// Create a context to attempt a graceful 5-second shutdown.
	const timeout = 5 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	// Attempt the graceful shutdown by closing the listener
	// and completing all inflight requests
	if err := c.server.Shutdown(ctx); err != nil {
		// Looks like we timed out on the graceful shutdown. Force close.
		if err = c.server.Close(); err != nil {
			return err
		}
		return err
	}

	c.wg.Wait()
	return nil
}
