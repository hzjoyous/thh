package template

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"time"
)

func init() {
	addConsole("tcp", "tcp",
		func() {

			fmt.Println("tcp")

			var tcpAddr *net.TCPAddr
			tcpAddr, _ = net.ResolveTCPAddr("tcp", "127.0.0.1:8282")

			conn, err := net.DialTCP("tcp", nil, tcpAddr)

			if err != nil {
				fmt.Println("Client connect error ! " + err.Error())
				return
			}

			defer conn.Close()

			fmt.Println(conn.LocalAddr().String() + " : Client connected!")

			onMessageReceived(conn)

		})
}

func onMessageReceived(conn *net.TCPConn) {

	reader := bufio.NewReader(conn)
	b := []byte(conn.LocalAddr().String() + " Say hello to Server... \n")
	_, err := conn.Write(b)
	if err != nil {
		return
	}
	for {
		msg, err := reader.ReadString('\n')
		fmt.Println("server:" + msg)

		if err != nil || err == io.EOF {
			fmt.Println(err)
			break
		}
		time.Sleep(time.Second * 13)

		fmt.Println("writing...")

		b := []byte(conn.LocalAddr().String() + " write dataRep to Server... \n")
		_, err = conn.Write(b)

		if err != nil {
			fmt.Println(err)
			break
		}
	}
}
