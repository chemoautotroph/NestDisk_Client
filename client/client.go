package client

import (
	"log"
	"myClient/config"
	"net"
	"strconv"
	"time"
)

func Client(){
	for i:=0; i<1; i++{
		conn := establishConn(i)
		// time.Sleep(2*time.Second)
		if conn != nil{
			// fmt.Println("conn is ",conn)
		}
	}
}

func establishConn(i int) net.Conn{
	address := config.Config.GetString("port")
	conn, err := net.DialTimeout("tcp", address,time.Second * 5)
	if err!= nil{
		log.Fatalf("can not connect with err: %v\n",err)
		return nil
	}

	log.Println(i, "connect to server ok")
	_, err = conn.Write([]byte("nky"+strconv.Itoa(i)))
	if err != nil {
		log.Fatalln("conn.Write Unexpected Error")
		return nil
	}
	
	return conn
}
