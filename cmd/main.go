package main

import tcpclient "github.com/darkphotonKN/tcp-client/internal/tcp_client"

func main() {

	// tcpclient.SimulateTCPConn(3600)
	tcpclient.SimulateTCPConnWithLogin(3600)
}
