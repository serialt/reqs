/*
 * @Description   : Blue Planet
 * @Author        : serialt
 * @Email         : tserialt@gmail.com
 * @Created Time  : 2023-04-18 22:02:19
 * @Last modified : 2023-10-11 20:12:16
 * @FilePath      : /req/example/main.go
 * @Other         :
 * @              :
 *
 *
 *
 */
package main

import (
	"fmt"

	req "github.com/serialt/reqs"
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
