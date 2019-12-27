package main

import (
	coachee "coachee-backend/gen/coachee"
	"coachee-backend/internal/repository/mysql"
	"coachee-backend/internal/repository/mysql/connector"
	"coachee-backend/internal/service"
	"coachee-backend/internal/stripe"
	"context"
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"os/signal"
	"strings"
	"sync"
)

func main() {
	// Define command line flags, add any other flag required to configure the
	// service.
	var (
		hostF     = flag.String("host", "development", "Server host (valid values: development)")
		domainF   = flag.String("domain", "", "Host domain name (overrides host domain specified in service design)")
		httpPortF = flag.String("http-port", "", "HTTP port (overrides host HTTP port specified in service design)")
		secureF   = flag.Bool("secure", false, "Use secure scheme (https or grpcs)")
		dbgF      = flag.Bool("debug", false, "Log request and response bodies")
		stripeKey = flag.String("stripe-Key", "sk_test_yKV7Mo9kSpokxpFvwxKRtbyd00knjXTpJh", "stripe key")
	)
	flag.Parse()
	// initialize app context
	appCtx := context.Background()

	// Setup logger. Replace logger with your own log package of choice.
	logger := log.New(os.Stderr, "[coacheeapi] ", log.Ltime)

	// Get connector
	conn, err := connector.Connect(appCtx)
	if err != nil {
		logger.Panicln("failed to connect to db:", err.Error())
	}

	// Initialize repositories
	coachRepository := mysql.NewCoachRepository(conn)
	clientRepository := mysql.NewClientRepository(conn)

	// Initialize stripe client
	stripeClient := stripe.NewClient(appCtx, *stripeKey)

	// Initialize the services.
	coacheeSvc := service.NewCoachee(appCtx, coachRepository, clientRepository, stripeClient)

	// Wrap the services in endpoints that can be invoked from other services
	// potentially running in different processes.
	coacheeEndpoints := coachee.NewEndpoints(coacheeSvc)

	// Create channel used by both the signal handler and server goroutines
	// to notify the main goroutine when to stop the server.
	errc := make(chan error)

	// Setup interrupt handler. This optional step configures the process so
	// that SIGINT and SIGTERM signals cause the services to stop gracefully.
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		errc <- fmt.Errorf("%s", <-c)
	}()

	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())

	// Start the servers and send errors (if any) to the error channel.
	switch *hostF {
	case "development":
		{
			addr := "http://localhost:80"
			u, err := url.Parse(addr)
			if err != nil {
				fmt.Fprintf(os.Stderr, "invalid URL %#v: %s\n", addr, err)
				os.Exit(1)
			}
			if *secureF {
				u.Scheme = "https"
			}
			if *domainF != "" {
				u.Host = *domainF
			}
			if *httpPortF != "" {
				h := strings.Split(u.Host, ":")[0]
				u.Host = h + ":" + *httpPortF
			} else if u.Port() == "" {
				u.Host += ":80"
			}
			handleHTTPServer(ctx, u, coacheeEndpoints, &wg, errc, logger, *dbgF)
		}

	default:
		fmt.Fprintf(os.Stderr, "invalid host argument: %q (valid hosts: development)\n", *hostF)
	}

	// Wait for signal.
	logger.Printf("exiting (%v)", <-errc)

	// Send cancellation signal to the goroutines.
	cancel()

	wg.Wait()
	logger.Println("exited")
}
