package main

import (
	"fmt"
	"os"
	"time"

	"github.com/apex/log"
	"github.com/apex/log/handlers/multi"
)

type handler struct {
	filep *os.File
}

func newhandler(filep *os.File) *handler {
	return &handler{filep: filep}
}

func (h *handler) HandleLog(e *log.Entry) error {
	_, err := fmt.Fprintf(h.filep, "%s\n", e.Message)
	return err
}

func startlogging() {
	filep, err := os.Create(time.Now().Format("20060102150405.txt"))
	if err != nil {
		log.WithError(err).Fatal("cannot open logfile")
	}
	log.SetHandler(multi.New(
		newhandler(os.Stderr),
		newhandler(filep),
	))
}
