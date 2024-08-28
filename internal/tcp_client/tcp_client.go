package tcpclient

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

/**
* Simulating and testing a tcp client connection and sending it messages from cmd.
**/

func SimulateTCPConn(port int) {
	conn, err := net.Dial("tcp", fmt.Sprintf(":%d", port))

	if err != nil {
		fmt.Println("Error when attempting to dial to tcp connection.")
	}

	defer conn.Close()

	// read in arguments to send over this tcp connection
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter message: ")
		msg, _ := reader.ReadString('\n')
		msg = strings.TrimSpace(msg)

		_, err := conn.Write([]byte(msg))
		if err != nil {
			fmt.Println("Error sending message:", err)
			return
		}

		// read response and log it
		res, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println("Error when attempting to read from connection.")
		}

		fmt.Println("reply from server:", res)
	}

}
