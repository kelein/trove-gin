package log

import (
	"log/slog"
	"os"
	"path/filepath"
)

const logTime = "2006-01-02T15:04:05.000"

// SetupSlog setting slog default logger
func SetupSlog() {
	replace := func(groups []string, a slog.Attr) slog.Attr {
		if a.Key == slog.SourceKey {
			source := a.Value.Any().(*slog.Source)
			source.File = filepath.Base(source.File)
		}
		if a.Value.Kind() == slog.KindTime {
			return slog.String(a.Key, a.Value.Time().Format(logTime))
		}
		return a
	}

	// * Text Log Format
	logger := slog.New(slog.NewTextHandler(
		os.Stdout, &slog.HandlerOptions{
			AddSource:   true,
			ReplaceAttr: replace,
		},
	))
	slog.SetDefault(logger)
}
