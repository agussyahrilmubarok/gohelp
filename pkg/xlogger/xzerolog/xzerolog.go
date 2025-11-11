package xzerolog

import (
	"os"
	"path/filepath"

	"github.com/rs/zerolog"
)

// NewLogger creates a Zerolog logger that writes to both a file and stdout.
// logPath: path to the log file.
// logLevel: logging level as a string (e.g., "info", "debug").
// Returns a zerolog.Logger and an error if any occurs.
func NewLogger(logPath, logLevel string) (zerolog.Logger, error) {
	// Ensure the directory for the log file exists
	logDir := filepath.Dir(logPath)
	if err := os.MkdirAll(logDir, os.ModePerm); err != nil {
		return zerolog.Logger{}, err
	}

	// Open or create the log file for appending logs
	logFile, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return zerolog.Logger{}, err
	}

	// Write logs to both stdout and the log file
	multi := zerolog.MultiLevelWriter(os.Stdout, logFile)

	// Parse log level; default to InfoLevel if invalid
	level, err := zerolog.ParseLevel(logLevel)
	if err != nil {
		level = zerolog.InfoLevel
	}
	zerolog.SetGlobalLevel(level)

	// Create a new logger with timestamp
	logger := zerolog.New(multi).With().Timestamp().Logger()

	return logger, nil
}
