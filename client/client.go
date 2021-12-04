package client

import (
	"log"
	"myClient/config"
	"net"
	"time"
)

func UserTest() {
	// establishConn()
}

func Init() {
	for {
		senderChannel := make(chan []byte, 1)
		go typeCommand(senderChannel)
		conn := establishConn(senderChannel)
		// time.Sleep(1 * time.Second)

		if conn != nil {
			// fmt.Println("conn is ",conn)
		}


	}
}

func establishConn(senderChannel chan []byte ,i ...string) net.Conn {
	address := config.Config.GetString("port")
	conn, err := net.DialTimeout("tcp", address, 2*time.Second)
	if err != nil {
		log.Fatalf("can not connect with err: %v\n", err)
		return nil
	}
	if i != nil {
		log.Println(i, "connect to server ok")
	}
	sendCommand(conn, senderChannel)

	return conn
}

