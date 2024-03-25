package config

import (
	"io"
	"log"
)

type Logger struct {
	debug   *log.Logger
	info    *log.Logger
	warning *log.Logger
	error   *log.Logger
	writer  io.Writer
}
