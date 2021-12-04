package client

import (
	"bufio"
	"fmt"
	"log"
	"myClient/protocol"
	"net"
	"os"
	"strings"
)

func send(conn net.Conn, message []byte) {
	//session := strconv.FormatInt(time.Now().Unix(), 10)
	//message = append([]byte(session), message...)
	_, err := conn.Write(protocol.Enpack(message))
	if err != nil {
		log.Fatalln("conn.Write Unexpected Error", err)
	}
}

func sendCommand(conn net.Conn, senderChannel chan []byte) {
	for{
		select {
		case message := <-senderChannel:
			send(conn, []byte(message))
			fmt.Println("sending command", message)
			break
		}
		buf := make([]byte, 1024)
		receiveCommand(buf, conn)
	}
}

func typeCommand (senderChannel chan []byte){
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("You can type now.")
	for{
		input, err := inputReader.ReadString('\n')
		command := strings.TrimSuffix(strings.TrimSuffix(input, "\n"), "\r")
		if err != nil {
			log.Fatalln("input err: ", err)
		} else {
			senderChannel <- []byte(command)
		}
	}
}


func receiveCommand (rep []byte, conn net.Conn) {
	tempbuf := make([]byte, 0)
	readerChannel := make(chan []byte, 16)
	buffer := make([]byte, 1024*1024)
	n, _ := conn.Read(buffer)

	tempbuf = protocol.Depack(append(tempbuf, buffer[:n]...), readerChannel)
	readCommand(conn, readerChannel)
}

func readCommand (conn net.Conn, readerChannel chan []byte){
	select {
	case data := <- readerChannel:
		fmt.Println("receive respond: ", string(data))
	}
}
