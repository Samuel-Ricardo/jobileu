package config

import (
	"io"
	"log"
	"os"
)

type Logger struct {
	debug   *log.Logger
	info    *log.Logger
	warning *log.Logger
	error   *log.Logger
	writer  io.Writer
}

func NewLogger(p string) *Logger {
	writter := io.Writer(os.Stdout)
	logger := log.New(writter, p, log.Ldate|log.Ltime)

	return &Logger{
		debug:   log.New(writter, "[DEBUG] | ", logger.Flags()),
		info:    log.New(writter, "[INFO] | ", logger.Flags()),
		warning: log.New(writter, "[WARNING] | ", logger.Flags()),
		error:   log.New(writter, "[ERROR] | ", logger.Flags()),
		writer:  writter,
	}
}
