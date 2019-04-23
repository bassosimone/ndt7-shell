package main

import (
	"github.com/measurement-kit/engine/nettest/ndt7/runner"
)

func main() {
	startlogging()
	FQDNs := []string{
		"ndt-iupui-mlab4-sin01.measurement-lab.org",
		"ndt-iupui-mlab4-bom02.measurement-lab.org",
	}
	for _, FQDN := range FQDNs {
		if subtest("download", FQDN, runner.StartDownload) {
			break
		}
	}
	for _, FQDN := range FQDNs {
		if subtest("upload", FQDN, runner.StartUpload) {
			break
		}
	}
}
