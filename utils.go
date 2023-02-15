package req

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"io/ioutil"
	"net/url"
	"os"
)

// handle URL params
func buildURLParams(userURL string, params map[string]string) (string, error) {
	Url, err := url.Parse(userURL)
	if err != nil {
		return "", errors.New(err.Error())
	}
	mkparams := url.Values{}
	for k, v := range params {
		mkparams.Set(k, v)

	}
	Url.RawQuery = Url.RawQuery + mkparams.Encode()

	userURL = Url.String()
	return userURL, nil
}

func openFile(filename string) *os.File {
	r, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	return r
}

func NewCAPool(crts ...[]byte) *x509.CertPool {
	clientCAPool := x509.NewCertPool()
	for _, crt := range crts {
		clientCAPool.AppendCertsFromPEM(crt)
	}
	return clientCAPool
}

func NewCAPoolFromFiles(files ...string) *x509.CertPool {
	clientCAPool := x509.NewCertPool()
	for _, file := range files {
		crtData, _ := ioutil.ReadFile(file)
		clientCAPool.AppendCertsFromPEM(crtData)
	}
	return clientCAPool
}

func NewCertificate(certData []byte, keyData []byte) (tls.Certificate, error) {
	clientCrt, err := tls.X509KeyPair(certData, keyData)
	if err != nil {
		return tls.Certificate{}, err
	}
	return clientCrt, nil

}

func NewCertificateFromFiles(certFile string, keyFile string) (tls.Certificate, error) {
	clientCrt, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return tls.Certificate{}, err
	}
	return clientCrt, nil
}
