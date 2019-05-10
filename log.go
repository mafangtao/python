package logging

import (
	"log"
)

type Level int

var (
	F *os.File

	DefaultPrefix = ""
	DefaultCallerDepth = 2

	logger *log.Logger
	logPrefix = ""
	levelFlags = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
)
