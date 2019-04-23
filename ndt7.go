package main

import (
	"context"
	"encoding/json"
	"time"

	"github.com/apex/log"
	"github.com/measurement-kit/engine/nettest/ndt7/runner/model"
)

type subfn func(context.Context, string) (<-chan model.Measurement, error)

func subtest(name, FQDN string, fn subfn) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	log.Infof("=== BEGIN %s %s ===", FQDN, name)
	defer log.Infof("=== END %s %s ===", FQDN, name)
	ch, err := fn(ctx, FQDN)
	if err != nil {
		log.Warnf("subtest function error: %s", err.Error())
		return false
	}
	for ev := range ch {
		data, err := json.Marshal(ev)
		if err != nil {
			log.Fatalf("cannot serialize JSON: %s", err.Error())
		}
		log.Infof("measurement: %s", string(data))
	}
	return true
}
