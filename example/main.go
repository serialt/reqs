package main

import (
	"fmt"

	"github.com/serialt/req"
)

func main() {
	// req.NewCAPool("/tmp/ccc/ca.crt")

	// certificate, _ := req.NewCertificateFromFiles("xxx.crt", "xxx.key")
	testConfig := req.TLSConfig{
		InsecureSkipVerify: true,
		ClientCAs:          req.NewCAPoolFromFiles("/tmp/ccc/ca.crt"),
		// Certificates:       []tls.Certificate{certificate},
	}

	resp, err := req.Get("https://example.com", testConfig)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resp.Text())
	}
}
