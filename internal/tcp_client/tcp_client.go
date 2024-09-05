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
		return
	}

	defer conn.Close()

	// read in arguments to send over this tcp connection
	reader := bufio.NewReader(os.Stdin)

	for {

		fmt.Print("Enter message: ")
		msg, _ := reader.ReadString('\n')

		_, err = conn.Write([]byte(msg))
		if err != nil {
			fmt.Println("Error sending message:", err)
			return
		}

		// read response and log it
		// res, err := bufio.NewReader(conn).ReadString('\n')
		// if err != nil {
		// 	fmt.Println("Error when attempting to read from connection.")
		// }
		//
		// fmt.Println(res)
	}

}

/**
* For simulating interaction between a TCP server that requires authetication before accessing the commands.
**/

const (
	AUTHENTICATED = "AUTHENTICATED"
)

func SimulateTCPConnWithLogin(port int) {

	conn, err := net.Dial("tcp", fmt.Sprintf(":%d", port))

	if err != nil {
		fmt.Println("Error when attempting to dial to tcp connection.")
		return
	}

	defer conn.Close()

	// read in arguments to send over this tcp connection
	reader := bufio.NewReader(os.Stdin)

	var accessToken string

	// -- attempt to authenticate --
	for {

		// read response and log it
		res, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println("Error when attempting to read from connection.")
		}

		fmt.Println("Message from server:", res)

		// attempt to check if server provided an authenticated status
		resPair := strings.SplitN(res, ":", 2)

		// check the correct format has been received before checking status
		if len(resPair) == 2 {
			status := resPair[0]
			serverMsg := resPair[1]

			fmt.Printf("From Server:\nStatus: %s\nMessage: %s\n\n", status, serverMsg)

			fmt.Println("If check between status:", status == AUTHENTICATED)

			if status == AUTHENTICATED {
				fmt.Println("Server authenticated.")
				accessToken = serverMsg // temp
				// exit out of the infinite auth loop
				break
			}
			fmt.Println("OUT OF BOUNDS")
		}

		fmt.Print("Enter: ")
		msg, _ := reader.ReadString('\n')

		_, err = conn.Write([]byte(msg))

		if err != nil {
			fmt.Println("Error sending message:", err)
			return
		}

	}

	// -- access server with authentication --

	for {
		fmt.Println("Entering Authenticated Message Loop")

		fmt.Printf("accessToken: %s\n\n", accessToken)

		fmt.Print("Enter [cmd] + [message]: ")
		msg, _ := reader.ReadString('\n')

		// append jwt to message for authorization
		// message format is [jwt accessToekn] [cmd] [message]
		authMsg := fmt.Sprintf("%s %s", accessToken, msg)

		// TODO: use jwt access token

		_, err = conn.Write([]byte(authMsg))

		if err != nil {
			fmt.Println("Error sending message:", err)
			return
		}

		// read response and log it
		res, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println("Error when attempting to read from connection.")
		}

		fmt.Println("Message from server:", res)
	}

}
