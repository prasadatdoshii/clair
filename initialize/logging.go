package initialize

import (
	"context"
	"os"
	"strings"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/quay/clair/v4/config"
)

// Logging will set the global logging level for Clair,
// create a global logger embedded into a CTX,
// and sets this CTX as our application's GlobalCTX.
func (i *Init) Logging() error {
	// global log level
	level := LogLevel(i.conf.LogLevel)
	zerolog.SetGlobalLevel(level)

	// attach global logger to ctx
	i.GlobalCTX, i.GlobalCancel = context.WithCancel(context.Background())
	globalLogger := log.With().Timestamp().Logger()
	i.GlobalCTX = globalLogger.WithContext(i.GlobalCTX)

	globalLogger.Info().Str("component", "init/Init.Logging").Msg("logging initialized")
	return nil
}

// LogLevel does a string-to-level mapping.
func LogLevel(level string) zerolog.Level {
	level = strings.ToLower(level)
	switch level {
	case "debug-color":
		// set global logger to use ConsoleWriter for colorized output
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
		return zerolog.DebugLevel
	case "debug":
		return zerolog.DebugLevel
	case "info":
		return zerolog.InfoLevel
	case "warn":
		return zerolog.WarnLevel
	case "error":
		return zerolog.ErrorLevel
	case "fatal":
		return zerolog.FatalLevel
	case "panic":
		return zerolog.PanicLevel
	default:
		return zerolog.InfoLevel
	}
}

// Logging configures a logger according to the provided configuration and returns
// a configured Context.
func Logging(ctx context.Context, cfg *config.Config) (context.Context, error) {
	zerolog.SetGlobalLevel(LogLevel(cfg.LogLevel))
	l := log.With().Timestamp().Logger()
	ctx = l.WithContext(ctx)
	l.Debug().Str("component", "initialize/Logging").Msg("logging initialized")
	return ctx, nil
}
