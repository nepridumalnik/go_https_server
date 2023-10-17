package web

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type configure struct {
	Address  string `json:"address,omitempty"`
	Port     uint16 `json:"port,omitempty"`
	CertFile string `json:"cert_file,omitempty"`
	KeyFile  string `json:"key_file,omitempty"`
}

func readConfigFromFile(name string) (*configure, error) {
	file, err := os.ReadFile(name)

	if err != nil {
		return nil, err
	}

	conf := configure{
		Address:  "127.0.0.1",
		Port:     8080,
		CertFile: "./crypto/certificate.pem",
		KeyFile:  "./crypto/private_key.pem",
	}

	err = json.Unmarshal(file, &conf)

	if err != nil {
		return nil, err
	}

	return &conf, nil
}

type webServer struct {
	server *http.Server
	config *configure
}

func MakeWebServer(config string) (*webServer, error) {
	conf, err := readConfigFromFile(config)

	if err != nil {
		log.Fatal(err)
	}

	server := &http.Server{
		Addr: fmt.Sprintf("%s:%d", conf.Address, conf.Port),
		TLSConfig: &tls.Config{
			Certificates: make([]tls.Certificate, 1),
		},
	}

	server.TLSConfig.Certificates[0], err = tls.LoadX509KeyPair(conf.CertFile, conf.KeyFile)

	if err != nil {
		log.Fatal(err)
	}

	return &webServer{
		server: server,
		config: conf,
	}, nil
}

func (ws *webServer) Run() error {
	fmt.Printf("Running server on: https://%s:%d\n", ws.config.Address, ws.config.Port)
	return ws.server.ListenAndServeTLS("", "")
}
