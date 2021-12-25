package client

import (
	"bufio"
	"fmt"
	"log"
	"myClient/protocol"
	"myClient/utils"
	"net"
	"os"
	"strings"
)


var UserName, tempName string


func send(conn net.Conn, message []byte) {
	//session := strconv.FormatInt(time.Now().Unix(), 10)
	//message = append([]byte(session), message...)
	_, err := conn.Write(protocol.Enpack(message))
	if err != nil {
		log.Fatalln("conn.Write Unexpected Error", err)
	}
}

func sendWhileLogin(conn net.Conn, message []byte) {
	if UserName == ""{
		fmt.Println("Err: Undefined UserName")
		return
	}
	m := utils.BytesCombine(message, []byte(" "), []byte(UserName))
	_, err := conn.Write(protocol.Enpack(m))
	if err != nil {
		log.Fatalln("conn.Write Unexpected Error", err)
	}
}



func sendCommand(conn net.Conn, senderChannel chan []byte) {
	for {
		select {
		case message := <-senderChannel:
			command := strings.Split(string(message), " ")
			if command[0] == "login"{
				tempName = command[1]
			}
			if UserName != ""{
				sendWhileLogin(conn, message)
				// fmt.Println("sending command while logged in")
				break
			}
			send(conn, []byte(message))
			// fmt.Println("sending command", message)
			break
		}
		buf := make([]byte, 1024)
		receiveCommand(buf, conn)
	}
}

func typeCommand(senderChannel chan []byte) {
	inputReader := bufio.NewReader(os.Stdin)
	logo := getLogo()
	fmt.Println(logo)
	for {
		input, err := inputReader.ReadString('\n')
		command := strings.TrimSuffix(strings.TrimSuffix(input, "\n"), "\r")
		if err != nil {
			log.Fatalln("input err: ", err)
		} else {
			senderChannel <- []byte(command)
		}
	}
}

func receiveCommand(rep []byte, conn net.Conn) {
	tempbuf := make([]byte, 0)
	readerChannel := make(chan []byte, 16)
	buffer := make([]byte, 1024*1024)
	n, _ := conn.Read(buffer)

	tempbuf = protocol.Depack(append(tempbuf, buffer[:n]...), readerChannel)
	readCommand(conn, readerChannel)
}

func readCommand(conn net.Conn, readerChannel chan []byte) {
	select {
	case data := <-readerChannel:
		switch string(data) {
		case "Login Successful":
			UserName = tempName
		default:
			fmt.Println("server respond -> ", string(data))
		}
	}
}

func getLogo() string {
	logo := "\n __    __              __      _______   __            __       \n/  \\  /  |            /  |    /       \\ /  |          /  |      \n$$  \\ $$ |  ______   _$$ |_   $$$$$$$  |$$/   _______ $$ |   __ \n$$$  \\$$ | /      \\ / $$   |  $$ |  $$ |/  | /       |$$ |  /  |\n$$$$  $$ |/$$$$$$  |$$$$$$/   $$ |  $$ |$$ |/$$$$$$$/ $$ |_/$$/ \n$$ $$ $$ |$$    $$ |  $$ | __ $$ |  $$ |$$ |$$      \\ $$   $$<  \n$$ |$$$$ |$$$$$$$$/   $$ |/  |$$ |__$$ |$$ | $$$$$$  |$$$$$$  \\ \n$$ | $$$ |$$       |  $$  $$/ $$    $$/ $$ |/     $$/ $$ | $$  |\n$$/   $$/  $$$$$$$/    $$$$/  $$$$$$$/  $$/ $$$$$$$/  $$/   $$/ \n                                                                \n                                                                \n                                                                \n"
	return logo
}
