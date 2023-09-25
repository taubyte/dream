package api

import (
	"context"
	"time"
	goHttp "net/http"

	"github.com/sirupsen/logrus" // Used logrus instead of pterm for structured logging
	httpIface "github.com/taubyte/http"
	http "github.com/taubyte/http/basic"
	"github.com/taubyte/http/options"
	"github.com/taubyte/tau/libdream/common"
	"github.com/taubyte/tau/libdream/services"
)

// multiverseService combines a REST-ful service interface with the core functionalities of a multiverse
type multiverseService struct {
	rest httpIface.Service // Interface for the REST-ful service
	common.Multiverse // Embedded core functionalities for a multiverse
}

// constants for server startup checks.
const serverCheckInterval = 100 * time.Millisecond // Time at which the server's readiness is checked
const serverStartTimeout  = 10 * time.Second // Maximum time to wait for the server to start

/* BigBang launches the service after initializing the multiverse, then checks to see if it is ready by continually
making HTTP requests up until a successful answer is returned or the maximum wait time is achieved. */
func BigBang() error {
	var err error

	// Initialize the multiverse
	srv := &multiverseService{
		Multiverse: services.NewMultiVerse(),
	}

	// Attempt to build a new HTTP service using the settings supplied.
	srv.rest, err = http.New(srv.Context(), options.Listen(common.DreamlandApiListen), options.AllowedOrigins(true, []string{".*"}))
	if err != nil {
		return err // Return error if there is a problem when building the HTTP service.
	}

	// Set up the service's HTTP routes and start it.
	srv.setUpHttpRoutes().Start()

	// Creates a context that will be terminated upon serverStartTimeout.
	waitCtx, waitCtxC := context.WithTimeout(srv.Context(), serverStartTimeout)
	defer waitCtxC() // Ensure that all context-related resources are freed at the end

	// Create a ticker that will cause events at predetermined intervals (serverCheckInterval)
	ticker := time.NewTicker(serverCheckInterval)
	defer ticker.Stop() // To release connected resources, ensure that the ticker is turned off

	for {
		select {
		// Return an error if the maximum wait time is exceeded without the server becoming ready
		case <-waitCtx.Done():
			return waitCtx.Err()

		// Check the server's readiness at the intervals specified by the ticker
		case <-ticker.C:
			// Handle any internal errors that the service reports
			if serviceError := srv.rest.Error(); serviceError != nil {
				logrus.Error("Dreamland failed to start") // Log the error and return it
				return serviceError
			}
			
			// Create a new HTTP GET request
			resp, err := goHttp.Get("http://" + common.DreamlandApiListen)
			defer resp.Body.Close() // Making certain that the body is closed
			
			if err == nil {
				logrus.Info("Dreamland ready") // If the request is successful, return and log a message.
				return nil
			}
		}
	}
}

/* Suggested Changes
	1. Moved the time intervals to constants
	2. Changed pterm logger with logrus to get the structured logs
	3. Changed Time.After with the Ticker because time.After in a loop creates a new timer on each iteration, which could 
		 lead to a slight overhead when the polling interval is very short. On the other hand Ticker is designed for cases
		 where we want to do something at regular intervals, which is precisely the use case here.
	4. Closed response body at the end in order to prevent the resources leaks.
*/