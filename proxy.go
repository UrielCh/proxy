package main

import (
	"flag"
	"log"
	"net"
	"net/http"

	"github.com/elazarl/goproxy"
)

func main() {
	verbose := flag.Bool("v", false, "should every proxy request be logged to stdout")
	addr := flag.String("addr", ":8080", "proxy listen address")
	bind := flag.String("bind", "", "local IP to bind ex: 192.168.0.1:")
	flag.Parse()
	proxy := goproxy.NewProxyHttpServer()

	proxy.Verbose = *verbose
	if *bind != "" {
		tcpAddr, err := net.ResolveTCPAddr("tcp", *bind)
		if err != nil {
			log.Fatalf("failt to read IP %s, Err: %v", *bind, err)
		}
		proxy.Bind = tcpAddr
		proxy.Tr.DialContext = (&net.Dialer{LocalAddr: &net.TCPAddr{IP: tcpAddr.IP}}).DialContext
	}
	log.Printf("Listening %s", *addr)
	log.Fatal(http.ListenAndServe(*addr, proxy))
}
