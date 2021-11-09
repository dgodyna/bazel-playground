package main

import (
	"context"
	"github.com/cenkalti/backoff/v4"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	log.Logger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout}).With().Timestamp().Logger().Level(zerolog.InfoLevel)

	ctx, cancelFunc := context.WithCancel(context.Background())
	wg := sync.WaitGroup{}

	b := backoff.NewExponentialBackOff()

	// Start background printing
	wg.Add(1)
	go func() {
		defer wg.Done()
		printStatus(ctx, b)
	}()

	// Graceful shutdown by cancelling application context on system signal
	wg.Add(1)
	go func() {
		defer wg.Done()
		termChan := make(chan os.Signal, 1)
		signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM)

		select {
		// received cancellation signal
		case sig := <-termChan:
			log.Info().Msgf("received cancellation signal '%s'", sig.String())
			cancelFunc()
		}
	}()

	wg.Wait()
}

// just print uptime with exponential backoff until context will be cancelled
func printStatus(ctx context.Context, b backoff.BackOff) {
	startTime := time.Now()
	for {
		nextBackOff := b.NextBackOff()
		log.Debug().Msgf("next backoff: %v", nextBackOff)
		select {
		case <-ctx.Done():
			return
		case <-time.Tick(nextBackOff):
			log.Info().Msgf("Program uptime: '%s'", time.Since(startTime))
		}
	}

}
