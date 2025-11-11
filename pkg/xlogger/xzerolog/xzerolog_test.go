package xzerolog_test

import (
	"io/fs"
	"os"
	"path/filepath"
	"testing"

	"github.com/agussyahrilmubarok/gox/pkg/xlogger/xzerolog"
	"github.com/rs/zerolog"
)

func tempLogPath(t *testing.T) string {
	t.Helper()
	dir := t.TempDir()
	return filepath.Join(dir, "test.log")
}

func TestNewLogger_Success(t *testing.T) {
	logPath := tempLogPath(t)
	logger, err := xzerolog.NewLogger(logPath, "debug")

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if _, err := os.Stat(logPath); os.IsNotExist(err) {
		t.Fatalf("expected log file to exist, but not found: %v", logPath)
	}

	if zerolog.GlobalLevel() != zerolog.DebugLevel {
		t.Errorf("expected global level = Debug, got %v", zerolog.GlobalLevel())
	}

	defer func() {
		if r := recover(); r != nil {
			t.Errorf("logging caused panic: %v", r)
		}
	}()
	logger.Info().Msg("test log message")
}

func TestNewLogger_InvalidLevel_DefaultsToInfo(t *testing.T) {
	logPath := tempLogPath(t)
	_, err := xzerolog.NewLogger(logPath, "invalid-level")

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if zerolog.GlobalLevel() != zerolog.InfoLevel {
		t.Errorf("expected default level = Info, got %v", zerolog.GlobalLevel())
	}
}

func TestNewLogger_CreateDirIfNotExists(t *testing.T) {
	dir := filepath.Join(t.TempDir(), "nested", "dir", "logs")
	logPath := filepath.Join(dir, "app.log")

	_, err := xzerolog.NewLogger(logPath, "info")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		t.Errorf("expected directory to be created: %v", dir)
	}
}

func TestNewLogger_FailToCreateDir(t *testing.T) {
	tmpFile := filepath.Join(t.TempDir(), "notadir")
	if err := os.WriteFile(tmpFile, []byte("data"), fs.ModePerm); err != nil {
		t.Fatalf("setup failed: %v", err)
	}

	logPath := filepath.Join(tmpFile, "log.log")
	_, err := xzerolog.NewLogger(logPath, "info")

	if err == nil {
		t.Fatal("expected error due to invalid directory path, got nil")
	}
}
