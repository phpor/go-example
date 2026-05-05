package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"net"
	"net/http"
)

func main() {

	//var typ, entry, pin string
	//flag.StringVar(&typ, "type", "idfa", "type")
	//flag.StringVar(&entry, "entry", "test", "entry")
	//flag.StringVar(&pin, "pin", "", "pin")
	//flag.Parse()

	client := http.Client{
		Transport: &http.Transport{
			DialTLSContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
				return tls.Dial(network, addr, &tls.Config{
					InsecureSkipVerify: false,
					ServerName:         "",
					MinVersion:         tls.VersionTLS12,
					CipherSuites: []uint16{
						tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
						tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
						tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
					},
				})
			},
		},
	}

	resp, err := client.Get("https://aaid.uyunad.com/")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))

}
