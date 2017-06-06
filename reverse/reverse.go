package reverse

import (
	"crypto/tls"
	"net/http"
	"net/url"
	"time"

	"net/http/httputil"

	"golang.org/x/crypto/acme/autocert"
)

// Config represents a configuration of the reverse proxy.
type Config struct {
	Addr                string
	ContactEmail        string
	Hosts               []string
	CertificateCacheDir string
	Target              string
}

// Start starts the reverse proxy.
func Start(config *Config) error {
	m := autocert.Manager{
		Email:      config.ContactEmail,
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist(config.Hosts...),
		Cache:      autocert.DirCache(config.CertificateCacheDir),
	}

	tlsConfig := &tls.Config{
		MinVersion:               tls.VersionTLS12,
		PreferServerCipherSuites: true,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
			tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,
			tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
		},
		GetCertificate: m.GetCertificate,
	}

	target, err := url.Parse(config.Target)
	if err != nil {
		return err
	}

	server := http.Server{
		Addr:         config.Addr,
		Handler:      httputil.NewSingleHostReverseProxy(target),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		TLSConfig:    tlsConfig,
	}

	return server.ListenAndServeTLS("", "")
}
