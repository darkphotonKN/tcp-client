package main

import tcpclient "github.com/darkphotonKN/tcp-client/internal/tcp_client"

func main() {

	// general
	// tcpclient.SimulateTCPConn(5555)

	// starlight cargo
	tcpclient.SimulateTCPConnWithLogin(3600)
}
