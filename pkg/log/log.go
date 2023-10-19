package log

import (
	"io"
	golog "log"
)

var (
	Debug = golog.New(io.Discard, "", golog.LstdFlags)
	Info  = golog.Default()
)
